package command

import (
	"reflect"
	"testing"
)

// TestNewCommand tests the NewCommand function.
func TestNewCommand(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(qemuCommandCreate, identifier, arguments...)

	if cmd.GetCommand() != qemuCommandCreate || cmd.GetIdentifier() != identifier {
		t.Errorf("NewCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("NewCommand did not set arguments correctly")
	}
}

// TestNewCreateCommand tests the NewCreateCommand function.
func TestNewCreateCommand(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCreateCommand(identifier, arguments...)

	if cmd.GetCommand() != qemuCommandCreate || cmd.GetIdentifier() != identifier {
		t.Errorf("NewCreateCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("NewCreateCommand did not set arguments correctly")
	}
}

// TestNewSetCommand tests the NewSetCommand function.
func TestNewSetCommand(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewSetCommand(identifier, arguments...)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("NewSetCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("NewSetCommand did not set arguments correctly")
	}
}

// TestNewDiskCommand tests the NewDiskCommand function.
func TestNewDiskCommand(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewDiskCommand(identifier, arguments...)

	if cmd.GetCommand() != qemuCommandDisk || cmd.GetIdentifier() != identifier {
		t.Errorf("NewDiskCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("NewDiskCommand did not set arguments correctly")
	}
}

// TestNewTemplateCommand tests the NewTemplateCommand function.
func TestNewTemplateCommand(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewTemplateCommand(identifier, arguments...)

	if cmd.GetCommand() != qemuCommandTemplate || cmd.GetIdentifier() != identifier {
		t.Errorf("NewTemplateCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("NewTemplateCommand did not set arguments correctly")
	}
}

// TestGetOrder tests the GetOrder function.
func TestGetOrder(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(qemuCommandCreate, identifier, arguments...)

	if cmd.GetOrder() != 0 {
		t.Errorf("GetOrder did not return 0")
	}
}

// TestSetOrder tests the SetOrder function.
func TestSetOrder(t *testing.T) {
	identifier := 1
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(qemuCommandCreate, identifier, arguments...)

	cmd.SetOrder(1)
	if cmd.GetOrder() != 1 {
		t.Errorf("SetOrder did not set order correctly")
	}
}
