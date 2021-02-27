package android

import (
	"fmt"
	"strings"

	"github.com/borosr/flutter-screenshot/src/device/types"
)

const (
	errFmtDeviceNotAbleToCreate = "android virtual device (AVD) not able to create, please create it for yourself or choose one of the following [%s]"
)

// Create doesn't supported on Android devices
func (d Device) Create(s string) (string, types.Kind, error) {
	return "", 0, fmt.Errorf(errFmtDeviceNotAbleToCreate, strings.Join(d.Config.Devices, ","))
}
