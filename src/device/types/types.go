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

type Instance struct {
	ID    string
	State state
	Kind  Kind
}

// key is the name of the device
// the value contains the ID and other details
type Pairs map[string]Instance

func (p Pairs) String() string {
	res, err := json.Marshal(p)
	if err != nil {
		return `{"error":"` + err.Error() + `"}`
	}
	return string(res)
}
