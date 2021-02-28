package main

import (
	"os"

	"github.com/borosr/flutter-screenshot/src/cmd"
	"github.com/borosr/flutter-screenshot/src/executor"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Commands: []*cli.Command{
		cmd.Init,
		cmd.Doctor,
	},
	Flags: []cli.Flag{
		cmd.FlagVerbose,
		cmd.FlagConfig,
	},
	Before: func(ctx *cli.Context) error {
		if ctx.Bool(cmd.FlagNameVerbose) {
			log.SetLevel(log.DebugLevel)
		}

		return nil
	},
	Action: func(ctx *cli.Context) error {
		return executor.Run(ctx.String(cmd.FlagNameConfig))
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
