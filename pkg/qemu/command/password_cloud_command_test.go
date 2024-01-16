package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"reflect"
	"strconv"
	"testing"
)

// TestNewCloudPasswordCommand tests the NewCloudPasswordCommand function.
func TestNewCloudPasswordCommand(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.SetPassword("password")

	cmd := NewCloudPasswordCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewCloudPasswordCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--cipassword", "password"}) {
		t.Errorf("TestNewCloudPasswordCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--cipassword", "password"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
