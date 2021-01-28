package device

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/borosr/flutter-screenshot/src/device/ios"
	log "github.com/sirupsen/logrus"
)

var idRegex = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}")

func Create(name string) (Instance, error) {
	if c, ok := getIosConfig(); ok {
		for _, deviceType := range c.DeviceTypes {
			if strings.EqualFold(deviceType.Name, name) {
				if deviceType.Name != name {
					log.Warnf("DeviceType's name is %s, but you're using %s, consider changing it", deviceType.Name, name)
				}
				deviceID, err := createIosDevice(name, deviceType.Identifier)
				if err != nil {
					return Instance{}, err
				}

				return Instance{
					ID:    deviceID,
					State: StateBooted,
					Kind:  KindIos,
				}, err
			}
		}
	}
	if _, ok := getAndroidConfig(name); ok { // TODO support this too
		return Instance{}, errors.New("android not supported yet")
	}

	return Instance{}, fmt.Errorf("unable to create device with name %s", name)
}

func getAndroidConfig(name string) (interface{}, bool) {
	return nil, false
}

func createIosDevice(name, identifier string) (string, error) {
	log.Infof("Creating device with name %s and device type %s", name, identifier)
	cmd := exec.Command("xcrun", "simctl", "create", name, identifier)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	log.Debugf("Running following command: %s", cmd.String())
	outputLine := out.String()
	strings.ReplaceAll(outputLine, "\n\t", "")
	log.Debugf("Got the following output: %s", outputLine)
	if idRegex.MatchString(outputLine) {
		log.Infof("Device with name %s successfully created, id is %s", name, outputLine)
		return outputLine, nil
	}

	return "", errors.New("device creation result invalid")
}

func getIosConfig() (ios.Config, bool) {
	if iosConfig.Loaded {
		return iosConfig, true
	}
	config, err := loadIosConfig()

	return config, err == nil
}
