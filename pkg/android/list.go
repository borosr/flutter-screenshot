package android

import (
	"strconv"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/brianvoe/gofakeit/v6"
)

// List converts the parsed devices to the common format
// and assign a random debug port to the device between 5555 and 5586
func (d Device) List() types.Pairs {
	p := types.Pairs{}
	for _, device := range d.Config.Devices {
		p[device] = types.Instance{
			ID:        device,
			DebugPort: strconv.Itoa(gofakeit.Number(5555, 5586)),
			State:     types.StateShutdown,
			Kind:      types.KindAndroid,
		}
	}

	return p
}
