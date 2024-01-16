package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
)

// NewCloudUsernameCommand creates a new command to set the username for cloud-init.
func NewCloudUsernameCommand(identifier int, cloudInit *ci.CloudInit) *Command {
	return NewSetCommand(
		identifier,
		"--ciuser",
		cloudInit.GetUsername(),
	)
}
