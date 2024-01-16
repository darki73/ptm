package command

// NewResizeCommand creates a new resize command.
func NewResizeCommand(identifier int, device string, size string) *Command {
	return NewDiskCommand(
		identifier,
		"resize",
		"%VM_ID%",
		device,
		size,
	)
}
