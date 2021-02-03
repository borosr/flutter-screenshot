package ios

import (
	"github.com/borosr/flutter-screenshot/pkg/exec"
)

var execute exec.CommandExecutor = exec.Command

func mockExecute(e exec.Executable) exec.CommandExecutor {
	return func(_ string, _ ...string) exec.Executable {
		return e
	}
}
