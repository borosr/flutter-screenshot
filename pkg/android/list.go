package android

import "github.com/borosr/flutter-screenshot/src/device/types"

func (d Device) List() types.Pairs {
	p := types.Pairs{}
	for _, device := range d.Config.Devices {
		p[device] = types.Instance{
			ID: device,
			State: types.StateShutdown,
			Kind: types.KindAndroid,
		}
	}

	return p
}
