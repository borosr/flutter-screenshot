package config

import "encoding/json"

// Config represents the highest level of the iOS configuration's output
type Config struct {
	DeviceTypes []DeviceType `json:"devicetypes"`
	Runtimes    []Runtime    `json:"runtimes"`
	Devices     Devices      `json:"devices"`
	Pairs       Pairs        `json:"pairs"`
	Loaded      bool         `json:"-"`
}

// DeviceType represents an iOS device type with details
type DeviceType struct {
	MinRuntimeVersion int    `json:"minRuntimeVersion"`
	BundlePath        string `json:"bundlePath"`
	MaxRuntimeVersion int    `json:"maxRuntimeVersion"`
	Name              string `json:"name"`
	Identifier        string `json:"identifier"`
	ProductFamily     string `json:"productFamily"`
}

// Runtime represents an iOS runtime with details
type Runtime struct {
	BundlePath   string `json:"bundlePath"`
	Buildversion string `json:"buildversion"`
	RuntimeRoot  string `json:"runtimeRoot"`
	Identifier   string `json:"identifier"`
	Version      string `json:"version"`
	IsAvailable  bool   `json:"isAvailable"`
	Name         string `json:"name"`
}

// Devices represents ID-Device pairs
type Devices map[string][]Device

// Device represents an iOS device with all the details
type Device struct {
	MinimalDevice
	AvailabilityError    string `json:"availabilityError"`
	DataPath             string `json:"dataPath"`
	LogPath              string `json:"logPath"`
	IsAvailable          bool   `json:"isAvailable"`
	DeviceTypeIdentifier string `json:"deviceTypeIdentifier"`
}

// MinimalDevice represents the common values of devices
type MinimalDevice struct {
	UDID  string `json:"udid"`
	State string `json:"state"`
	Name  string `json:"name"`
}

// Pairs represents multiple iPhone-Watch pairs
type Pairs map[string]Pair

// Pair represents an iPhone-Watch pair
type Pair struct {
	State string        `json:"state"`
	Watch MinimalDevice `json:"watch"`
	Phone MinimalDevice `json:"phone"`
}

// Unmarshal creates a Config from the external command calls output
func Unmarshal(data []byte) (Config, error) {
	var c Config
	err := json.Unmarshal(data, &c)

	return c, err
}
