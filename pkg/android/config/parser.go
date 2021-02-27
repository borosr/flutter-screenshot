package config

import (
	"regexp"
	"strings"
)

// Config stores the names of the existing Android Virtual Devices (AVD)
type Config struct {
	Devices []string
	Loaded  bool
}

// UnmarshalDevices building the device name list from the external command calls output
func UnmarshalDevices(data []byte) []string {
	var d []string

	str := string(regexp.MustCompile("Parsing (.*)\n").
		ReplaceAll(data, []byte("")))

	str = strings.ReplaceAll(str, "\r", "\n")

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if line != "" {
			d = append(d, line)
		}
	}

	return d
}
