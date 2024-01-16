package cloud_init

import (
	"testing"
)

// TestInitializeCloudInitNetworkIPv6WithDefaults tests the initialization of the CloudInitNetworkIPv6 configuration with default values.
func TestInitializeCloudInitNetworkIPv6WithDefaults(t *testing.T) {
	config := InitializeCloudInitNetworkIPv6WithDefaults()

	if config.AutoConfigure != true {
		t.Errorf("Expected AutoConfigure to be true, got %v", config.AutoConfigure)
	}
	if config.Address != "" {
		t.Errorf("Expected Address to be empty, got %s", config.Address)
	}
	if config.Gateway != "" {
		t.Errorf("Expected Gateway to be empty, got %s", config.Gateway)
	}
}

// TestCloudInitNetworkIPv6Getters tests the getters of the CloudInitNetworkIPv6 configuration.
func TestCloudInitNetworkIPv6Getters(t *testing.T) {
	config := &CloudInitNetworkIPv6{
		AutoConfigure: false,
		Address:       "2001:db8::1",
		Gateway:       "2001:db8::ff",
	}

	if config.GetAutoConfigure() != config.AutoConfigure {
		t.Errorf("GetAutoConfigure() = %v; want %v", config.GetAutoConfigure(), config.AutoConfigure)
	}
	if config.GetAddress() != config.Address {
		t.Errorf("GetAddress() = %s; want %s", config.GetAddress(), config.Address)
	}
	if config.GetGateway() != config.Gateway {
		t.Errorf("GetGateway() = %s; want %s", config.GetGateway(), config.Gateway)
	}
}

// TestCloudInitNetworkIPv6IsConfiguredWithoutAutoConfigure tests the IsConfigured function of the CloudInitNetworkIPv6 configuration without AutoConfigure.
func TestCloudInitNetworkIPv6IsConfiguredWithoutAutoConfigure(t *testing.T) {
	config := &CloudInitNetworkIPv6{
		AutoConfigure: false,
		Address:       "2001:db8::1",
		Gateway:       "2001:db8::ff",
	}

	if config.IsConfigured() != true {
		t.Errorf("IsConfigured() = %v; want %v", config.IsConfigured(), true)
	}
}

// TestCloudInitNetworkIPv6IsConfiguredWithAutoConfigure tests the IsConfigured function of the CloudInitNetworkIPv6 configuration with AutoConfigure.
func TestCloudInitNetworkIPv6IsConfiguredWithAutoConfigure(t *testing.T) {
	config := &CloudInitNetworkIPv6{
		AutoConfigure: true,
		Address:       "",
		Gateway:       "",
	}

	if config.IsConfigured() == false {
		t.Errorf("IsConfigured() = %v; want %v", config.IsConfigured(), false)
	}
}
