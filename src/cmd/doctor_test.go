package cmd

import (
	"bytes"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestCheckiOS(t *testing.T) {
	b := &bytes.Buffer{}
	log.SetOutput(b)
	checkiOS(true)
	if !strings.Contains(b.String(), checkingIosCommandsMsg) {
		t.Errorf("Invalid output, missing %s", checkingIosCommandsMsg)
	}
	b.Reset()
	checkiOS(false)
	if b.String() != "" {
		t.Error("The log output should be empty")
	}
}

func TestCheckAndroid(t *testing.T) {
	b := &bytes.Buffer{}
	log.SetOutput(b)
	checkAndroid()
	if !strings.Contains(b.String(), checkingAndroidEnvsMsg) {
		t.Errorf("Invalid output, missing %s", checkingAndroidEnvsMsg)
	}
	if !strings.Contains(b.String(), androidEnvJavaHome) {
		t.Errorf("Invalid output, missing %s", androidEnvJavaHome)
	}
	if !strings.Contains(b.String(), androidEnvAndroidHome) {
		t.Errorf("Invalid output, missing %s", androidEnvAndroidHome)
	}
	if !strings.Contains(b.String(), androidEnvAndroidSdkRoot) {
		t.Errorf("Invalid output, missing %s", androidEnvAndroidSdkRoot)
	}
	if !strings.Contains(b.String(), checkingAndroidCommandsMsg) {
		t.Errorf("Invalid output, missing %s", checkingAndroidCommandsMsg)
	}
	if !strings.Contains(b.String(), androidCommandEmulator) {
		t.Errorf("Invalid output, missing %s", androidCommandEmulator)
	}
	if !strings.Contains(b.String(), androidCommandAdb) {
		t.Errorf("Invalid output, missing %s", androidCommandAdb)
	}
}

func TestGetResultIcon(t *testing.T) {
	const notFoundIcon = "\u274c"
	if s := getResultIcon(false); bytes.Equal([]byte(s), []byte(notFoundIcon)) {
		t.Errorf("Invalid result, state should be %s, instead of %s", notFoundIcon, s)
	}
	const foundIcon = "\u2705"
	if s := getResultIcon(true); bytes.Equal([]byte(s), []byte(foundIcon)) {
		t.Errorf("Invalid result, state should be %s, instead of %s", foundIcon, s)
	}
}
