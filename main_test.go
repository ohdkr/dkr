package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

var executedCommands []string

func fakeExecCommand(command string, args []string) {
	cmd := []string{command}
	cmd = append(cmd, args...)
	executedCommands = append(executedCommands, strings.Join(cmd, " "))
}

var lastExitCode int

func fakePrexit(code int) {
	lastExitCode = code
}

func runWithArgs(t *testing.T, args []string) string {
	oldArgs := os.Args
	os.Args = args
	defer func() { os.Args = oldArgs }()

	execCommand = fakeExecCommand
	prexiter = fakePrexit

	out := captureStdout(main)

	return out
}

func TestVersion(t *testing.T) {
	out := runWithArgs(t, []string{"dkr", "--version"})

	expectedOut := "Dkr version: " + version + "\n"
	assert.Equal(t, expectedOut, out)
	assert.Equal(t, "docker --version", executedCommands[0])
	assert.Equal(t, "docker-compose --version", executedCommands[1])
	assert.Equal(t, 0, lastExitCode)
}

// Source https://gist.github.com/mindscratch/0faa78bd3c0005d080bf
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
