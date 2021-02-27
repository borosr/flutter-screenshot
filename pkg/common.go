package pkg

import "github.com/borosr/flutter-screenshot/src/device/types"

// DeviceAction defines the common functionalities of virtual devices.
// It's a strategy interface for virtual device management.
type DeviceAction interface {
	List() types.Pairs
	Create(string) (string, types.Kind, error)
	Boot(types.Instance) error
	WaitUntilBooted(types.Instance) error
	SetTheme(types.Instance, string) error
	Shutdown(types.Instance) error
}
