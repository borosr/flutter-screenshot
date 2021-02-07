package android

import (
	"strconv"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/brianvoe/gofakeit/v6"
)

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
