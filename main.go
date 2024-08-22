package main

import (
	"os"

	"github.com/gabe565/ansi2txt/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
