package executor

import (
	"os"
	"os/exec"
	"strings"

	"github.com/borosr/flutter-screenshot/pkg"
	"github.com/borosr/flutter-screenshot/pkg/ios"
	"github.com/borosr/flutter-screenshot/src/config"
	"github.com/borosr/flutter-screenshot/src/device/types"
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

	execute(conf.Devices.IOS, conf.Cmd, ios.New())
	//execute(conf.Devices.Android, conf.Cmd, android.New())

	return nil
}

func execute(devices []config.Device, cmd string, device pkg.DeviceAction) {
	existingDevices := device.List()
	for _, d := range devices {
		d.Mode = strings.ToLower(d.Mode)
		log.Infof("Starting iOS device %s", d.Name)
		instance, ok := existingDevices[d.Name]
		if !ok {
			log.Warnf("Device %s not found in creted list, starting creation...", d.Name)
			id, kind, err := device.Create(d.Name)
			if err != nil {
				log.Errorf("Device %s not creatable, maybe misspelled the name, continue with the next device!", d.Name)
				log.Errorf("Error is: %v", err)

				continue
			}
			instance = types.Instance{
				ID:    id,
				Kind:  kind,
				State: types.StateBooted,
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

		switch {
		case d.Mode == "both":
			for _, t := range []string{DarkTheme, LightTheme} {
				d.Mode = t
				setThemeAndExecute(device, d, instance, cmd)
			}
		case d.Mode == "":
			d.Mode = LightTheme
			setThemeAndExecute(device, d, instance, cmd)
		case d.Mode == LightTheme || d.Mode == DarkTheme:
			setThemeAndExecute(device, d, instance, cmd)
		}

		log.Infof("Shutdown the device %s", instance.ID)
		if err := device.Shutdown(instance.ID); err != nil {
			log.Errorf("Shutdown error: %v", err)
		}
	}

}

func setThemeAndExecute(da pkg.DeviceAction, d config.Device, instance types.Instance, cmd string) {
	log.Infof("Set theme to %s at device %s", d.Mode, d.Name)
	if err := da.SetTheme(instance, d.Mode); err != nil {
		log.Errorf("Error setting the theme command: %v", err)
	}

	setScreenshotSubdirectoryName(d)
	if err := executeCommand(cmd, instance.ID); err != nil {
		log.Errorf("Error executing command: %v", err)
	} else {
		log.Info("Command successfully executed!")
	}
}

func setScreenshotSubdirectoryName(d config.Device) {
	emuDevName := strings.ReplaceAll(d.Name, " ", "_")
	subDirName := emuDevName + "_" + d.Mode
	log.Debugf("Subdiectory name %s", subDirName)
	if err := os.Setenv(emuDeviceEnv, subDirName); err != nil {
		log.Errorf("Can't set %s env to %s", emuDeviceEnv, emuDevName)
	}
}

func executeCommand(cmd, deviceID string) error {
	log.Infof("Executing command %s...", cmd)

	c := exec.Command("/bin/sh", "-c", cmd, "-d", deviceID)
	log.Debugf("Execute: Executing cmd: %s", c.String())
	c.Stdout = os.Stdout

	return c.Run()
}
