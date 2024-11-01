package utils

import (
	"os"
	"os/exec"
)

type Executor interface {
	Execute(arg string, args ...string) error
}

type OsExecutor struct {
}

func (o OsExecutor) Execute(arg string, args ...string) error {
	cmd := exec.Command(arg, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
