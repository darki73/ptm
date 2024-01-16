package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewResizeCommand tests the NewResizeCommand function.
func TestNewResizeCommand(t *testing.T) {
	identifier := 1
	cmd := NewResizeCommand(identifier, "scsi0", "4G")

	if cmd.GetCommand() != qemuCommandDisk || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewResizeCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"resize", "%VM_ID%", "scsi0", "4G"}) {
		t.Errorf("TestNewResizeCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandDisk, "resize", strconv.Itoa(identifier), "scsi0", "4G"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
