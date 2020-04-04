package commands

import (
	"fmt"
	"os"
)

func handleSH(shArgs []string) {
	if len(shArgs) == 0 {
		fmt.Printf("Please provide container name. E.x dkr sh SOME_NAME\n")
		os.Exit(1)
	}
	container := shArgs[0]
	ExecCommand("docker", []string{"exec", "-it", container, "/bin/sh"})
	os.Exit(0)
}

func DetectAndCallAliases() {
	// No arguments after `dkr` or `dkr c` called
	if len(os.Args) == 1 ||
		len(os.Args) >= 2 &&
			os.Args[1] == "c" {
		return
	}

	args := os.Args[1:]
	alias := args[0]
	var rest []string

	if len(args) > 1 {
		rest = args[1:]
	}

	switch alias {
	case "sh":
		handleSH(rest)
	}
}
