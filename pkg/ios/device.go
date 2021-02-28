package ios

import (
	"bytes"
	"regexp"
	"runtime"

	"github.com/borosr/flutter-screenshot/pkg/ios/config"
	log "github.com/sirupsen/logrus"
)

const (
	skippingIosSimulatorMsg = "Skipping iOS simulators..."
)

var idRegex = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}")

// Device represents the iOS device action strategy
type Device struct {
	Config config.Config
}

// New is creating a new iOS device action strategy
func New() Device {
	return Device{
		Config: loadConfig(runtime.GOOS == "darwin"),
	}
}

func loadConfig(isMac bool) config.Config {
	if !isMac {
		log.Info(skippingIosSimulatorMsg)

		return config.Config{Loaded: true}
	}
	cmd := execute("xcrun", "simctl", "list", "-j")

	log.Debugf("Running following command: %s", cmd.String())

	var out bytes.Buffer
	cmd.Stdout(&out)
	if err := cmd.Run(); err != nil {
		log.Errorf("Error executing command: %v", err)

		return config.Config{}
	}

	var err error
	iosConfig, err := config.Unmarshal(out.Bytes())
	if err != nil {
		log.Errorf("Error unmashal ios config: %v", err)

		return config.Config{}
	}
	iosConfig.Loaded = true

	return iosConfig
}
