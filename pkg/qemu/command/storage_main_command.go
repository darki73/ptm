package command

import "fmt"

// NewMainStorageCommand creates a new main storage command (scsi0).
func NewMainStorageCommand(identifier int, storage string, image string) *Command {
	scsiCommand := fmt.Sprintf("%s:0,import-from=%s,discard=on", storage, image)
	return NewSetCommand(
		identifier,
		"--scsi0",
		scsiCommand,
	)
}
