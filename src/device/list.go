package device

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"runtime"

	"github.com/borosr/flutter-screenshot/src/device/ios"
	log "github.com/sirupsen/logrus"
)

const (
	StateShutdown state = iota
	StateBooted
)

type state uint8

const (
	KindIos kind = iota
	KindAndroid
)

type kind uint8

var (
	iosConfig ios.Config
)

type Instance struct {
	ID    string
	State state
	Kind  kind
}

// key is the name of the device
// the value contains the ID and other details
type Pairs map[string]Instance

func (p Pairs) buildIosDevices() (Pairs, error) {
	config, err := loadIosConfig()
	if err != nil {
		return p, err
	}
	for _, devices := range config.Devices {
		for _, device := range devices {
			if device.AvailabilityError != "" {
				log.Debugf("Device unavailable %s", device.Name)

				continue
			}
			var s = StateShutdown
			if device.State == "Booted" {
				s = StateBooted
			}
			p[device.Name] = Instance{
				ID:    device.UDID,
				State: s,
				Kind:  KindIos,
			}
		}
	}

	return p, nil
}

func loadIosConfig() (ios.Config, error) {
	if runtime.GOOS != "darwin" {
		log.Info("Skipping iOS simulators...")

		return ios.Config{Loaded: true}, nil
	}
	cmd := exec.Command("xcrun", "simctl", "list", "-j")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return ios.Config{}, err
	}

	var err error
	iosConfig, err = ios.UnmarshalConfig(out.Bytes())
	if err != nil {
		return ios.Config{}, err
	}
	iosConfig.Loaded = true
	return iosConfig, nil
}

func (p Pairs) buildAndroidDevices() (Pairs, error) {
	// TODO setup android here
	log.Info("Android not supported yet")
	return p, nil
}

func (p Pairs) String() string {
	res, err := json.Marshal(p)
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return string(res)
}

func List() (Pairs, error) {
	var p Pairs = make(map[string]Instance)
	var err error
	p, err = p.buildIosDevices()
	if err != nil {
		return nil, err
	}
	p, err = p.buildAndroidDevices()
	if err != nil {
		return nil, err
	}

	log.Infof("Found %d devices", len(p))
	log.Debugf("Devices are: %v", p.String())

	return p, nil
}
