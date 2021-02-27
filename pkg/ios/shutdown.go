package ios

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

// Shutdown is trying to terminate the virtual device
func (d Device) Shutdown(i types.Instance) error {
	cmd := execute("xcrun", "simctl", "shutdown", i.ID)
	log.Debugf("Shutdown: Executing cmd: %s", cmd.String())

	return cmd.Run()
}
