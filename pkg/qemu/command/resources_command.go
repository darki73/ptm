package command

// NewResourcesCommand creates a new memory and cpu cores command.
func NewResourcesCommand(identifier int, cores int, memory int, cpuType string) *Command {
	return NewSetCommand(
		identifier,
		"--cores",
		cores,
		"--memory",
		memory,
		"--cpu",
		cpuType,
	)
}
