package commands

import (
	"fmt"
	"os"
)

// Raw proxy to docker or docker-compose.
func Proxy() {
	var args []string = nil
	application := "docker"
	rawArgs := os.Args

	if len(rawArgs) > 1 {
		args = rawArgs[1:]
		if rawArgs[1] == "c" {
			application = "docker-compose"
			if len(rawArgs) > 2 {
				args = rawArgs[2:]
			}
		}
	}

	fmt.Printf("Calling %s with %v\n", application, args)
	ExecCommand(application, args)
}
