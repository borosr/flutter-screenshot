package executor

import (
	"fmt"
	"os"
	"strings"

	"github.com/borosr/flutter-screenshot/pkg"
	"github.com/borosr/flutter-screenshot/pkg/android"
	"github.com/borosr/flutter-screenshot/pkg/exec"
	"github.com/borosr/flutter-screenshot/pkg/ios"
	"github.com/borosr/flutter-screenshot/src/config"
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

const (
	emuDeviceEnv = "EMU_DEVICE"

	errFmtExecuteIos              = "error running iOS devices: %w"
	errFmtCreateDevice            = "error when creating the device: %w|"
	errFmtBootDevice              = "error when booting the device: %w|"
	errFmtWaitUntilBooted         = "error waiting for simulator boot %w"
	errFmtSetThemeAndExecute      = "setThemeAndExecute error: %w"
	errFmtShutdownDevice          = "shutdown error: %w"
	errFmtSetThemeDevice          = "error setting the theme command: %w"
	errFmtSetScreenshotSubdirName = "error setScreenshotSubdirectoryName: %w"
	errFmtExecuteCmd              = "error executing command: %w"
	errFmtSetEnv                  = "cannot set %s env to %s"
)

var invoke exec.CommandExecutor = exec.Command

// Run is the entrypoint of the tool.
// First it reads the configurations from yaml,
// then runs sequentially the iOS and then the Android execution
func Run(configName string) error {
	conf, err := config.Read(configName)
	if err != nil {
		log.Error("Error when reading yaml file")

		return err
	}

	if err := execute(conf.Devices.IOS, conf.Cmd, ios.New()); err != nil {
		return fmt.Errorf(errFmtExecuteIos, err)
	}
	if err := execute(conf.Devices.Android, conf.Cmd, android.New()); err != nil {
		return fmt.Errorf(errFmtExecuteIos, err)
	}

	return nil
}

func execute(devices []config.Device, cmd string, device pkg.DeviceAction) error {
	existingDevices := device.List()
	var loopErr error
	for _, d := range devices {
		d.Mode = strings.ToLower(d.Mode)
		log.Infof("Starting device %s", d.Name)
		instance, ok := existingDevices[d.Name]
		if !ok {
			log.Warnf("Device %s not found in created list, starting creation...", d.Name)
			id, kind, err := device.Create(d.Name)
			if err != nil {
				log.Errorf("Device %s not creatable, maybe misspelled the name, continue with the next device!", d.Name)
				loopErr = fmt.Errorf(errFmtCreateDevice, err)

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
			loopErr = fmt.Errorf(errFmtBootDevice, err)

			continue
		}

		log.Info("Waiting a few seconds to startup...")
		if err := device.WaitUntilBooted(instance); err != nil {
			loopErr = fmt.Errorf(errFmtWaitUntilBooted, err)

			continue
		}

		switch {
		case d.Mode == config.ModeBoth.String():
			for _, t := range config.BothModes {
				d.Mode = t
				if err := setThemeAndExecute(device, d, instance, cmd); err != nil {
					loopErr = fmt.Errorf(errFmtSetThemeAndExecute, err)
				}
			}
		case d.Mode == "":
			d.Mode = config.ModeLight.String()
			if err := setThemeAndExecute(device, d, instance, cmd); err != nil {
				loopErr = fmt.Errorf(errFmtSetThemeAndExecute, err)
			}
		case d.Mode == config.ModeLight.String() || d.Mode == config.ModeDark.String():
			if err := setThemeAndExecute(device, d, instance, cmd); err != nil {
				loopErr = fmt.Errorf(errFmtSetThemeAndExecute, err)
			}
		}

		log.Infof("Shutdown the device %s", instance.ID)
		if err := device.Shutdown(instance); err != nil {
			loopErr = fmt.Errorf(errFmtShutdownDevice, err)
		}
	}

	return loopErr
}

func setThemeAndExecute(da pkg.DeviceAction, d config.Device, instance types.Instance, cmd string) error {
	log.Infof("Set theme to %s at device %s", d.Mode, d.Name)
	if err := da.SetTheme(instance, d.Mode); err != nil {
		return fmt.Errorf(errFmtSetThemeDevice, err)
	}

	if err := setScreenshotSubdirectoryName(d); err != nil {
		return fmt.Errorf(errFmtSetScreenshotSubdirName, err)
	}

	if err := executeCommand(cmd, instance.ID); err != nil {
		return fmt.Errorf(errFmtExecuteCmd, err)
	}
	log.Info("Command successfully executed!")

	return nil
}

func setScreenshotSubdirectoryName(d config.Device) error {
	emuDevName := strings.ReplaceAll(d.Name, " ", "_")
	subDirName := emuDevName + "_" + d.Mode
	log.Debugf("Subdiectory name %s", subDirName)
	if err := os.Setenv(emuDeviceEnv, subDirName); err != nil {
		return fmt.Errorf(errFmtSetEnv, emuDeviceEnv, emuDevName)
	}

	return nil
}

func executeCommand(cmd, deviceID string) error {
	log.Infof("Executing command %s...", cmd)

	c := invoke("/bin/sh", "-c", cmd, "-d", deviceID)
	log.Debugf("Execute: Executing cmd: %s", c.String())
	c.Stdout(os.Stdout)

	return c.Run()
}

func mockExecute(e exec.Executable) exec.CommandExecutor {
	return func(_ string, _ ...string) exec.Executable {
		return e
	}
}
