package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ReturnCommand(application string, args []string) []byte {
	cmd := exec.Command(application, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return out
}

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

func Prexit(code int) {
	os.Exit(code)
}
