package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ModeBoth is set if the device should start both light and dark modes
// ModeDark if the device should start in dark mode
// ModeLight if the device should start in light mode
const (
	ModeBoth  deviceMode = "both"
	ModeDark  deviceMode = "dark"
	ModeLight deviceMode = "light"
)

type deviceMode string

// Data is the highest level of the configuration
// contains the command for execution
// and the devices for run.
// More details is under Devices
type Data struct {
	Cmd     string  `yaml:"command"`
	Devices Devices `yaml:"devices"`
}

// Devices is wrapping all IOS and Android devices
// in two separate lists
type Devices struct {
	IOS     []Device `yaml:"ios"`
	Android []Device `yaml:"android"`
}

// Device represents the possible configuration properties of a device
type Device struct {
	Name string `yaml:"name"`
	Mode string `yaml:"mode"` // can be dark, light, both
}

// Read is reading the screenshots.yaml file and returning the parsed config.Data
func Read() (Data, error) {
	file, err := os.Open("screenshots.yaml")
	if err != nil {
		return Data{}, err
	}
	var data Data
	lines, err := ioutil.ReadAll(file)
	if err != nil {
		return Data{}, err
	}
	if err := yaml.Unmarshal(lines, &data); err != nil {
		return Data{}, err
	}

	return data, nil
}
