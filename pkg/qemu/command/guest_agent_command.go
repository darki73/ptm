package command

import (
	"fmt"
	"github.com/darki73/ptm/pkg/utils"
)

// NewGuestAgentCommand returns a new GuestAgentCommand.
func NewGuestAgentCommand(identifier int, enabled bool, trim bool) *Command {
	enabledInt := utils.BooleanToInteger(enabled)
	trimInt := utils.BooleanToInteger(trim)

	guestAgent := fmt.Sprintf("enabled=%d", enabledInt)
	if trim {
		guestAgent = fmt.Sprintf("%s,fstrim_cloned_disks=%d", guestAgent, trimInt)
	}

	return NewSetCommand(
		identifier,
		"--agent",
		guestAgent,
	)
}
