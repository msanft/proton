package main

import (
	"fmt"
	"os/exec"

	"github.com/msanft/proton/daemon"
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

	daemonConn, err := daemon.NewConn(stdout, stdin)
	if err != nil {
		panic(err)
	}

	storePaths := []string{
		"/nix/store/yda60ngw0yiknd3xx1yrszkv2s6askrf-libcap-2.69-man",
		"/nix/store/yrik1ppdagk2y6pn3yaly2lz90ll92v2-cargo-package-libc-0.2.foo",
		// "/abc", // TODO: Add the error handling for this
	}

	for _, path := range storePaths {
		valid, err := daemonConn.IsValidPath(path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s is valid: %t\n", path, valid)
	}

	b := make([]byte, 1024)
	n, err := stdout.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", b[:n])
}
