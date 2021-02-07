package android

import (
	"errors"
	"testing"
)

func TestDevice_Create(t *testing.T) {
	d := Device{}
	_, _, err := d.Create("")
	if !errors.Is(err, ErrDeviceNotAbleToCreate) {
		t.Errorf("Create should throw the error %v, instead of throwing %v", ErrDeviceNotAbleToCreate, err)
	}
}
