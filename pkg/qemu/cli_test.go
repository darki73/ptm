package qemu

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"github.com/darki73/ptm/pkg/qemu/command"
	"os"
	"reflect"
	"testing"
)

// TestNewCommandLineInterface tests the NewCommandLineInterface function.
func TestNewCommandLineInterface(t *testing.T) {
	config := &Qemu{} // Create a Qemu object with necessary mock data
	cli := NewCommandLineInterface(config)

	if cli == nil {
		t.Fatal("NewCommandLineInterface() returned nil")
	}

	if cli.configuration != config {
		t.Error("NewCommandLineInterface() did not set configuration correctly")
	}
}

// TestAddCommand tests the addCommand method.
func TestAddCommand(t *testing.T) {
	cli := NewCommandLineInterface(&Qemu{})
	commandToAdd := &command.Command{}

	cli.addCommand(commandToAdd)

	if len(cli.commands) != 1 {
		t.Errorf("Expected 1 command, got %d", len(cli.commands))
	}
	if cli.commands[1] != commandToAdd {
		t.Error("addCommand did not add the command correctly")
	}
}

// TestAddCleanupFunction tests the addCleanupFunction method.
func TestAddCleanupFunction(t *testing.T) {
	cli := NewCommandLineInterface(&Qemu{})
	funcToAdd := func() error { return nil }

	cli.addCleanupFunction(funcToAdd)

	if len(cli.cleanupFunctions) != 1 {
		t.Errorf("Expected 1 cleanup function, got %d", len(cli.cleanupFunctions))
	}
	if reflect.ValueOf(cli.cleanupFunctions[0]).Pointer() != reflect.ValueOf(funcToAdd).Pointer() {
		t.Error("addCleanupFunction did not add the function correctly")
	}
}

// TestCreateTemporarySshKeysFile tests the createTemporarySshKeysFile method.
func TestCreateTemporarySshKeysFile(t *testing.T) {
	cli := NewCommandLineInterface(&Qemu{})
	cloudInit := &ci.CloudInit{} // Mock necessary CloudInit data
	// Set a temporary file path for the test
	cloudInit.SetSSHKeysTemporaryFilePath("test_ssh_keys.temp")

	// Cleanup after the test
	defer os.Remove("test_ssh_keys.temp")

	fileHandle, err := cli.createTemporarySshKeysFile(cloudInit)
	if err != nil {
		t.Fatalf("createTemporarySshKeysFile() returned error: %v", err)
	}
	if fileHandle == nil {
		t.Fatal("createTemporarySshKeysFile() returned a nil file handle")
	}

	// Ensure the file exists
	if _, err := os.Stat(cloudInit.GetSSHKeysTemporaryFilePath()); os.IsNotExist(err) {
		t.Error("Temporary SSH keys file was not created")
	}
}

// TestCleanup tests the cleanup method.
func TestCleanup(t *testing.T) {
	cli := NewCommandLineInterface(&Qemu{})
	called := false
	cli.addCleanupFunction(func() error {
		called = true
		return nil
	})

	err := cli.cleanup()
	if err != nil {
		t.Errorf("cleanup() returned error: %v", err)
	}
	if !called {
		t.Error("cleanup() did not call the cleanup function")
	}
}
