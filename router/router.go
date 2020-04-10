package router

import (
	. "dkr/commands"
	"fmt"
)

var execCommand = ExecCommand
var prexiter = Prexit
var proxy = Proxy
var detectAndCallAliases = DetectAndCallAliases
var version = "0.1.0"

func Route(showVersion bool, osArgs []string) {
	if showVersion {
		fmt.Printf("Dkr version: %s\n", version)
		execCommand("docker", []string{"--version"})
		execCommand("docker-compose", []string{"--version"})
		prexiter(0)
		return
	}

	// Check if passed known aliases. If yes, this will exit inside.
	shouldFinish, code := detectAndCallAliases(osArgs)
	if shouldFinish {
		prexiter(code)
		return
	}
	// No known command found, proceeds as a raw proxy.
	proxy(osArgs)
	prexiter(0)
}
