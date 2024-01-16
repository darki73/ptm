package command

import (
	"reflect"
	"testing"
)

// TestNewInstallCommand tests the NewInstallCommand function.
func TestNewInstallCommand(t *testing.T) {
	packages := []string{"package1", "package2"}
	cmd := NewInstallCommand(imagePathForTesting, packages)

	expectedCommand := []string{"-a", imagePathForTesting, "--install", "package1,package2"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expectedCommand) {
		t.Errorf("Expected command to be %v, but got %v", expectedCommand, result)
	}
}
