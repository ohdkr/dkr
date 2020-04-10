package commands

import (
	"fmt"
)

// Raw proxy to docker or docker-compose.
func Proxy(osArgs []string) {
	var args []string = nil
	application := "docker"
	rawArgs := osArgs

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
