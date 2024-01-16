package command

import (
	"reflect"
	"testing"
)

// TestNewUpdateCommand tests the NewUpdateCommand function.
func TestNewUpdateCommand(t *testing.T) {
	cmd := NewUpdateCommand(imagePathForTesting)

	expectedCommand := []string{"-a", imagePathForTesting, "--update"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expectedCommand) {
		t.Errorf("Expected command to be %v, but got %v", expectedCommand, result)
	}
}
