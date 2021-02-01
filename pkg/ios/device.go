package ios

import (
	"bytes"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/borosr/flutter-screenshot/pkg/ios/config"
	log "github.com/sirupsen/logrus"
)

var idRegex = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}")

type Device struct {
	Config config.Config
}

func New() Device {
	return Device{
		Config: loadConfig(),
	}
}

func loadConfig() config.Config {
	if runtime.GOOS != "darwin" {
		log.Info("Skipping iOS simulators...")

		return config.Config{Loaded: true}
	}
	cmd := exec.Command("xcrun", "simctl", "list", "-j")

	log.Debugf("Running following command: %s", cmd.String())

	var out bytes.Buffer
	cmd.Stdout = &out
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
