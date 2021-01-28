package device

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func Boot(i Instance) error {
	if i.State == StateBooted {
		log.Infof("%s is already booted, shutting it down", i.ID)
		if err := Shutdown(i.ID, i.Kind); err != nil {
			return err
		}
	}

	log.Infof("Booting device with id %s", i.ID)
	switch i.Kind {
	case KindIos:
		cmd := exec.Command("xcrun", "simctl", "boot", i.ID)
		log.Debugf("Boot: Executing cmd: %s", cmd.String())

		return cmd.Run()
	case KindAndroid:
		log.Info("Android not supported yet")
	}

	return nil
}
