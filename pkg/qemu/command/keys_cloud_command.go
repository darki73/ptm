package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
)

// NewCloudKeysCommand creates a new cloud ssh keys command.
func NewCloudKeysCommand(identifier int, cloudInit *ci.CloudInit) *Command {
	return NewSetCommand(
		identifier,
		"--sshkey",
		cloudInit.GetSSHKeysTemporaryFilePath(),
	)
}
