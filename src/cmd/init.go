package cmd

import (
	"io/ioutil"
	"os"

	"github.com/borosr/flutter-screenshot/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var Init = &cli.Command{
	Name: "init",
	Action: func(ctx *cli.Context) error {
		filePath := ctx.String(FlagNameConfig)
		log.Infof("Generating init file to: %s", filePath)

		conf, err := yaml.Marshal(config.Data{
			Cmd: `echo "hello world"`,
			Devices: config.Devices{
				IOS: []config.Device{{
					Name: "iPhone X",
					Mode: config.ModeLight.String(),
				}},
			},
		})
		if err != nil {
			return err
		}
		log.Debugf("Exported config is: %s", string(conf))

		return ioutil.WriteFile(filePath, conf, os.FileMode(0666))
	},
}
