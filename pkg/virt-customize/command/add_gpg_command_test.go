package command

import (
	"fmt"
	"reflect"
	"testing"
)

// TestNewAddGPGCommand tests the NewAddGPGCommand function.
func TestNewAddGPGCommand(t *testing.T) {
	cmd := NewAddGPGCommand(imagePathForTesting, repositoryConfigurationForTesting)

	expectedCommand := []string{
		"-a",
		imagePathForTesting,
		"--run-command",
		fmt.Sprintf("curl -fsSL %s | gpg --dearmor -o %s", repositoryConfigurationForTesting.GetGPG(), repositoryConfigurationForTesting.GetKeyFullPath()),
	}

	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expectedCommand) {
		t.Errorf("Expected command to be %v, but got %v", expectedCommand, result)
	}
}
