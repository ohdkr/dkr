package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var command = exec.Command

func ReturnCommand(application string, args []string) []byte {
	cmd := command(application, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func ExecCommand(application string, args []string) {
	cmd := command(application, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("There was an error when trying to execute the command, %s", err)
	}
}

func Prexit(code int) {
	os.Exit(code)
}
