package ios

import (
	"os"

	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

// WaitUntilBooted is waiting until virtual device booted
func (d Device) WaitUntilBooted(i types.Instance) error {
	cmd := execute("xcrun", "simctl", "bootstatus", i.ID)
	log.Debugf("Bootstatus: Executing cmd: %s", cmd.String())
	cmd.Stdout(os.Stdout)

	return cmd.Run()
}
