package command

import "fmt"

// NewCloudStorageCommand creates a new cloud storage command (ide2).
func NewCloudStorageCommand(identifier int, storage string) *Command {
	return NewSetCommand(
		identifier,
		"--ide2",
		fmt.Sprintf(
			"%s:cloudinit",
			storage,
		),
	)
}
