// main.go
package main

import (
	"fmt"
	"os"

	"vugit/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vugit <git-command> [args...]")
		os.Exit(1)
	}

	subcommand := os.Args[1]
	var err error

	switch subcommand {
	case "log":
		err = cmd.HandleLogCommand()
	default:
		err = cmd.PassThroughCommand()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
