package android

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

func (d Device) Boot(i types.Instance) error {
	log.Infof("Booting device with id %s", i.ID)

	cmd := execute("emulator",
		"-port", i.DebugPort,
		"-avd", i.ID,
		"-no-boot-anim",
		"-netdelay", "none",
		"-no-snapshot")

	log.Debugf("Running following command: %s", cmd.String())

	return cmd.Start()
}
