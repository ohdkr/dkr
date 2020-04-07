package main

import (
	"flag"
	"fmt"
	"os"

	. "dkr/commands"
)

var version = "0.1.0"

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
