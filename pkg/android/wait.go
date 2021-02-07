package android

import (
	"errors"

	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

const (
	maxErrorCount = 10
	deviceState   = "device"
)

var (
	ErrWaitingTooManyErrors = errors.New("too many errors got when waiting to startup")
)

func (d Device) WaitUntilBooted(instance types.Instance) error {
	cmd := execute("adb", "-s", "emulator-"+instance.DebugPort, "wait-for-device")
	log.Debugf("Emulator state: Executing cmd: %s", cmd.String())
	return cmd.Run()
}
