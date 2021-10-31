package main

import (
	"os"

	"github.com/rafaelcr1dev/Golang-Blockchain.git/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
