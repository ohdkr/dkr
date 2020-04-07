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

	// Not called just `dkr`
	if len(rawArgs) > 1 {
		// Assign all args after `dkr`
		args = rawArgs[1:]
		// Second argument is c => `dkr c`
		if rawArgs[1] == "c" {
			application = "docker-compose"
			// There are even more arguments, assigning them.
			if len(rawArgs) > 2 {
				args = rawArgs[2:]
			}
		}
	}

	fmt.Printf("Calling %s with %v\n", application, args)
	ExecCommand(application, args)
}
