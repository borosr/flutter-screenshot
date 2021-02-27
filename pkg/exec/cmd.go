package exec

import (
	"io"
	"os/exec"
)

// Cmd is doing the same as the os/exec's Cmd,
// but this Cmd implements Executable interface
// to make mocking possible
type Cmd struct {
	*exec.Cmd
}

// CommandExecutor defines a function for Cmd instance creation
type CommandExecutor func(string, ...string) Executable

// Executable defines all methods which are needed at external calls
type Executable interface {
	Run() error
	Start() error
	String() string
	Stdout(io.Writer)
}

// Command is the basic implementation of Cmd creation
func Command(name string, args ...string) Executable {
	return &Cmd{
		Cmd: exec.Command(name, args...),
	}
}

// Stdout sets os/exec.Cmd's Stdout field
func (c *Cmd) Stdout(f io.Writer) {
	c.Cmd.Stdout = f
}
