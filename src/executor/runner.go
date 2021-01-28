package executor

import (
	"os"
	"os/exec"
	"strings"

	"github.com/borosr/flutter-screenshot/src/config"
	"github.com/borosr/flutter-screenshot/src/device"
	log "github.com/sirupsen/logrus"
)

const (
	emuDeviceEnv = "EMU_DEVICE"
	LightTheme   = "light"
	DarkTheme    = "dark"
)

func Run() error {
	conf, err := config.Read()
	if err != nil {
		log.Error("Error when reading yaml file")

		return err
	}

	devices, err := device.List()
	if err != nil {
		log.Error("Error when listing devices")

		return err
	}

	executeIos(devices, conf.Devices.IOS, conf.Cmd)

	return nil
}

func executeIos(devices device.Pairs, iosDevices []config.Device, cmd string) {
	for _, d := range iosDevices {
		d.Mode = strings.ToLower(d.Mode)
		log.Infof("Starting iOS device %s", d.Name)
		instance, ok := devices[d.Name]
		if !ok {
			log.Warnf("Device %s not found in creted list, starting creation...", d.Name)
			var err error
			instance, err = device.Create(d.Name)
			if err != nil {
				log.Errorf("Device %s not creatable, maybe misspelled the name, continue with the next device!", d.Name)
				log.Errorf("Error is: %v", err)

				continue
			}
		}

		log.Infof("Booting device %s %s...", d.Name, instance.ID)
		if err := device.Boot(instance); err != nil {
			log.Errorf("Error when booting the device: %v", err)

			continue
		}

		log.Info("Waiting a few seconds to startup...")
		if err := device.WaitUntilBooted(instance); err != nil {
			log.Errorf("Error waiting for simulator boot %v", err)
		}

		if d.Mode == "both" {
			for _, t := range []string{DarkTheme, LightTheme} {
				setIosThemeAndExecute(d, instance, cmd, t)
			}
		} else if d.Mode == "" {
			setIosThemeAndExecute(d, instance, cmd, LightTheme)
		} else if d.Mode == LightTheme || d.Mode == DarkTheme {
			setIosThemeAndExecute(d, instance, cmd, d.Mode)
		}

		log.Infof("Shutdown the device %s", instance.ID)
		if err := device.Shutdown(instance.ID, instance.Kind); err != nil {
			log.Errorf("Shutdown error: %v", err)
		}
	}

}

func setIosThemeAndExecute(d config.Device, instance device.Instance, cmd, t string) {
	log.Infof("Set theme to %s at device %s", t, d.Name)
	if err := changeIosTheme(instance.ID, t); err != nil {
		log.Errorf("Error executing command: %v", err)
	}

	setScreenshotSubdirectoryName(d, t)
	if err := executeCommand(cmd, instance.ID); err != nil {
		log.Errorf("Error executing command: %v", err)
	} else {
		log.Info("Command successfully executed!")
	}
}

func setScreenshotSubdirectoryName(d config.Device, theme string) {
	emuDevName := strings.ReplaceAll(d.Name, " ", "_")
	subDirName := emuDevName + "_" + theme
	log.Debugf("Subdiectory name %s", subDirName)
	if err := os.Setenv(emuDeviceEnv, subDirName); err != nil {
		log.Errorf("Can't set %s env to %s", emuDeviceEnv, emuDevName)
	}
}

func changeIosTheme(id, theme string) error {
	cmd := exec.Command("xcrun", "simctl", "ui", id, "appearance", theme)
	log.Debugf("Execute: Executing cmd: %s", cmd.String())

	return cmd.Run()
}

func executeCommand(cmd, deviceID string) error {
	log.Infof("Executing command %s...", cmd)

	c := exec.Command("/bin/sh", "-c", cmd, "-d", deviceID)
	log.Debugf("Execute: Executing cmd: %s", c.String())
	c.Stdout = os.Stdout

	return c.Run()
}
