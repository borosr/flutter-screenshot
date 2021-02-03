package ios

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

func (d Device) SetTheme(i types.Instance, theme string) error {
	cmd := execute("xcrun", "simctl", "ui", i.ID, "appearance", theme)
	log.Debugf("Execute: Executing cmd: %s", cmd.String())

	return cmd.Run()
}
