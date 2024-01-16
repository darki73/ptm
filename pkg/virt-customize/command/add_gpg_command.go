package command

import (
	"fmt"
	"github.com/darki73/ptm/pkg/configuration/repositories"
)

// NewAddGPGCommand creates a new add gpg command.
func NewAddGPGCommand(image string, repository *repositories.Configuration) *Command {
	return NewCommand(
		image,
		"--run-command",
		fmt.Sprintf(
			"curl -fsSL %s | gpg --dearmor -o %s",
			repository.GetGPG(),
			repository.GetKeyFullPath(),
		),
	)
}
