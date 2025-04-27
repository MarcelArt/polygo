package main

import (
	"os"

	"github.com/MarcelArt/polygo/cmd"
)

func main() {
	args := os.Args
	argsLength := len(args)
	if argsLength > 1 {
		cmd.Manager(args)
	} else {
		cmd.Polygo()
	}
}
