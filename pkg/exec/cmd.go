package exec

import (
	"os"
	"os/exec"
)

type Cmd struct {
	*exec.Cmd
}

type CommandExecutor func(string, ...string) Executable

type Executable interface {
	Run() error
	String() string
	Stdout(*os.File)
}

func Command(name string, args ...string) Executable {
	return &Cmd{
		Cmd: exec.Command(name, args...),
	}
}

func (c *Cmd) Stdout(f *os.File) {
	c.Cmd.Stdout = f
}
