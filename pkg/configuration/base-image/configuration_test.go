package base_image

import (
	"testing"
)

// TestInitializeWithDefaults tests the initialization of the base image configuration with default values.
func TestInitializeWithDefaults(t *testing.T) {
	config := InitializeWithDefaults()

	expectedDistribution := "ubuntu"
	expectedRelease := "jammy"
	expectedMinimal := true
	expectedArchitecture := "amd64"
	expectedFormat := "img"

	if config.Distribution != expectedDistribution {
		t.Errorf("Expected Distribution %s, got %s", expectedDistribution, config.Distribution)
	}
	if config.Release != expectedRelease {
		t.Errorf("Expected Release %s, got %s", expectedRelease, config.Release)
	}
	if config.Minimal != expectedMinimal {
		t.Errorf("Expected Minimal %v, got %v", expectedMinimal, config.Minimal)
	}
	if config.Architecture != expectedArchitecture {
		t.Errorf("Expected Architecture %s, got %s", expectedArchitecture, config.Architecture)
	}
	if config.Format != expectedFormat {
		t.Errorf("Expected Format %s, got %s", expectedFormat, config.Format)
	}
}

// TestConfigurationGetters tests the getters of the base image configuration.
func TestConfigurationGetters(t *testing.T) {
	config := &Configuration{
		Distribution: "debian",
		Release:      "buster",
		Minimal:      false,
		Architecture: "arm",
		Format:       "qcow2",
	}

	if config.GetDistribution() != config.Distribution {
		t.Errorf("GetDistribution() = %s; want %s", config.GetDistribution(), config.Distribution)
	}
	if config.GetRelease() != config.Release {
		t.Errorf("GetRelease() = %s; want %s", config.GetRelease(), config.Release)
	}
	if config.GetMinimal() != config.Minimal {
		t.Errorf("GetMinimal() = %v; want %v", config.GetMinimal(), config.Minimal)
	}
	if config.GetArchitecture() != config.Architecture {
		t.Errorf("GetArchitecture() = %s; want %s", config.GetArchitecture(), config.Architecture)
	}
	if config.GetFormat() != config.Format {
		t.Errorf("GetFormat() = %s; want %s", config.GetFormat(), config.Format)
	}
}
