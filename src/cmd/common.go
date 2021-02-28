package cmd

import "github.com/urfave/cli/v2"

// FlagNameVerbose the full name of the verbose flag
// FlagNameConfig the full name of the config flag
// DefaultConfigFileName the default value of the config file
const (
	FlagNameVerbose = "verbose"
	FlagNameConfig  = "config"

	DefaultConfigFileName = "screenshots.yaml"
)

// FlagVerbose the cli tools verbose flag
// FlagConfig the cli tools config flag
var (
	FlagVerbose = &cli.BoolFlag{
		Name:     FlagNameVerbose,
		Aliases:  []string{"v"},
		Required: false,
		Value:    false,
	}
	FlagConfig = &cli.StringFlag{
		Name:     FlagNameConfig,
		Aliases:  []string{"c"},
		Required: false,
		Value:    DefaultConfigFileName,
	}
)
