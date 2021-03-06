package config

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	mockCommand    = "echo \"hello world\""
	mockDeviceName = "iPhone X"
	mockContent    = "command: " + mockCommand + "\ndevices:\n  ios:\n  - name: " + mockDeviceName + "\n    mode: light\n  android: []\n"
)

func TestReadOnMac(t *testing.T) {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(file.Name(), []byte(mockContent), os.FileMode(0666)); err != nil {
		t.Fatal(err)
	}
	read, err := Read(file.Name(), true)
	if err != nil {
		t.Fatal(err)
	}
	if read.Cmd != mockCommand {
		t.Errorf("Command should be %s, instead of %s", mockCommand, read.Cmd)
	}
	if len(read.Devices.IOS) != 1 {
		t.Fatal("Invalid device count, only 1 iOS device should exist")
	}
	if read.Devices.IOS[0].Name != mockDeviceName {
		t.Errorf("IOS device name should be %s, instead of %s", mockCommand, read.Cmd)
	}
}

func TestReadOnNotMac(t *testing.T) {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(file.Name(), []byte(mockContent), os.FileMode(0666)); err != nil {
		t.Fatal(err)
	}
	read, err := Read(file.Name(), false)
	if err != nil {
		t.Fatal(err)
	}
	if read.Cmd != mockCommand {
		t.Errorf("Command should be %s, instead of %s", mockCommand, read.Cmd)
	}
	if len(read.Devices.IOS) != 0 {
		t.Fatal("Invalid device count, no iOS device should exist")
	}
}

func TestNotExistFile(t *testing.T) {
	_, err := Read(gofakeit.UUID()+".yaml", false)
	if !errors.Is(err, os.ErrNotExist) {
		t.Error("Read should return with error if file not exists")
	}
}

func TestInvalidFileFormat(t *testing.T) {
	temp, err := os.CreateTemp("", gofakeit.UUID())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := temp.WriteString("some string..."); err != nil {
		t.Fatal(err)
	}
	if _, err := Read(temp.Name(), false); err == nil {
		t.Error("Read should return with error if format is invalid")
	}
}
