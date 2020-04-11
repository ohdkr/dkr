package router

import (
	. "dkr/commands"
	"fmt"
	"os"
)

var version = "0.2.0"

func Route(showVersion bool, osArgs []string) {
	if showVersion {
		fmt.Printf("Dkr version: %s\n", version)
		ExecCommand("docker", []string{"--version"})
		ExecCommand("docker-compose", []string{"--version"})
		os.Exit(0)
		return
	}

	// Check if passed known aliases. If yes, this will exit inside.
	shouldFinish, code := DetectAndCallAliases(osArgs)
	if shouldFinish {
		os.Exit(code)
		return
	}
	// No known command found, proceeds as a raw proxy.
	Proxy(osArgs)
	os.Exit(0)
}
