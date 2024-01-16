package command

// NewUpdateCommand creates a new update command.
func NewUpdateCommand(image string) *Command {
	return NewCommand(
		image,
		"--update",
	)
}
