package device

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func Shutdown(id string, k kind) error {
	switch k {
	case KindIos:
		cmd := exec.Command("xcrun", "simctl", "shutdown", id)
		log.Debugf("Shutdown: Executing cmd: %s", cmd.String())

		return cmd.Run()
	case KindAndroid:
		log.Info("Android not supported yet")
	}
	return nil
}
