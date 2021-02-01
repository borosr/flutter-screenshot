package ios

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func (d Device) Shutdown(id string) error {
	cmd := exec.Command("xcrun", "simctl", "shutdown", id)
	log.Debugf("Shutdown: Executing cmd: %s", cmd.String())

	return cmd.Run()
}
