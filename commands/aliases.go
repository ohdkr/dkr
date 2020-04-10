package commands

import (
	"fmt"
	"strings"
)

var prexiter = Prexit

func handleJumpIntoTerminal(mode string, shArgs []string) int {
	if len(shArgs) == 0 {
		fmt.Printf("Please provide container name. E.x dkr sh SOME_NAME\n")
		return 1
	}
	container := shArgs[0]
	ExecCommand("docker", []string{"exec", "-it", container, mode})
	return 0
}

func handleKillAll() int {
	ids := string(ReturnCommand("docker", []string{"ps", "-q"}))
	idsArr := strings.Split(strings.Trim(ids, "\n"), "\n")

	if idsArr[0] == "" {
		fmt.Print("There's nothing to kill.")
		return 0
	}

	for _, element := range idsArr {
		ExecCommand("docker", []string{"kill", element})
	}

	return 0
}

func handleCleanup() int {
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
	return 0
}

func DetectAndCallAliases(osArgs []string) (bool, int) {
	// No arguments after `dkr` or `dkr c` called
	if len(osArgs) == 1 ||
		len(osArgs) >= 2 &&
			osArgs[1] == "c" {
		return false, 0
	}

	args := osArgs[1:]
	alias := args[0]
	var rest []string

	if len(args) > 1 {
		rest = args[1:]
	}

	switch alias {
	case "sh":
		return true, handleJumpIntoTerminal("/bin/sh", rest)
	case "bash":
		return true, handleJumpIntoTerminal("/bin/bash", rest)
	case "killall":
		return true, handleKillAll()
	case "cleanup":
		return true, handleCleanup()
	}

	return false, 0
}
