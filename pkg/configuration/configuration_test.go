package configuration

import (
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
	ci "github.com/darki73/ptm/pkg/configuration/cloud-init"
	"github.com/darki73/ptm/pkg/configuration/downloader"
	"github.com/darki73/ptm/pkg/configuration/repositories"
	uu "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
	"reflect"
	"testing"
)

// TestConfigurationGetters tests the getter methods of the Configuration struct.
func TestConfigurationGetters(t *testing.T) {
	expectedConfig := &Configuration{
		BaseImage:          &bi.Configuration{},
		CloudInit:          &ci.Configuration{},
		Downloader:         &downloader.Configuration{SaveTo: "/test"},
		Repositories:       []*repositories.Configuration{},
		UnattendedUpgrades: &uu.Configuration{},
		BasePackages:       []string{"package1", "package2"},
		ExtraPackages:      []string{"extra1", "extra2"},
		LogLevel:           "info",
	}

	configuration = expectedConfig

	// Test each getter method
	if !reflect.DeepEqual(configuration.GetBaseImage(), expectedConfig.BaseImage) {
		t.Errorf("GetBaseImage() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetCloudInit(), expectedConfig.CloudInit) {
		t.Errorf("GetCloudInit() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetDownloader(), expectedConfig.Downloader) {
		t.Errorf("GetDownloader() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetRepositories(), expectedConfig.Repositories) {
		t.Errorf("GetRepositories() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetUnattendedUpgrades(), expectedConfig.UnattendedUpgrades) {
		t.Errorf("GetUnattendedUpgrades() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetBasePackages(), expectedConfig.BasePackages) {
		t.Errorf("GetBasePackages() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetExtraPackages(), expectedConfig.ExtraPackages) {
		t.Errorf("GetExtraPackages() did not return expected value")
	}
	if !reflect.DeepEqual(configuration.GetLogLevel(), expectedConfig.LogLevel) {
		t.Errorf("GetLogLevel() did not return expected value")
	}
}
