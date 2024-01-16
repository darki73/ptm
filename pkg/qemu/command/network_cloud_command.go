package command

import (
	"fmt"
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
)

// NewNetworkCloudCommand  creates a new cloud network command.
func NewNetworkCloudCommand(identifier int, cloudInit *ci.CloudInit) *Command {
	ipv4 := cloudInit.GetIPv4()
	if ipv4 != "dhcp" {
		ipv4 = fmt.Sprintf(
			"gw4=%s,ip=%s",
			cloudInit.GetIPv4Gateway(),
			cloudInit.GetIPv4(),
		)
	} else {
		ipv4 = fmt.Sprintf(
			"ip=%s",
			cloudInit.GetIPv4(),
		)
	}
	ipv6 := cloudInit.GetIPv6()
	if ipv6 != "auto" {
		ipv6 = fmt.Sprintf(
			"gw6=%s,ip6=%s",
			cloudInit.GetIPv6Gateway(),
			cloudInit.GetIPv6(),
		)
	} else {
		ipv6 = fmt.Sprintf(
			"ip6=%s",
			cloudInit.GetIPv6(),
		)
	}

	return NewSetCommand(
		identifier,
		"--ipconfig0",
		fmt.Sprintf(
			"%s,%s",
			ipv6,
			ipv4,
		),
	)
}
