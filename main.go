package main

import (
	"os"

	"github.com/gabe565/ansi2txt/cmd"
)

var version = "beta"

func main() {
	if err := cmd.New(cmd.WithVersion(version)).Execute(); err != nil {
		os.Exit(1)
	}
}
