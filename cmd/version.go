package cmd

import "fmt"

const POLYGO_VERSION = "v0.0.3"

func Version() {
	fmt.Println("polygo version:", POLYGO_VERSION)
}
