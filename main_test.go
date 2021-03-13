package main

import (
	"os"
	"testing"

	"github.com/borosr/flutter-screenshot/src/cmd"
)

func TestInitCmdVerbose(t *testing.T) {
	os.Args = []string{"app", "--"+cmd.FlagNameVerbose, cmd.Init.Name}
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()
	main()
}
