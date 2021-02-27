package types

import "encoding/json"

const (
	StateShutdown state = iota
	StateBooted
)

type state uint8

const (
	KindUnknow Kind = iota
	KindIos
	KindAndroid
)

type Kind uint8

// Instance represents a virtual device
// ID is the unique identifier of the virtual device
// DebugPort is only available on Android devices
// State stores the current state of the device, if successfully started setting it to Booted, otherwise Shutdown
// Kind stores the type of the device, it can be iOS or Android
type Instance struct {
	ID        string `json:"id"`
	DebugPort string `json:"debug_port"`
	State     state  `json:"state"`
	Kind      Kind   `json:"kind"`
}

// Pairs is an in memory key-value store for devices
// key is the name of the device
// the value stores the details
type Pairs map[string]Instance

// String marshals the Pairs instance to JSON string
func (p Pairs) String() string {
	if p == nil {
		return "null"
	}
	res, err := json.Marshal(p)
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return string(res)
}
