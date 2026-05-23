package main

import (
	"os"

	"github.com/hashicorp/boundary/internal/cmd"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"dev"}
	}
	os.Exit(cmd.Run(args))
}