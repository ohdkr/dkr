package main

import (
	. "dkr/router"
	"flag"
	"fmt"
	"os"
)

var router = Route

func main() {
	// Prepares app description.
	var Usage = func() {
		fmt.Println("Welcome to docker CLI tool.\r")

		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of dkr:\n")
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

	router(*showVersion, os.Args)
}
