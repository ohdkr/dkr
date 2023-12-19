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
		fmt.Println("Welcome to docker CLI tool.")

		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of dkr:")
		if err != nil {
			fmt.Printf("Error when trying to format the output string, %s\n", err)
		}

		flag.PrintDefaults()
		// Subcommands
		fmt.Println("aliases:")
		fmt.Println("  follow - Follows all running container logs")
		fmt.Println("  sh CONTAINER_NAME - Jumps into running container sh.")
		fmt.Println("  bash CONTAINER_NAME - Jumps into running container bash.")
		fmt.Println("  killall - Kills all active container. Equivalent: docker kill $(docker ps -q).")
		fmt.Println("  cleanup - Removes all containers and volumes. Equivalent: docker rm $(docker ps -a -q) & docker rmi $(docker images -q).")
		fmt.Println("  nuke - Removes everything. Alias to cleanup and docker system prune --volumes -f. ")
	}
	flag.Usage = Usage

	showVersion := flag.Bool("version", false, "Prints version")
	flag.Parse()

	router(*showVersion, os.Args)
}
