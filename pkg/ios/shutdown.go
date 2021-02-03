package ios

import (
	log "github.com/sirupsen/logrus"
)

func (d Device) Shutdown(id string) error {
	cmd := execute("xcrun", "simctl", "shutdown", id)
	log.Debugf("Shutdown: Executing cmd: %s", cmd.String())

	return cmd.Run()
}
