package utils

import (
	"os"

	"github.com/logdisplayplatform/logdisplayplatform/pkg/cmd/logdisplayplatform-cli/logger"
)

func GetLogDisplayPlatformPluginDir(currentOS string) string {
	//currentOS := runtime.GOOS

	if currentOS == "windows" {
		return returnOsDefault(currentOS)
	}

	pwd, err := os.Getwd()

	if err != nil {
		logger.Error("Could not get current path. using default")
		return returnOsDefault(currentOS)
	}

	if isDevenvironment(pwd) {
		return "../data/plugins"
	}

	return returnOsDefault(currentOS)
}

func isDevenvironment(pwd string) bool {
	// if ../conf/defaults.ini exists, logdisplayplatform is not installed as package
	// that its in development environment.
	_, err := os.Stat("../conf/defaults.ini")
	return err == nil
}

func returnOsDefault(currentOs string) string {
	switch currentOs {
	case "windows":
		return "../data/plugins"
	case "darwin":
		return "/usr/local/var/lib/logdisplayplatform/plugins"
	case "freebsd":
		return "/var/db/logdisplayplatform/plugins"
	default: //"linux"
		return "/var/lib/logdisplayplatform/plugins"
	}
}
