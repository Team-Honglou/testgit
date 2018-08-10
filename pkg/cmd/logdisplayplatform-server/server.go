package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/facebookgo/inject"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/bus"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/middleware"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/registry"

	"golang.org/x/sync/errgroup"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/api"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/log"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/login"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/setting"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/social"

	// self registering services
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/extensions"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/metrics"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/plugins"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/alerting"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/cleanup"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/notifications"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/provisioning"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/rendering"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/search"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/services/sqlstore"
	_ "github.com/logdisplayplatform/logdisplayplatform/pkg/tracing"
)

func NewLogDisplayPlatformServer() *LogDisplayPlatformServerImpl {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	return &LogDisplayPlatformServerImpl{
		context:       childCtx,
		shutdownFn:    shutdownFn,
		childRoutines: childRoutines,
		log:           log.New("server"),
		cfg:           setting.NewCfg(),
	}
}

type LogDisplayPlatformServerImpl struct {
	context            context.Context
	shutdownFn         context.CancelFunc
	childRoutines      *errgroup.Group
	log                log.Logger
	cfg                *setting.Cfg
	shutdownReason     string
	shutdownInProgress bool

	RouteRegister api.RouteRegister `inject:""`
	HttpServer    *api.HTTPServer   `inject:""`
}

func (g *LogDisplayPlatformServerImpl) Run() error {
	g.loadConfiguration()
	g.writePIDFile()

	login.Init()
	social.NewOAuthService()

	serviceGraph := inject.Graph{}
	serviceGraph.Provide(&inject.Object{Value: bus.GetBus()})
	serviceGraph.Provide(&inject.Object{Value: g.cfg})
	serviceGraph.Provide(&inject.Object{Value: api.NewRouteRegister(middleware.RequestMetrics, middleware.RequestTracing)})

	// self registered services
	services := registry.GetServices()

	// Add all services to dependency graph
	for _, service := range services {
		serviceGraph.Provide(&inject.Object{Value: service.Instance})
	}

	serviceGraph.Provide(&inject.Object{Value: g})

	// Inject dependencies to services
	if err := serviceGraph.Populate(); err != nil {
		return fmt.Errorf("Failed to populate service dependency: %v", err)
	}

	// Init & start services
	for _, service := range services {
		if registry.IsDisabled(service.Instance) {
			continue
		}

		g.log.Info("Initializing " + service.Name)

		if err := service.Instance.Init(); err != nil {
			return fmt.Errorf("Service init failed: %v", err)
		}
	}

	// Start background services
	for _, srv := range services {
		// variable needed for accessing loop variable in function callback
		descriptor := srv
		service, ok := srv.Instance.(registry.BackgroundService)
		if !ok {
			continue
		}

		if registry.IsDisabled(descriptor.Instance) {
			continue
		}

		g.childRoutines.Go(func() error {
			// Skip starting new service when shutting down
			// Can happen when service stop/return during startup
			if g.shutdownInProgress {
				return nil
			}

			err := service.Run(g.context)

			// If error is not canceled then the service crashed
			if err != context.Canceled && err != nil {
				g.log.Error("Stopped "+descriptor.Name, "reason", err)
			} else {
				g.log.Info("Stopped "+descriptor.Name, "reason", err)
			}

			// Mark that we are in shutdown mode
			// So more services are not started
			g.shutdownInProgress = true
			return err
		})
	}

	sendSystemdNotification("READY=1")

	return g.childRoutines.Wait()
}

func (g *LogDisplayPlatformServerImpl) loadConfiguration() {
	err := g.cfg.Load(&setting.CommandLineArgs{
		Config:   *configFile,
		HomePath: *homePath,
		Args:     flag.Args(),
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start logdisplayplatform. error: %s\n", err.Error())
		os.Exit(1)
	}

	g.log.Info("Starting "+setting.ApplicationName, "version", version, "commit", commit, "compiled", time.Unix(setting.BuildStamp, 0))
	g.cfg.LogConfigSources()
}

func (g *LogDisplayPlatformServerImpl) Shutdown(reason string) {
	g.log.Info("Shutdown started", "reason", reason)
	g.shutdownReason = reason
	g.shutdownInProgress = true

	// call cancel func on root context
	g.shutdownFn()

	// wait for child routines
	g.childRoutines.Wait()
}

func (g *LogDisplayPlatformServerImpl) Exit(reason error) {
	// default exit code is 1
	code := 1

	if reason == context.Canceled && g.shutdownReason != "" {
		reason = fmt.Errorf(g.shutdownReason)
		code = 0
	}

	g.log.Error("Server shutdown", "reason", reason)

	log.Close()
	os.Exit(code)
}

func (g *LogDisplayPlatformServerImpl) writePIDFile() {
	if *pidFile == "" {
		return
	}

	// Ensure the required directory structure exists.
	err := os.MkdirAll(filepath.Dir(*pidFile), 0700)
	if err != nil {
		g.log.Error("Failed to verify pid directory", "error", err)
		os.Exit(1)
	}

	// Retrieve the PID and write it.
	pid := strconv.Itoa(os.Getpid())
	if err := ioutil.WriteFile(*pidFile, []byte(pid), 0644); err != nil {
		g.log.Error("Failed to write pidfile", "error", err)
		os.Exit(1)
	}

	g.log.Info("Writing PID file", "path", *pidFile, "pid", pid)
}

func sendSystemdNotification(state string) error {
	notifySocket := os.Getenv("NOTIFY_SOCKET")

	if notifySocket == "" {
		return fmt.Errorf("NOTIFY_SOCKET environment variable empty or unset.")
	}

	socketAddr := &net.UnixAddr{
		Name: notifySocket,
		Net:  "unixgram",
	}

	conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)

	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(state))

	conn.Close()

	return err
}
