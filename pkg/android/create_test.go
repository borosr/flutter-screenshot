package android

import (
	"fmt"
	"strings"
	"testing"
)

func TestDevice_Create(t *testing.T) {
	d := Device{}
	_, _, err := d.Create("")
	if err == nil {
		t.Fatal("Missing error")
	}
	if exp := fmt.Sprintf(errFmtDeviceNotAbleToCreate, strings.Join(d.Config.Devices, ",")); err.Error() != exp {
		t.Errorf("Create should throw the error %v, instead of throwing %v", exp, err)
	}
}
