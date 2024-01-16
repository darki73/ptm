package command

import "fmt"

// NewUploadCommand creates a new upload command.
func NewUploadCommand(image string, source string, target string) *Command {
	return NewCommand(
		image,
		"--upload",
		fmt.Sprintf("%s:%s", source, target),
	)
}
