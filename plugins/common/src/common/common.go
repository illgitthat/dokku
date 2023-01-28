package main

import (
	"fmt"
	"os"

	"github.com/dokku/dokku/plugins/common"
	flag "github.com/spf13/pflag"
)

func main() {
	quiet := flag.Bool("quiet", false, "--quiet: set DOKKU_QUIET_OUTPUT=1")
	global := flag.Bool("global", false, "--global: Whether global or app-specific")
	flag.Parse()
	cmd := flag.Arg(0)

	if *quiet {
		os.Setenv("DOKKU_QUIET_OUTPUT", "1")
	}

	var err error
	switch cmd {
	case "clone-data-directory":
		pluginName := flag.Arg(1)
		oldAppName := flag.Arg(2)
		newAppName := flag.Arg(3)
		err = common.CloneAppData(pluginName, oldAppName, newAppName)
	case "create-data-directory":
		pluginName := flag.Arg(1)
		appName := flag.Arg(2)
		err = common.CreateAppDataDirectory(pluginName, appName)
	case "delete-data-directory":
		pluginName := flag.Arg(1)
		appName := flag.Arg(2)
		err = common.RemoveAppDataDirectory(pluginName, appName)
	case "docker-cleanup":
		appName := flag.Arg(1)
		force := common.ToBool(flag.Arg(2))
		if *global {
			appName = "--global"
		}
		err = common.DockerCleanup(appName, force)
	case "is-deployed":
		appName := flag.Arg(1)
		if !common.IsDeployed(appName) {
			err = fmt.Errorf("App %v not deployed", appName)
		}
	case "image-is-cnb-based":
		image := flag.Arg(1)
		if common.IsImageCnbBased(image) {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case "image-is-herokuish-based":
		image := flag.Arg(1)
		appName := flag.Arg(2)
		if common.IsImageHerokuishBased(image, appName) {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case "migrate-data-directory":
		pluginName := flag.Arg(1)
		oldAppName := flag.Arg(2)
		newAppName := flag.Arg(3)
		err = common.MigrateAppDataDirectory(pluginName, oldAppName, newAppName)
	case "scheduler-detect":
		appName := flag.Arg(1)
		if *global {
			appName = "--global"
		}
		fmt.Print(common.GetAppScheduler(appName))
	case "setup-data-directory":
		pluginName := flag.Arg(1)
		err = common.SetupAppData(pluginName)
	case "verify-app-name":
		appName := flag.Arg(1)
		err = common.VerifyAppName(appName)
	default:
		err = fmt.Errorf("Invalid common command call: %v", cmd)
	}

	if err != nil {
		common.LogFailWithErrorQuiet(err)
	}
}
