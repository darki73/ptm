package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewNetworkCommand tests the NewNetworkCommand function.
func TestNewNetworkCommand(t *testing.T) {
	identifier := 1
	cmd := NewNetworkCommand(identifier, "virtio", "vmbr0")

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewNetworkCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--net0", "virtio,bridge=vmbr0"}) {
		t.Errorf("TestNewNetworkCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--net0", "virtio,bridge=vmbr0"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
