package cmd

import (
	"io/ioutil"
	"os"

	"github.com/borosr/flutter-screenshot/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const (
	exampleCommandValue = `echo "hello world"`
	exampleDeviceName   = "iPhone X"
)

var Init = &cli.Command{
	Name:   "init",
	Action: initCmd,
}

func initCmd(ctx *cli.Context) error {
	return createFile(ctx.String(FlagNameConfig))
}

func createFile(filePath string) error {
	log.Infof("Generating init file to: %s", filePath)

	conf, err := yaml.Marshal(config.Data{
		Cmd: exampleCommandValue,
		Devices: config.Devices{
			IOS: []config.Device{{
				Name: exampleDeviceName,
				Mode: config.ModeLight.String(),
			}},
		},
	})
	if err != nil {
		return err
	}
	log.Debugf("Exported config is: %s", string(conf))

	return ioutil.WriteFile(filePath, conf, os.FileMode(0666))
}
