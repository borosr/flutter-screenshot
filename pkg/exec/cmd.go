package exec

import "os/exec"

type Cmd = *exec.Cmd

type CommandExecutor func(string, ...string) Executable

type Executable interface {
	Run() error
	String() string
}

func Command(name string, args ...string) Executable {
	return exec.Command(name, args...)
}
