package command

import (
	"reflect"
	"testing"
)

// TestNewUploadCommand tests the NewUploadCommand function.
func TestNewUploadCommand(t *testing.T) {
	source := "/tmp/source.file"
	target := "/tmp/target.file"
	cmd := NewUploadCommand(imagePathForTesting, source, target)

	expectedCommand := []string{"-a", imagePathForTesting, "--upload", "/tmp/source.file:/tmp/target.file"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expectedCommand) {
		t.Errorf("Expected command to be %v, but got %v", expectedCommand, result)
	}
}
