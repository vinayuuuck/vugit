package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Check if any arguments were passed
	if len(os.Args) < 2 {
		fmt.Println("Usage: vugit <git-args>")
		os.Exit(1)
	}

	// All args after the program name
	args := os.Args[1:]

	// Prepare git command with args
	cmd := exec.Command("git", args...)

	// Connect command output directly to terminal stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running git command: %v\n", err)
		os.Exit(1)
	}
}
