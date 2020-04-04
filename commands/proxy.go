package commands

import (
	"fmt"
	"os"
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
	ExecCommand(application, args)
}
