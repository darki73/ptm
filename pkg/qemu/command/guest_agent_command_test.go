package command

import (
	"reflect"
	"strconv"
	"testing"
)

// TestNewGuestAgentCommandWithAllEnabled tests the NewGuestAgentCommand function with all options enabled.
func TestNewGuestAgentCommandWithAllEnabled(t *testing.T) {
	identifier := 1
	cmd := NewGuestAgentCommand(identifier, true, true)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewGuestAgentCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--agent", "enabled=1,fstrim_cloned_disks=1"}) {
		t.Errorf("TestNewGuestAgentCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--agent", "enabled=1,fstrim_cloned_disks=1"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}

// TestNewGuestAgentCommandWithAllDisabled tests the NewGuestAgentCommand function with all options disabled.
func TestNewGuestAgentCommandWithAllDisabled(t *testing.T) {
	identifier := 1
	cmd := NewGuestAgentCommand(identifier, false, false)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewGuestAgentCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--agent", "enabled=0"}) {
		t.Errorf("TestNewGuestAgentCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--agent", "enabled=0"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}

// TestNewGuestAgentCommandWithAgentEnabledAndTrimDisabled tests the NewGuestAgentCommand function with agent enabled and trim disabled.
func TestNewGuestAgentCommandWithAgentEnabledAndTrimDisabled(t *testing.T) {
	identifier := 1
	cmd := NewGuestAgentCommand(identifier, true, false)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewGuestAgentCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--agent", "enabled=1"}) {
		t.Errorf("TestNewGuestAgentCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--agent", "enabled=1"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}

// TestNewGuestAgentCommandWithAgentDisabledAndTrimEnabled tests the NewGuestAgentCommand function with agent disabled and trim enabled.
func TestNewGuestAgentCommandWithAgentDisabledAndTrimEnabled(t *testing.T) {
	identifier := 1
	cmd := NewGuestAgentCommand(identifier, false, true)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewGuestAgentCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--agent", "enabled=0,fstrim_cloned_disks=1"}) {
		t.Errorf("TestNewGuestAgentCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--agent", "enabled=0,fstrim_cloned_disks=1"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
