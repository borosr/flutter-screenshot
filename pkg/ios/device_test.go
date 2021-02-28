package ios

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/borosr/flutter-screenshot/pkg/ios/config"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		raw, _ := json.Marshal(config.Config{})
		_, _ = w.Write(raw)
	})

	d := New()
	if !d.Config.Loaded {
		t.Error("config isn't successfully loaded")
	}
}

func TestLoadConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		raw, _ := json.Marshal(config.Config{})
		_, _ = w.Write(raw)
	})

	c := loadConfig(true)
	if !c.Loaded {
		t.Error("config isn't successfully loaded")
	}
}

func TestLoadConfigCmdRunError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(errors.New("invalid command"))
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		raw, _ := json.Marshal(config.Config{})
		_, _ = w.Write(raw)
	})

	c := loadConfig(true)
	if c.Loaded {
		t.Error("config shouldn't be loaded")
	}
}

func TestLoadConfigInvalidFormat(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockExecutable := NewMockExecutable(ctrl)
	execute = mockExecute(mockExecutable)

	mockExecutable.EXPECT().Run().Return(nil)
	mockExecutable.EXPECT().String().Return("")
	mockExecutable.EXPECT().Stdout(gomock.Any()).DoAndReturn(func(w io.Writer) {
		_, _ = w.Write([]byte("invalid json format"))
	})

	c := loadConfig(true)
	if c.Loaded {
		t.Error("config shouldn't be loaded")
	}
}

func TestIsNotMac(t *testing.T) {
	b := &bytes.Buffer{}
	log.SetOutput(b)
	c := loadConfig(false)
	if !c.Loaded {
		t.Error("config shouldn't be loaded on mac")
	}
	if !strings.Contains(b.String(), skippingIosSimulatorMsg) {
		t.Errorf("Console should contains the following log: %s", skippingIosSimulatorMsg)
	}
}
