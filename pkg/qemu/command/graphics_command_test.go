package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewGraphicsCommand tests the NewGraphicsCommand function.
func TestNewGraphicsCommand(t *testing.T) {
	identifier := 1
	cmd := NewGraphicsCommand(identifier)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewGraphicsCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--serial0", "socket", "--vga", "serial0"}) {
		t.Errorf("TestNewGraphicsCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--serial0", "socket", "--vga", "serial0"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
