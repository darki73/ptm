package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewNameCommand tests the NewNameCommand function.
func TestNewNameCommand(t *testing.T) {
	identifier := 1
	cmd := NewNameCommand(identifier, "ubuntu-cloudinit")

	if cmd.GetCommand() != qemuCommandCreate || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewNameCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--name", "ubuntu-cloudinit"}) {
		t.Errorf("TestNewNameCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandCreate, strconv.Itoa(identifier), "--name", "ubuntu-cloudinit"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
