package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/cmd/logdisplayplatform-cli/commands"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/cmd/logdisplayplatform-cli/logger"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/cmd/logdisplayplatform-cli/services"
	"github.com/logdisplayplatform/logdisplayplatform/pkg/cmd/logdisplayplatform-cli/utils"
)

var version = "master"

func main() {
	setupLogging()

	app := cli.NewApp()
	app.Name = "LogDisplayPlatform cli"
	app.Usage = ""
	app.Author = "LogDisplayPlatform Project"
	app.Email = "https://github.com/logdisplayplatform/logdisplayplatform"
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "pluginsDir",
			Usage:  "path to the logdisplayplatform plugin directory",
			Value:  utils.GetLogDisplayPlatformPluginDir(runtime.GOOS),
			EnvVar: "GF_PLUGIN_DIR",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "url to the plugin repository",
			Value:  "https://logdisplayplatform.com/api/plugins",
			EnvVar: "GF_PLUGIN_REPO",
		},
		cli.StringFlag{
			Name:   "pluginUrl",
			Usage:  "Full url to the plugin zip file instead of downloading the plugin from logdisplayplatform.com/api",
			Value:  "",
			EnvVar: "GF_PLUGIN_URL",
		},
		cli.BoolFlag{
			Name:  "insecure",
			Usage: "Skip TLS verification (insecure)",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable debug logging",
		},
	}

	app.Before = func(c *cli.Context) error {
		services.Init(version, c.GlobalBool("insecure"))
		return nil
	}
	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound

	if err := app.Run(os.Args); err != nil {
		logger.Errorf("%v", err)
	}
}

func setupLogging() {
	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			logger.SetDebug(true)
		}
	}
}

func cmdNotFound(c *cli.Context, command string) {
	fmt.Printf(
		"%s: '%s' is not a %s command. See '%s --help'.\n",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}
