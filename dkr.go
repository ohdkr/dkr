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
		println("Welcome to docker CLi tool.")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Usage = Usage

	showVersion := flag.Bool("version", false, "Prints version")
	flag.Parse()

	if *showVersion {
		println(version)
		os.Exit(0)
	}

	// No known command found, proceeds as a raw proxy.
	Proxy()
}
