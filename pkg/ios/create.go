package ios

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/borosr/flutter-screenshot/src/device/types"
	log "github.com/sirupsen/logrus"
)

const (
	errFmtDeviceNameNotFound = "error, device name not found: %s"
)

var (
	ErrDeviceCreationInvalidResult = errors.New("device creation result invalid")
)

func (d Device) Create(name string) (string, types.Kind, error) {
	for _, deviceType := range d.Config.DeviceTypes {
		if strings.EqualFold(deviceType.Name, name) {
			if deviceType.Name != name {
				log.Warnf("DeviceType's name is %s, but you're using %s, consider changing it", deviceType.Name, name)
			}

			return create(name, deviceType.Identifier)
		}
	}

	return "", types.KindUnknow, fmt.Errorf(errFmtDeviceNameNotFound, name)
}

func create(name, identifier string) (string, types.Kind, error) {
	log.Infof("Creating device with name %s and device type %s", name, identifier)
	cmd := execute("xcrun", "simctl", "create", name, identifier)
	var out bytes.Buffer
	cmd.Stdout(&out)
	if err := cmd.Run(); err != nil {
		return "", types.KindUnknow, err
	}

	log.Debugf("Running following command: %s", cmd.String())
	outputLine := out.String()
	outputLine = strings.ReplaceAll(outputLine, "\n\t", "")
	log.Debugf("Got the following output: %s", outputLine)
	if idRegex.MatchString(outputLine) {
		log.Infof("Device with name %s successfully created, id is %s", name, outputLine)

		return outputLine, types.KindIos, nil
	}

	return "", types.KindUnknow, ErrDeviceCreationInvalidResult
}
