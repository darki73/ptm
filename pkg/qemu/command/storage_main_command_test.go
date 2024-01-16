package command

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// TestNewMainStorageCommand tests the NewMainStorageCommand function.
func TestNewMainStorageCommand(t *testing.T) {
	identifier := 1
	storage := "local-lvm"
	image := "/etc/ptm/images/ubuntu-22.04-cloudimg-amd64.img"
	scsiCommand := fmt.Sprintf("%s:0,import-from=%s,discard=on", storage, image)

	cmd := NewMainStorageCommand(identifier, storage, image)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewMainStorageCommand did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--scsi0", scsiCommand}) {
		t.Errorf("TestNewMainStorageCommand did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--scsi0", scsiCommand}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
