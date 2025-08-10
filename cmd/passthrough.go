package cmd

import (
	"os"
	"os/exec"
)

func PassThroughCommand() error {
	args := os.Args[1:]

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
