package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecCommand(application string, args []string) {
	cmd := exec.Command(application, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("There was an error when trying to execute the command, %s", err)
	}
}
