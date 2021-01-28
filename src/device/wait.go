package device

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func WaitUntilBooted(i Instance) error {
	if i.Kind == KindIos {
		cmd := exec.Command("xcrun", "simctl", "bootstatus", i.ID)
		log.Debugf("Bootstatus: Executing cmd: %s", cmd.String())
		cmd.Stdout = os.Stdout

		return cmd.Run()
	}
	log.Info("Android not supported yet")

	return nil
}
