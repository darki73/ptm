package virt_customize

import (
	"github.com/darki73/ptm/pkg/configuration/repositories"
	uu "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
)

// VirtCustomize is a structure that holds information for virt-customize configurator.
type VirtCustomize struct {
	// image is the path to the image to be customized.
	image string
	// basePackages is a list of packages to be installed.
	basePackages []string
	// extraPackages is a list of packages to be installed.
	extraPackages []string
	// repositoriesConfiguration is a list of repositories to be added.
	repositoriesConfiguration []*repositories.Configuration
	// unattendedUpgradesConfiguration is a reference to unattended upgrades configuration.
	unattendedUpgradesConfiguration *uu.Configuration
}

// NewVirtCustomizeConfiguration creates a new virt-customize configuration.
func NewVirtCustomizeConfiguration(
	image string,
	basePackages []string,
	extraPackages []string,
	repositoriesConfiguration []*repositories.Configuration,
	unattendedUpgradesConfiguration *uu.Configuration,
) *VirtCustomize {
	return &VirtCustomize{
		image:                           image,
		basePackages:                    basePackages,
		extraPackages:                   extraPackages,
		repositoriesConfiguration:       repositoriesConfiguration,
		unattendedUpgradesConfiguration: unattendedUpgradesConfiguration,
	}
}

// GetImage returns the path to the image to be customized.
func (vc *VirtCustomize) GetImage() string {
	return vc.image
}

// GetBasePackages returns a list of packages to be installed.
func (vc *VirtCustomize) GetBasePackages() []string {
	return vc.basePackages
}

// GetExtraPackages returns a list of packages to be installed.
func (vc *VirtCustomize) GetExtraPackages() []string {
	return vc.extraPackages
}

// GetPackages returns a list of packages to be installed.
func (vc *VirtCustomize) GetPackages() []string {
	return append(vc.basePackages, vc.extraPackages...)
}

// GetRepositoriesConfiguration returns a list of repositories to be added.
func (vc *VirtCustomize) GetRepositoriesConfiguration() []*repositories.Configuration {
	return vc.repositoriesConfiguration
}

// GetUnattendedUpgradesConfiguration returns a reference to unattended upgrades configuration.
func (vc *VirtCustomize) GetUnattendedUpgradesConfiguration() *uu.Configuration {
	return vc.unattendedUpgradesConfiguration
}
