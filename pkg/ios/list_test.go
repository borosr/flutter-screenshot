package ios

import (
	"testing"

	"github.com/borosr/flutter-screenshot/pkg/ios/config"
	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/brianvoe/gofakeit/v6"
)

func TestDevice_ListShutdown(t *testing.T) {
	deviceID := gofakeit.UUID()
	deviceName := gofakeit.AppName()
	d := Device{
		Config: config.Config{
			Devices: config.Devices{
				deviceID + deviceName: []config.Device{
					{
						MinimalDevice: config.MinimalDevice{
							UDID:  deviceID,
							Name:  deviceName,
							State: "Shutdown",
						},
					},
				},
			},
		},
	}

	thenCheckTheListResult(t,
		d.List(),
		deviceName, deviceID,
		func(t *testing.T, p types.Instance) {
			if p.State != types.StateShutdown {
				t.Errorf("device state should be %v, instead of %v", types.StateShutdown, p.State)
			}
			if p.Kind != types.KindIos {
				t.Errorf("device kind should be %v, instead of %v", types.KindIos, p.Kind)
			}
		})
}

func TestDevice_ListBooted(t *testing.T) {
	deviceID := gofakeit.UUID()
	deviceName := gofakeit.AppName()
	d := Device{
		Config: config.Config{
			Devices: config.Devices{
				deviceID + deviceName: []config.Device{
					{
						MinimalDevice: config.MinimalDevice{
							UDID:  deviceID,
							Name:  deviceName,
							State: "Booted",
						},
					},
				},
			},
		},
	}

	thenCheckTheListResult(t,
		d.List(),
		deviceName, deviceID,
		func(t *testing.T, p types.Instance) {
			if p.State != types.StateBooted {
				t.Errorf("device state should be %v, instead of %v", types.StateBooted, p.State)
			}
			if p.Kind != types.KindIos {
				t.Errorf("device kind should be %v, instead of %v", types.KindIos, p.Kind)
			}
		})
}

func TestDevice_ListErrored(t *testing.T) {
	deviceID := gofakeit.UUID()
	deviceName := gofakeit.AppName()
	d := Device{
		Config: config.Config{
			Devices: config.Devices{
				deviceID + deviceName: []config.Device{
					{
						MinimalDevice: config.MinimalDevice{
							UDID:  deviceID,
							Name:  deviceName,
							State: "Errored",
						},
						AvailabilityError: "Error",
					},
				},
			},
		},
	}

	pairs := d.List()
	if pairs == nil {
		t.Fatal("missing result")
	}
	if len(pairs) != 0 {
		t.Fatalf("generated pairs amount should be 0, instead of %d", len(pairs))
	}
}

func thenCheckTheListResult(
	t *testing.T,
	pairs types.Pairs,
	deviceName, deviceID string,
	stateKindCheck func(t *testing.T, p types.Instance)) {
	if pairs == nil {
		t.Fatal("missing result")
	}
	if len(pairs) != 1 {
		t.Fatalf("generated pairs amount should be 1, instead of %d", len(pairs))
	}
	if p, ok := pairs[deviceName]; !ok {
		t.Errorf("missing key from map, it should be %s", deviceName)
	} else {
		if p.ID != deviceID {
			t.Errorf("device id should be %s, instead of %s", deviceID, p.ID)
		}
		stateKindCheck(t, p)
	}
}
