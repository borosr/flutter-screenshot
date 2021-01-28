package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	ModeBoth  deviceMode = "both"
	ModeDark  deviceMode = "dark"
	LightDark deviceMode = "light"
)

type deviceMode string

type Data struct {
	Cmd     string  `yaml:"command"`
	Devices Devices `yaml:"devices"`
}

type Devices struct {
	IOS     []Device `yaml:"ios"`
	Android []Device `yaml:"android"`
}

type Device struct {
	Name string `yaml:"name"`
	Mode string `yaml:"mode"` // can be dark, light, both
}

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
