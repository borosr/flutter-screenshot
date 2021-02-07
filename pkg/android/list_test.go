package android

import (
	"testing"

	"github.com/borosr/flutter-screenshot/pkg/android/config"
	"github.com/borosr/flutter-screenshot/src/device/types"
)

func TestDevice_List(t *testing.T) {
	deviceID := "Pixel_API_30"
	d := Device{
		Config: config.Config{
			Devices: []string{deviceID},
		},
	}

	list := d.List()
	if list == nil {
		t.Fatal("Result shouldn't be nil")
	}
	if l := len(list); l == 0 || l > 1 {
		t.Error("List length should not be zero or more then 1")
	}
	if d, ok := list[deviceID]; !ok {
		t.Errorf("Device name should be %s", deviceID)
	} else if d.ID != deviceID {
		t.Errorf("Device ID should be %s, instead of %s", deviceID, d.ID)
	} else if d.State != types.StateShutdown {
		t.Error("Device State should be shutdown, instead of booted")
	} else if d.Kind != types.KindAndroid {
		t.Error("Device Kind should be android, instead of ios")
	}
}
