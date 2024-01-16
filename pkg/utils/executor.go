package utils

import (
	"os/exec"
)

// ExecuteCommand executes a command and returns the output.
func ExecuteCommand(command string, arguments ...string) (string, error) {
	cmd := exec.Command(command, arguments...)
	output, err := cmd.Output()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}
