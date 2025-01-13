package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/argandtech/gatekeeper/internal/cli"
)

var (
	version string
)

func main() {
	if err := cli.Run(strings.TrimSuffix(version, "\n"), os.Args...); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute command: %v", err)
		os.Exit(1)
	}
}
