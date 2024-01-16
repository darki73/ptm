package command

import "strings"

// NewInstallCommand creates a new install command.
func NewInstallCommand(image string, packages []string) *Command {
	return NewCommand(
		image,
		"--install",
		strings.Join(packages, ","),
	)
}
