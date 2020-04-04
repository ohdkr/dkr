package proxy

import (
	"fmt"
	"os"
	"os/exec"
)

// Raw proxy to docker or docker-compose.
func Proxy() {
	application := "docker"
	args := os.Args[1:]
	secondArg := os.Args[1]

	if secondArg == "c" {
		application = "docker-compose"
		args = os.Args[2:]
	}

	fmt.Printf("Calling docker with %v\n", args)
	cmd := exec.Command(application, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Run()
}
