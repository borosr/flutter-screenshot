package android

import (
	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

// SetTheme doesn't supported on Android devices
func (d Device) SetTheme(instance types.Instance, s string) error {
	log.Debug("Theme setup is not available on android device")
	return nil
}
