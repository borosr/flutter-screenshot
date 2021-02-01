package ios

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/borosr/flutter-screenshot/src/device/types"
	"github.com/sirupsen/logrus"
)

func (d Device) Create(name string) (string, types.Kind, error) {
	for _, deviceType := range d.Config.DeviceTypes {
		if strings.EqualFold(deviceType.Name, name) {
			if deviceType.Name != name {
				logrus.Warnf("DeviceType's name is %s, but you're using %s, consider changing it", deviceType.Name, name)
			}

			return create(name, deviceType.Identifier)
		}
	}

	return "", types.KindUnknow, fmt.Errorf("error, device name not found: %s", name)
}

func create(name, identifier string) (string, types.Kind, error) {
	logrus.Infof("Creating device with name %s and device type %s", name, identifier)
	cmd := exec.Command("xcrun", "simctl", "create", name, identifier)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", types.KindUnknow, err
	}

	logrus.Debugf("Running following command: %s", cmd.String())
	outputLine := out.String()
	outputLine = strings.ReplaceAll(outputLine, "\n\t", "")
	logrus.Debugf("Got the following output: %s", outputLine)
	if idRegex.MatchString(outputLine) {
		logrus.Infof("Device with name %s successfully created, id is %s", name, outputLine)

		return outputLine, types.KindIos, nil
	}

	return "", types.KindUnknow, errors.New("device creation result invalid")
}
