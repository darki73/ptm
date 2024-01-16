package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"reflect"
	"strconv"
	"testing"
)

// TestNewCloudUsernameCommand tests the NewCloudUsernameCommand function.
func TestNewCloudUsernameCommand(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.SetUsername("test")

	cmd := NewCloudUsernameCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewCloudUsernameCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--ciuser", "test"}) {
		t.Errorf("TestNewCloudUsernameCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--ciuser", "test"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
