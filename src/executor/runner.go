package executor

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/borosr/flutter-screenshot/src/config"
	"github.com/borosr/flutter-screenshot/src/device"
	log "github.com/sirupsen/logrus"
)

const (
	startupWaiting = 30 * time.Second

	emuDeviceEnv = "EMU_DEVICE"
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

func executeIos(devices device.Pairs, ios []config.Device, cmd string) {
	for _, i := range ios {
		log.Infof("Starting iOS device %s", i.Name)
		// TODO set mode light/dark
		instance, ok := devices[i.Name]
		if !ok {
			log.Warnf("Device %s not found in creted list, starting creation...", i.Name)
			var err error
			instance, err = device.Create(i.Name)
			if err != nil {
				log.Errorf("Device %s not creatable, maybe misspelled the name, continue with the next device!", i.Name)
				log.Errorf("Error is: %v", err)

				continue
			}
		}

		log.Infof("Booting device %s %s...", i.Name, instance.ID)
		if err := device.Boot(instance); err != nil {
			log.Errorf("Error when booting the device: %v", err)

			continue
		}
		log.Info("Waiting a few seconds to startup...")
		time.Sleep(startupWaiting)
		emuDevName := strings.ReplaceAll(i.Name, " ", "_")
		if err := os.Setenv(emuDeviceEnv, emuDevName); err != nil {
			log.Errorf("Can't set %s env to %s", emuDeviceEnv, emuDevName)
		}
		if err := executeCommand(cmd); err != nil {
			log.Errorf("Error executing command: %v", err)

			continue
		}

		log.Info("Command successfully executed!")
		log.Infof("Shutdown the device %s", instance.ID)
		if err := device.Shutdown(instance.ID, instance.Kind); err != nil {
			log.Errorf("Shutdown error: %v", err)
		}
	}

}

func executeCommand(cmd string) error {
	log.Infof("Executing command %s...", cmd)

	c := exec.Command("/bin/sh", "-c", cmd)
	c.Stdout = os.Stdout

	return c.Run()
}
