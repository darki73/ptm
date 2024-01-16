package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewCloudStorageCommand tests the NewCloudStorageCommand function.
func TestNewCloudStorageCommand(t *testing.T) {
	identifier := 1
	cmd := NewCloudStorageCommand(identifier, "local-lvm")

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewCloudStorageCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--ide2", "local-lvm:cloudinit"}) {
		t.Errorf("TestNewCloudStorageCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--ide2", "local-lvm:cloudinit"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
