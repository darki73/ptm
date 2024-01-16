package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewResourcesCommand tests the NewResourcesCommand function.
func TestNewResourcesCommand(t *testing.T) {
	identifier := 1
	cmd := NewResourcesCommand(identifier, 4, 8192, "host")

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewResourcesCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--cores", "4", "--memory", "8192", "--cpu", "host"}) {
		t.Errorf("TestNewResourcesCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--cores", "4", "--memory", "8192", "--cpu", "host"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
