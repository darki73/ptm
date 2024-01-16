package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
)

// NewCloudPasswordCommand creates a new command to set the username for cloud-init.
func NewCloudPasswordCommand(identifier int, cloudInit *ci.CloudInit) *Command {
	return NewSetCommand(
		identifier,
		"--cipassword",
		cloudInit.GetPassword(),
	)
}
