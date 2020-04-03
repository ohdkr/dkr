package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	const version = "0.0.1"

	var Usage = func() {
		println("Welcome to docker CLi tool.")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Usage = Usage

	var showVersion = flag.Bool("version", false, "Prints version")
	flag.Parse()

	if *showVersion {
		println(version)
		os.Exit(0)
	}
}
