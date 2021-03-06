package android

import (
	"bytes"

	"github.com/borosr/flutter-screenshot/pkg/android/config"
)

// Device represents the Android device action strategy
type Device struct {
	Config config.Config
}

// New is creating a new Android device action strategy
func New() Device {
	return Device{
		Config: loadConfig(),
	}
}

func loadConfig() config.Config {
	c := config.Config{}
	cmd := execute("emulator", "-list-avds")
	var out bytes.Buffer
	cmd.Stdout(&out)
	if err := cmd.Run(); err != nil {
		return config.Config{}
	}

	c.Devices = config.UnmarshalDevices(out.Bytes())
	c.Loaded = true

	return c
}
