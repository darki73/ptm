package virt_customize

import (
	"github.com/darki73/ptm/pkg/configuration/repositories"
	uu "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
	"reflect"
	"testing"
)

var (
	// imagePathForTesting is the path to the image used for testing.
	imagePathForTesting    = "/etc/ptm/images/ubuntu-22.04-cloudimage-amd64.img"
	basePackagesForTesting = []string{"pkg1", "pkg2"}
	// extraPackagesForTesting is a list of packages to be installed.
	extraPackagesForTesting = []string{"extra1", "extra2"}
	// repositoriesConfigurationForTesting is a configuration of repositories for testing.
	repositoriesConfigurationForTesting = []*repositories.Configuration{
		{
			Name:      "docker",
			GPG:       "https://download.docker.com/linux/ubuntu/gpg",
			URL:       "https://download.docker.com/linux/ubuntu",
			Release:   "jammy",
			Component: "stable",
			KeyName:   "docker-archive-keyring",
		},
	}
	// unattendedUpgradesConfigurationForTesting is a configuration of unattended upgrades for testing.
	unattendedUpgradesConfigurationForTesting = &uu.Configuration{
		Enabled: true,
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
)

// TestVirtCustomizeGetImage tests the GetImage() function.
func TestVirtCustomizeGetImage(t *testing.T) {
	vc := NewVirtCustomizeConfiguration(
		imagePathForTesting,
		basePackagesForTesting,
		extraPackagesForTesting,
		repositoriesConfigurationForTesting,
		unattendedUpgradesConfigurationForTesting,
	)

	result := vc.GetImage()

	if result != imagePathForTesting {
		t.Errorf("Expected GetImage() to return '%s', but got '%s'", imagePathForTesting, result)
	}
}

// TestVirtCustomizeGetBasePackages tests the GetBasePackages() method.
func TestVirtCustomizeGetBasePackages(t *testing.T) {
	vc := NewVirtCustomizeConfiguration(
		imagePathForTesting,
		basePackagesForTesting,
		extraPackagesForTesting,
		repositoriesConfigurationForTesting,
		unattendedUpgradesConfigurationForTesting,
	)

	result := vc.GetBasePackages()

	if !reflect.DeepEqual(result, basePackagesForTesting) {
		t.Errorf("Expected GetBasePackages() to return '%v', but got '%v'", basePackagesForTesting, result)
	}
}

// TestVirtCustomizeGetExtraPackages tests the GetExtraPackages function.
func TestVirtCustomizeGetExtraPackages(t *testing.T) {
	vc := NewVirtCustomizeConfiguration(
		imagePathForTesting,
		basePackagesForTesting,
		extraPackagesForTesting,
		repositoriesConfigurationForTesting,
		unattendedUpgradesConfigurationForTesting,
	)

	result := vc.GetExtraPackages()

	if !reflect.DeepEqual(result, extraPackagesForTesting) {
		t.Errorf("Expected GetExtraPackages() to return '%v', but got '%v'", extraPackagesForTesting, result)
	}
}

// TestVirtCustomizeGetPackages tests the GetPackages function.
func TestVirtCustomizeGetPackages(t *testing.T) {
	vc := NewVirtCustomizeConfiguration(
		imagePathForTesting,
		basePackagesForTesting,
		extraPackagesForTesting,
		repositoriesConfigurationForTesting,
		unattendedUpgradesConfigurationForTesting,
	)

	result := vc.GetPackages()

	expectedPackages := append(basePackagesForTesting, extraPackagesForTesting...)

	if !reflect.DeepEqual(result, expectedPackages) {
		t.Errorf("Expected GetPackages() to return '%v', but got '%v'", expectedPackages, result)
	}
}

// TestVirtCustomizeGetRepositories tests the GetRepositories function.
func TestVirtCustomizeGetRepositories(t *testing.T) {
	vc := NewVirtCustomizeConfiguration(
		imagePathForTesting,
		basePackagesForTesting,
		extraPackagesForTesting,
		repositoriesConfigurationForTesting,
		unattendedUpgradesConfigurationForTesting,
	)

	result := vc.GetRepositoriesConfiguration()

	if !reflect.DeepEqual(result, repositoriesConfigurationForTesting) {
		t.Errorf("Expected GetRepositories() to return '%v', but got '%v'", repositoriesConfigurationForTesting, result)
	}
}

// TestVirtCustomizeGetUnattendedUpgrades tests the GetUnattendedUpgrades function.
func TestVirtCustomizeGetUnattendedUpgrades(t *testing.T) {
	vc := NewVirtCustomizeConfiguration(
		imagePathForTesting,
		basePackagesForTesting,
		extraPackagesForTesting,
		repositoriesConfigurationForTesting,
		unattendedUpgradesConfigurationForTesting,
	)

	result := vc.GetUnattendedUpgradesConfiguration()

	if !reflect.DeepEqual(result, unattendedUpgradesConfigurationForTesting) {
		t.Errorf("Expected GetUnattendedUpgrades() to return '%v', but got '%v'", unattendedUpgradesConfigurationForTesting, result)
	}
}
