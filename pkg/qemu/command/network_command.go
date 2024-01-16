package command

import "fmt"

// NewNetworkCommand creates a new network command.
func NewNetworkCommand(identifier int, driver string, bridge string) *Command {
	return NewSetCommand(
		identifier,
		"--net0",
		fmt.Sprintf(
			"%s,bridge=%s",
			driver,
			bridge,
		),
	)
}
