package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var calledShowVersion bool
var calledOsArgs []string

func fakeRouter(showVersion bool, osArgs []string) {
	calledShowVersion = showVersion
	calledOsArgs = osArgs
}

func runWithArgs(t *testing.T, args []string) {
	oldArgs := os.Args
	os.Args = args
	defer func() { os.Args = oldArgs }()
	main()
}

func TestVersion(t *testing.T) {
	router = fakeRouter
	runWithArgs(t, []string{"dkr", "--version"})

	assert.Equal(t, true, calledShowVersion)
	assert.Equal(t, "dkr", calledOsArgs[0])
	assert.Equal(t, "--version", calledOsArgs[1])
}
