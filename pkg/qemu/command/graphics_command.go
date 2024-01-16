package command

// NewGraphicsCommand creates a new graphics command.
func NewGraphicsCommand(identifier int) *Command {
	return NewSetCommand(
		identifier,
		"--serial0",
		"socket",
		"--vga",
		"serial0",
	)
}
