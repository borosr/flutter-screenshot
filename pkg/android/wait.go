package android

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

const (
	deviceState = "device"
)

// WaitUntilBooted is waiting until virtual device booted
func (d Device) WaitUntilBooted(instance types.Instance) error {
	cmd := execute("adb", "-s", "emulator-"+instance.DebugPort, "wait-for-device")
	log.Debugf("Emulator state: Executing cmd: %s", cmd.String())
	return cmd.Run()
}
