package android

import (
	"errors"

	"github.com/borosr/flutter-screenshot/src/device/types"
)

var (
	ErrDeviceNotAbleToCreate = errors.New("android virtual device (AVD) not able to create, please create it for yourself")
)

func (d Device) Create(s string) (string, types.Kind, error) {
	return "", 0, ErrDeviceNotAbleToCreate
}
