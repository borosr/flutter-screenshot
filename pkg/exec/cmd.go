package exec

import (
	"io"
	"os/exec"
)

type Cmd struct {
	*exec.Cmd
}

type CommandExecutor func(string, ...string) Executable

type Executable interface {
	Run() error
	Start() error
	String() string
	Stdout(io.Writer)
}

func Command(name string, args ...string) Executable {
	return &Cmd{
		Cmd: exec.Command(name, args...),
	}
}

func (c *Cmd) Stdout(f io.Writer) {
	c.Cmd.Stdout = f
}
