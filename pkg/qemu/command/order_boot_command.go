package command

import "fmt"

// NewBootOrderCommand creates a new boot order command.
func NewBootOrderCommand(identifier int, device string, driver string) *Command {
	return NewSetCommand(
		identifier,
		"--boot",
		fmt.Sprintf(
			"order=%s",
			device,
		),
		"--scsihw",
		driver,
	)
}
