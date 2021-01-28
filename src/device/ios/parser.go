package ios

import "encoding/json"

type Config struct {
	DeviceTypes []DeviceType `json:"devicetypes"`
	Runtimes    []Runtime    `json:"runtimes"`
	Devices     Devices      `json:"devices"`
	Pairs       Pairs        `json:"pairs"`
	Loaded      bool         `json:"-"`
}

type DeviceType struct {
	MinRuntimeVersion int    `json:"minRuntimeVersion"`
	BundlePath        string `json:"bundlePath"`
	MaxRuntimeVersion int    `json:"maxRuntimeVersion"`
	Name              string `json:"name"`
	Identifier        string `json:"identifier"`
	ProductFamily     string `json:"productFamily"`
}

type Runtime struct {
	BundlePath   string `json:"bundlePath"`
	Buildversion string `json:"buildversion"`
	RuntimeRoot  string `json:"runtimeRoot"`
	Identifier   string `json:"identifier"`
	Version      string `json:"version"`
	IsAvailable  bool   `json:"isAvailable"`
	Name         string `json:"name"`
}

type Devices map[string][]Device

type Device struct {
	MinimalDevice
	AvailabilityError    string `json:"availabilityError"`
	DataPath             string `json:"dataPath"`
	LogPath              string `json:"logPath"`
	IsAvailable          bool   `json:"isAvailable"`
	DeviceTypeIdentifier string `json:"deviceTypeIdentifier"`
}

type MinimalDevice struct {
	UDID  string `json:"udid"`
	State string `json:"state"`
	Name  string `json:"name"`
}

type Pairs map[string]Pair

type Pair struct {
	State string        `json:"state"`
	Watch MinimalDevice `json:"watch"`
	Phone MinimalDevice `json:"phone"`
}

func UnmarshalConfig(data []byte) (Config, error) {
	var c Config
	err := json.Unmarshal(data, &c)

	return c, err
}
