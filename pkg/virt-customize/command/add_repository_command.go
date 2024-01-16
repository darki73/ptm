package command

import (
	"fmt"
	"github.com/darki73/ptm/pkg/configuration/repositories"
)

// NewAddRepositoryCommand creates a new add repository command.
func NewAddRepositoryCommand(image string, repository *repositories.Configuration) *Command {
	return NewCommand(
		image,
		"--run-command",
		fmt.Sprintf(
			"echo \"%s\" > %s",
			repository.GetConfigurationContents(),
			repository.GetConfigurationFullPath(),
		),
	)
}
