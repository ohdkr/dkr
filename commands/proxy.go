package proxy

import (
	"fmt"
	"os"
	"os/exec"
)

// Raw proxy to docker or docker-compose.
func Proxy() {
	var args []string
	var secondArg string

	application := "docker"

	if len(os.Args) > 1 {
		args = os.Args[1:]
		secondArg = os.Args[1]
	}

	if secondArg == "c" {
		application = "docker-compose"
		if len(os.Args) > 2 {
			args = os.Args[2:]
		} else {
			args = nil
		}
	}

	fmt.Printf("Calling %s with %v\n", application, args)
	cmd := exec.Command(application, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}
