package unattended_upgrades

import (
	"github.com/darki73/ptm/pkg/log"
	"github.com/google/go-cmp/cmp"
)

// Configuration is the structure that holds configuration for unattended upgrades.
type Configuration struct {
	// Enabled is the flag that enables unattended upgrades.
	Enabled bool `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	// Whitelist is the list of packages that should be upgraded.
	Whitelist []string `json:"whitelist" yaml:"whitelist" xml:"whitelist" toml:"whitelist" mapstructure:"whitelist"`
	// Blacklist is the list of packages that should not be upgraded.
	Blacklist []string `json:"blacklist" yaml:"blacklist" xml:"blacklist" toml:"blacklist" mapstructure:"blacklist"`
	// DevRelease is the flag that enables upgrades to development releases.
	DevRelease string `json:"dev_release" yaml:"dev_release" xml:"dev_release" toml:"dev_release" mapstructure:"dev_release"`
	// FixInterrupted is the flag that enables fixing of interrupted upgrades.
	FixInterrupted bool `json:"fix_interrupted" yaml:"fix_interrupted" xml:"fix_interrupted" toml:"fix_interrupted" mapstructure:"fix_interrupted"`
	// MinimalSteps is the flag that enables minimal steps for upgrades.
	MinimalSteps bool `json:"minimal_steps" yaml:"minimal_steps" xml:"minimal_steps" toml:"minimal_steps" mapstructure:"minimal_steps"`
	// InstallOnShutdown is the flag that enables installation of upgrades on shutdown.
	InstallOnShutdown bool `json:"install_on_shutdown" yaml:"install_on_shutdown" xml:"install_on_shutdown" toml:"install_on_shutdown" mapstructure:"install_on_shutdown"`
	// RemoveUnusedKernel is the flag that enables removal of unused kernel packages.
	RemoveUnusedKernel bool `json:"remove_unused_kernel" yaml:"remove_unused_kernel" xml:"remove_unused_kernel" toml:"remove_unused_kernel" mapstructure:"remove_unused_kernel"`
	// RemoveUnusedDependencies is the flag that enables removal of unused dependencies.
	RemoveUnusedDependencies bool `json:"remove_unused_dependencies" yaml:"remove_unused_dependencies" xml:"remove_unused_dependencies" toml:"remove_unused_dependencies" mapstructure:"remove_unused_dependencies"`
	// RemoveUnusedAutoDepend is the flag that enables removal of unused automatically installed dependencies.
	RemoveUnusedAutoDepend bool `json:"remove_unused_auto_depend" yaml:"remove_unused_auto_depend" xml:"remove_unused_auto_depend" toml:"remove_unused_auto_depend" mapstructure:"remove_unused_auto_depend"`
	// AutomaticReboot is the flag that enables automatic reboot after upgrades.
	AutomaticReboot bool `json:"automatic_reboot" yaml:"automatic_reboot" xml:"automatic_reboot" toml:"automatic_reboot" mapstructure:"automatic_reboot"`
	// AutomaticRebootWithUsers is the flag that enables automatic reboot even if there are users logged in.
	AutomaticRebootWithUsers bool `json:"automatic_reboot_with_users" yaml:"automatic_reboot_with_users" xml:"automatic_reboot_with_users" toml:"automatic_reboot_with_users" mapstructure:"automatic_reboot_with_users"`
	// AutomaticRebootTime is the time of automatic reboot after upgrades.
	AutomaticRebootTime string `json:"automatic_reboot_time" yaml:"automatic_reboot_time" xml:"automatic_reboot_time" toml:"automatic_reboot_time" mapstructure:"automatic_reboot_time"`
}

// InitializeWithDefaults initializes the configuration struct with default values.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		Enabled: false,
		Whitelist: []string{
			"${distro_id}:${distro_codename}",
			"${distro_id}:${distro_codename}-security",
			"${distro_id}ESMApps:${distro_codename}-apps-security",
			"${distro_id}ESM:${distro_codename}-infra-security",
			"${distro_id}:${distro_codename}-updates",
		},
		Blacklist:                []string{},
		DevRelease:               "auto",
		FixInterrupted:           true,
		MinimalSteps:             true,
		InstallOnShutdown:        false,
		RemoveUnusedKernel:       true,
		RemoveUnusedDependencies: true,
		RemoveUnusedAutoDepend:   true,
		AutomaticReboot:          false,
		AutomaticRebootWithUsers: false,
		AutomaticRebootTime:      "04:00",
	}
}

// GetEnabled returns the flag that enables unattended upgrades.
func (configuration *Configuration) GetEnabled() bool {
	return configuration.Enabled
}

// GetWhitelist returns the list of packages that should be upgraded.
func (configuration *Configuration) GetWhitelist() []string {
	return configuration.Whitelist
}

// SetWhitelist sets the list of packages that should be upgraded.
func (configuration *Configuration) SetWhitelist(whitelist []string) *Configuration {
	if !configuration.IsAutoConfigured() {
		log.Fatal("cannot set whitelist when configuration is not autoconfigured")
	}
	configuration.Whitelist = whitelist
	return configuration
}

// GetBlacklist returns the list of packages that should not be upgraded.
func (configuration *Configuration) GetBlacklist() []string {
	return configuration.Blacklist
}

// SetBlacklist sets the list of packages that should not be upgraded.
func (configuration *Configuration) SetBlacklist(blacklist []string) *Configuration {
	if !configuration.IsAutoConfigured() {
		log.Fatal("cannot set blacklist when configuration is not autoconfigured")
	}
	configuration.Blacklist = blacklist
	return configuration
}

// GetDevRelease returns the flag that enables upgrades to development releases.
func (configuration *Configuration) GetDevRelease() string {
	return configuration.DevRelease
}

// GetFixInterrupted returns the flag that enables fixing of interrupted upgrades.
func (configuration *Configuration) GetFixInterrupted() bool {
	return configuration.FixInterrupted
}

// GetMinimalSteps returns the flag that enables minimal steps for upgrades.
func (configuration *Configuration) GetMinimalSteps() bool {
	return configuration.MinimalSteps
}

// GetInstallOnShutdown returns the flag that enables installation of upgrades on shutdown.
func (configuration *Configuration) GetInstallOnShutdown() bool {
	return configuration.InstallOnShutdown
}

// GetRemoveUnusedKernel returns the flag that enables removal of unused kernel packages.
func (configuration *Configuration) GetRemoveUnusedKernel() bool {
	return configuration.RemoveUnusedKernel
}

// GetRemoveUnusedDependencies returns the flag that enables removal of unused dependencies.
func (configuration *Configuration) GetRemoveUnusedDependencies() bool {
	return configuration.RemoveUnusedDependencies
}

// GetRemoveUnusedAutoDepend returns the flag that enables removal of unused automatically installed dependencies.
func (configuration *Configuration) GetRemoveUnusedAutoDepend() bool {
	return configuration.RemoveUnusedAutoDepend
}

// GetAutomaticReboot returns the flag that enables automatic reboot after upgrades.
func (configuration *Configuration) GetAutomaticReboot() bool {
	return configuration.AutomaticReboot
}

// GetAutomaticRebootWithUsers returns the flag that enables automatic reboot even if there are users logged in.
func (configuration *Configuration) GetAutomaticRebootWithUsers() bool {
	return configuration.AutomaticRebootWithUsers
}

// GetAutomaticRebootTime returns the time of automatic reboot after upgrades.
func (configuration *Configuration) GetAutomaticRebootTime() string {
	return configuration.AutomaticRebootTime
}

// IsAutoConfigured returns true if the configuration is autoconfigured.
func (configuration *Configuration) IsAutoConfigured() bool {
	defaultConfigurationWithDisabled := InitializeWithDefaults()
	defaultConfigurationWithEnabled := InitializeWithDefaults()
	defaultConfigurationWithEnabled.Enabled = true

	if cmp.Equal(configuration, defaultConfigurationWithDisabled) {
		return true
	}

	if cmp.Equal(configuration, defaultConfigurationWithEnabled) {
		return true
	}

	return false
}
