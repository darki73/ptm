package command

import (
	"fmt"
	"reflect"
	"testing"
)

// TestNewAddRepositoryCommand tests the NewAddRepositoryCommand function.
func TestNewAddRepositoryCommand(t *testing.T) {
	cmd := NewAddRepositoryCommand(imagePathForTesting, repositoryConfigurationForTesting)

	expectedCommand := []string{
		"-a",
		imagePathForTesting,
		"--run-command",
		fmt.Sprintf("echo \"%s\" > %s", repositoryConfigurationForTesting.GetConfigurationContents(), repositoryConfigurationForTesting.GetConfigurationFullPath()),
	}

	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expectedCommand) {
		t.Errorf("Expected command to be %v, but got %v", expectedCommand, result)
	}
}
