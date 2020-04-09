package main

import (
	"flag"
	"fmt"
	"os"

	. "dkr/commands"
)

var version = "0.1.0"
var execCommand = ExecCommand
var prexiter = Prexit

func main() {
	// Prepares app description.
	var Usage = func() {
		fmt.Println("Welcome to docker CLI tool.\r")

		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		if err != nil {
			fmt.Printf("Error when trying to format the output string, %s", err)
		}

		flag.PrintDefaults()
		// Subcommands
		fmt.Println("aliases:\r")
		fmt.Println("  sh CONTAINER_NAME - Jumps into running container sh.\r")
		fmt.Println("  bash CONTAINER_NAME - Jumps into running container bash.\r")
		fmt.Println("  killall - Kills all active container. Equivalent: docker kill $(docker ps -q).\r")
		fmt.Println("  cleanup - Removes all containers and volumes. Equivalent: docker rm $(docker ps -a -q) & docker rmi $(docker images -q).\r")
	}
	flag.Usage = Usage

	showVersion := flag.Bool("version", false, "Prints version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Dkr version: %s\n", version)
		execCommand("docker", []string{"--version"})
		execCommand("docker-compose", []string{"--version"})
		prexiter(0)
		return
	}

	// Check if passed known aliases. If yes, this will exit inside.
	shouldFinish, code := DetectAndCallAliases()
	if shouldFinish {
		prexiter(code)
		return
	}
	// No known command found, proceeds as a raw proxy.
	Proxy()
	prexiter(0)
}
