package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewBootOrderCommand tests the NewBootOrderCommand function.
func TestNewBootOrderCommand(t *testing.T) {
	identifier := 1
	cmd := NewBootOrderCommand(identifier, "scsi0", "virtio-scsi-single")

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewBootOrderCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--boot", "order=scsi0", "--scsihw", "virtio-scsi-single"}) {
		t.Errorf("TestNewBootOrderCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--boot", "order=scsi0", "--scsihw", "virtio-scsi-single"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
