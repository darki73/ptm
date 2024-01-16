package qemu

import (
	"testing"
)

// TestInitializeQemuNetworkWithDefaults tests the initialization with default values.
func TestInitializeQemuNetworkWithDefaults(t *testing.T) {
	qemuNetwork := InitializeQemuNetworkWithDefaults()

	if qemuNetwork.Driver != "virtio" {
		t.Errorf("Expected Driver to be 'virtio', got %s", qemuNetwork.Driver)
	}
}

// TestGetDriver tests the GetDriver method.
func TestGetDriver(t *testing.T) {
	qemuNetwork := &QemuNetwork{Driver: "e1000"}
	if driver := qemuNetwork.GetDriver(); driver != "e1000" {
		t.Errorf("GetDriver() = %s, want %s", driver, "e1000")
	}
}

// TestGetBridge tests the GetBridge method.
func TestGetBridge(t *testing.T) {
	qemuNetwork := &QemuNetwork{Bridge: "test-bridge"}
	if bridge := qemuNetwork.GetBridge(); bridge != "test-bridge" {
		t.Errorf("GetBridge() = %s, want %s", bridge, "test-bridge")
	}
}

// TestNetworkIsConfigured tests the IsConfigured method.
func TestNetworkIsConfigured(t *testing.T) {
	qemuNetwork := &QemuNetwork{Driver: "e1000", Bridge: "test-bridge"}
	if configured := qemuNetwork.IsConfigured(); !configured {
		t.Errorf("IsConfigured() = %t, want %t", configured, true)
	}
}
