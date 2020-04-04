package commands

import (
	"os"
	"os/exec"
)

func ExecCommand(application string, args []string) {
	cmd := exec.Command(application, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}
