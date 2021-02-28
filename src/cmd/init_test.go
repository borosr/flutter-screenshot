package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	if err := createFile(DefaultConfigFileName); err != nil {
		t.Fatal(err)
	}
	file, err := ioutil.ReadFile(DefaultConfigFileName)
	if err != nil {
		t.Fatal(err)
	}
	if string(file) == "" {
		t.Error("File shouldn't be empty")
	}
	if !bytes.Contains(file, []byte(exampleCommandValue)) {
		t.Errorf("Invalid example file, it should contains the following command %s", exampleCommandValue)
	}
	if !bytes.Contains(file, []byte(exampleDeviceName)) {
		t.Errorf("Invalid example file, it should contains the following device name %s", exampleDeviceName)
	}
	t.Cleanup(func() {
		if err := os.Remove(DefaultConfigFileName); err != nil {
			t.Error(err)
		}
	})
}
