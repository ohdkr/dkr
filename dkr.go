package main

import (
	. "dkr/commands"
	"flag"
	"fmt"
	"os"
)

var version = "0.1.0"

func main() {
	// Prepares app description.
	var Usage = func() {
		println("Welcome to docker CLi tool.\r")

		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		// Subcommands
		println("aliases:\r")
		println("  sh CONTAINER_NAME - Jumps into running container sh.\r")
		println("  bash CONTAINER_NAME - Jumps into running container bash.\r")
	}
	flag.Usage = Usage

	showVersion := flag.Bool("version", false, "Prints version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Dkr version: %s\n", version)
		ExecCommand("docker", []string{"--version"})
		ExecCommand("docker-compose", []string{"--version"})
		os.Exit(0)
	}

	// Check if passed known aliases. If yes, this will exit inside.
	DetectAndCallAliases()
	// No known command found, proceeds as a raw proxy.
	Proxy()
}
