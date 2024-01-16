package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"reflect"
	"strconv"
	"testing"
)

// TestNewCloudKeysCommand tests the NewCloudKeysCommand function.
func TestNewCloudKeysCommand(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.SetSSHKeysTemporaryFilePath("/tmp/ptm-ssh-keys")

	cmd := NewCloudKeysCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewCloudKeysCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--sshkey", "/tmp/ptm-ssh-keys"}) {
		t.Errorf("TestNewCloudKeysCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--sshkey", "/tmp/ptm-ssh-keys"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
