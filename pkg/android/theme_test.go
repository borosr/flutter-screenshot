package android

import (
	"testing"

	"github.com/borosr/flutter-screenshot/src/device/types"
)

func TestDevice_SetTheme(t *testing.T) {
	d := Device{}
	if err := d.SetTheme(types.Instance{}, "dark"); err != nil {
		t.Error(err)
	}
}
