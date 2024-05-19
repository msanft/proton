package main

import (
	"os/exec"

	"github.com/msanft/proton/internal/daemon"
)

func main() {
	c := exec.Command("nix-daemon", "--stdio")
	stdin, err := c.StdinPipe()
	if err != nil {
		panic(err)
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := c.Start(); err != nil {
		panic(err)
	}

	_, err = daemon.NewConn(stdout, stdin)
	if err != nil {
		panic(err)
	}
}
