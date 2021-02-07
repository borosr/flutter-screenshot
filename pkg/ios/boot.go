package ios

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

func (d Device) Boot(i types.Instance) error {
	if i.State == types.StateBooted {
		log.Infof("%s is already booted, shutting it down", i.ID)
		if err := d.Shutdown(i); err != nil {
			return err
		}
	}
	log.Infof("Booting device with id %s", i.ID)
	cmd := execute("xcrun", "simctl", "boot", i.ID)
	log.Debugf("Boot: Executing cmd: %s", cmd.String())

	return cmd.Run()
}
