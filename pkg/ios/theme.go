package ios

import (
	"os/exec"

	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

func (d Device) SetTheme(i types.Instance, theme string) error {
	cmd := exec.Command("xcrun", "simctl", "ui", i.ID, "appearance", theme)
	log.Debugf("Execute: Executing cmd: %s", cmd.String())

	return cmd.Run()
}
