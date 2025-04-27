package cmd

import (
	"fmt"
	"os"
)

func Manager(args []string) {
	argsLength := len(args)
	fn := args[1]
	switch fn {
	case "add":
		if argsLength < 3 {
			fmt.Println("Please input model name")
			os.Exit(1)
		}
		Scaffolder(args[2])
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}

}
