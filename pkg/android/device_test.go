package android

import (
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte("Parsing something...\nPixel_API_30\n"))
	})

	d := New()
	if !d.Config.Loaded {
		t.Error("config isn't successfully loaded")
	}
	if l := len(d.Config.Devices); l != 1 {
		t.Fatalf("Device count should be 1, instead of %d", l)
	}
	if d.Config.Devices[0] != "Pixel_API_30" {
		t.Errorf("Device should be %s, instead of %s", "Pixel_API_30", d.Config.Devices[0])
	}
}

func TestLoadConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte("Parsing something...\nPixel_API_30\n"))
	})

	c := loadConfig()
	if !c.Loaded {
		t.Error("config isn't successfully loaded")
	}
	if l := len(c.Devices); l != 1 {
		t.Fatalf("Device count should be 1, instead of %d", l)
	}
	if c.Devices[0] != "Pixel_API_30" {
		t.Errorf("Device should be %s, instead of %s", "Pixel_API_30", c.Devices[0])
	}
}

func TestLoadConfigCmdRunError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(errors.New("invalid command"))
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).
		DoAndReturn(func(w io.Writer) {
			_, _ = w.Write([]byte(""))
		})

	c := loadConfig()
	if c.Loaded {
		t.Error("config shouldn't be loaded on mac")
	}
}
