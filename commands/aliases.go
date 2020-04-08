package commands

import (
	"fmt"
	"os"
	"strings"
)

func handleJumpIntoTerminal(mode string, shArgs []string) {
	if len(shArgs) == 0 {
		fmt.Printf("Please provide container name. E.x dkr sh SOME_NAME\n")
		os.Exit(1)
	}
	container := shArgs[0]
	ExecCommand("docker", []string{"exec", "-it", container, mode})
	os.Exit(0)
}

func handleKillAll() {
	ids := string(ReturnCommand("docker", []string{"ps", "-q"}))
	idsArr := strings.Split(strings.Trim(ids, "\n"), "\n")

	if idsArr[0] == "" {
		fmt.Print("There's nothing to kill.")
		os.Exit(0)
	}

	for _, element := range idsArr {
		ExecCommand("docker", []string{"kill", element})
	}

	os.Exit(0)
}

func handleCleanup() {
	ids := string(ReturnCommand("docker", []string{"ps", "-aq"}))
	idsArr := strings.Split(strings.Trim(ids, "\n"), "\n")

	if idsArr[0] != "" {
		for _, element := range idsArr {
			ExecCommand("docker", []string{"rm", element})
		}
	} else {
		fmt.Print("No containers to remove.\n")
	}

	imagesIds := string(ReturnCommand("docker", []string{"images", "-q"}))
	imagesIdsArr := strings.Split(strings.Trim(imagesIds, "\n"), "\n")

	if imagesIdsArr[0] != "" {
		for _, element := range imagesIdsArr {
			ExecCommand("docker", []string{"rmi", "-f", element})
		}
	} else {
		fmt.Print("No volumes to remove.\n")
	}
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
		handleJumpIntoTerminal("/bin/sh", rest)
	case "bash":
		handleJumpIntoTerminal("/bin/bash", rest)
	case "killall":
		handleKillAll()
	case "cleanup":
		handleCleanup()
	}
}
