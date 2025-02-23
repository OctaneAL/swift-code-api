package main

import (
	"os"

	"github.com/OctaneAL/swift-code-api/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
