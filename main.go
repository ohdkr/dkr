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
		fmt.Println("Welcome to docker CLI tool.\n")

		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of dkr:\n")
		if err != nil {
			fmt.Printf("Error when trying to format the output string, %s", err)
		}

		flag.PrintDefaults()
		// Subcommands
		fmt.Println("aliases:\n")
		fmt.Println("  sh CONTAINER_NAME - Jumps into running container sh.\n")
		fmt.Println("  bash CONTAINER_NAME - Jumps into running container bash.\n")
		fmt.Println("  killall - Kills all active container. Equivalent: docker kill $(docker ps -q).\n")
		fmt.Println("  cleanup - Removes all containers and volumes. Equivalent: docker rm $(docker ps -a -q) & docker rmi $(docker images -q).\n")
		fmt.Println("  nuke - Removes everything. Alias to cleanup and docker system prune --volumes -f. \n")
	}
	flag.Usage = Usage

	showVersion := flag.Bool("version", false, "Prints version")
	flag.Parse()

	router(*showVersion, os.Args)
}
