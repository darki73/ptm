package command

// NewNameCommand creates a new name command.
func NewNameCommand(identifier int, name string) *Command {
	return NewCreateCommand(
		identifier,
		"--name",
		name,
	)
}
