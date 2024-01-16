package qemu

import (
	"testing"
)

// TestInitializeWithDefaults tests the initialization with default values.
func TestInitializeWithDefaults(t *testing.T) {
	config := InitializeWithDefaults()

	if config.Network == nil {
		t.Errorf("Expected Network to be initialized, got nil")
	}
	if config.Resources == nil {
		t.Errorf("Expected Resources to be initialized, got nil")
	}
	if config.Storage == nil {
		t.Errorf("Expected Storage to be initialized, got nil")
	}
}

// TestConfigurationGetIdentifier tests the GetIdentifier method.
func TestConfigurationGetIdentifier(t *testing.T) {
	config := &Configuration{Identifier: 1234}
	if id := config.GetIdentifier(); id != 1234 {
		t.Errorf("GetIdentifier() = %d, want %d", id, 1234)
	}
}

// TestConfigurationGetName tests the GetName method.
func TestConfigurationGetName(t *testing.T) {
	config := &Configuration{Name: "test-vm"}
	if name := config.GetName(); name != "test-vm" {
		t.Errorf("GetName() = %s, want %s", name, "test-vm")
	}
}

// TestConfigurationGetImage tests the GetImage method.
func TestConfigurationGetImage(t *testing.T) {
	config := &Configuration{Image: "/path/to/image"}
	if image := config.GetImage(); image != "/path/to/image" {
		t.Errorf("GetImage() = %s, want %s", image, "/path/to/image")
	}
}

// TestConfigurationGetNetwork tests the GetNetwork method.
func TestConfigurationGetNetwork(t *testing.T) {
	network := InitializeQemuNetworkWithDefaults()
	config := &Configuration{Network: network}
	if gotNetwork := config.GetNetwork(); gotNetwork != network {
		t.Errorf("GetNetwork() did not return expected network configuration")
	}
}

// TestConfigurationGetResources tests the GetResources method.
func TestConfigurationGetResources(t *testing.T) {
	resources := InitializeQemuResourcesWithDefaults()
	config := &Configuration{Resources: resources}
	if gotResources := config.GetResources(); gotResources != resources {
		t.Errorf("GetResources() did not return expected resources configuration")
	}
}

// TestConfigurationGetStorage tests the GetStorage method.
func TestConfigurationGetStorage(t *testing.T) {
	storage := InitializeQemuStorageWithDefaults()
	config := &Configuration{Storage: storage}
	if gotStorage := config.GetStorage(); gotStorage != storage {
		t.Errorf("GetStorage() did not return expected storage configuration")
	}
}
