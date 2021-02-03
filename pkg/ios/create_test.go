package ios

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/borosr/flutter-screenshot/pkg/ios/config"
	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func TestDevice_Create(t *testing.T) {
	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	d := Device{
		Config: config.Config{
			DeviceTypes: []config.DeviceType{
				{
					Name: deviceName,
				},
			},
			Loaded: true,
		},
	}

	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte(deviceID))
	})

	id, kind, err := d.Create(deviceName)
	if err != nil {
		t.Error(err)
	}
	if kind != types.KindIos {
		t.Error("kind should be ios")
	}
	if id != deviceID {
		t.Errorf("device id should be %s, instead of %s", deviceID, id)
	}
}

func TestDevice_CreateInvalidName(t *testing.T) {
	const deviceName = "iPhone X"
	d := Device{
		Config: config.Config{
			DeviceTypes: []config.DeviceType{
				{
					Name: "",
				},
			},
			Loaded: true,
		},
	}

	_, _, err := d.Create(deviceName)
	if err == nil {
		t.Fatal("missing error")
	}
	if want := fmt.Sprintf(errFmtDeviceNameNotFound, deviceName); err.Error() != want {
		t.Errorf("Error msg should be [%v], instead of [%v]", want, err)
	}
}

func TestDevice_CreateCmdRunError(t *testing.T) {
	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	d := Device{
		Config: config.Config{
			DeviceTypes: []config.DeviceType{
				{
					Name: deviceName,
				},
			},
			Loaded: true,
		},
	}

	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	errRunCmd := errors.New("invalid command")
	mockExecutable.EXPECT().Run().Return(errRunCmd)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte(deviceID))
	})

	id, kind, err := d.Create(deviceName)
	if err == nil {
		t.Fatal(err)
	}
	if kind != types.KindUnknow {
		t.Error("kind should be ios")
	}
	if !errors.Is(err, errRunCmd) {
		t.Errorf("device id should be %s, instead of %s", deviceID, id)
	}
}

func TestDevice_CreateCmdResultInvalid(t *testing.T) {
	const deviceName = "iPhone X"
	deviceID := gofakeit.UUID()
	d := Device{
		Config: config.Config{
			DeviceTypes: []config.DeviceType{
				{
					Name: deviceName,
				},
			},
			Loaded: true,
		},
	}

	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte("invalid id"))
	})

	id, kind, err := d.Create(deviceName)
	if err == nil {
		t.Fatal(err)
	}
	if kind != types.KindUnknow {
		t.Error("kind should be ios")
	}
	if !errors.Is(err, ErrDeviceCreationInvalidResult) {
		t.Errorf("device id should be %s, instead of %s", deviceID, id)
	}
}
