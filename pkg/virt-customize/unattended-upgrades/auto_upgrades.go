package unattended_upgrades

// autoUpgradesTemplate represents the template for the auto-upgrades configuration.
// /etc/apt/apt.conf.d/20auto-upgrades
const autoUpgradesTemplate = `
APT::Periodic::Update-Package-Lists "1";
APT::Periodic::Download-Upgradeable-Packages "1";
APT::Periodic::AutocleanInterval "7";
APT::Periodic::Unattended-Upgrade "1";
`

// GetAutoUpgradesTemplate returns the auto-upgrades template.
func GetAutoUpgradesTemplate() string {
	return autoUpgradesTemplate
}

// GetAutoUpgradesConfigurationPath returns the auto-upgrades configuration path.
func GetAutoUpgradesConfigurationPath() string {
	return "/etc/apt/apt.conf.d/20auto-upgrades"
}

// GetAutoUpgradesTemporaryPath returns the auto-upgrades temporary path.
func GetAutoUpgradesTemporaryPath() string {
	return "/tmp/ptm-20auto-upgrades"
}
