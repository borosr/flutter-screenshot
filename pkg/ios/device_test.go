package ios

import (
	"encoding/json"
	"errors"
	"io"
	"runtime"
	"testing"

	"github.com/borosr/flutter-screenshot/pkg/ios/config"
	"github.com/golang/mock/gomock"
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

	c := loadConfig()
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

	c := loadConfig()
	if runtime.GOOS != "darwin" && !c.Loaded {
		t.Error("config should be loaded on this os")
	} else if c.Loaded {
		t.Error("config loading isn't failed")
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

	c := loadConfig()
	if runtime.GOOS != "darwin" && !c.Loaded {
		t.Error("config should be loaded on this os")
	} else if c.Loaded {
		t.Error("config loading isn't failed")
	}
}
