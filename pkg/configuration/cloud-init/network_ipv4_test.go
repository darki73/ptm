package cloud_init

import (
	"testing"
)

// TestInitializeCloudInitNetworkIPv4WithDefaults tests the initialization of the CloudInitNetworkIPv4 configuration with default values.
func TestInitializeCloudInitNetworkIPv4WithDefaults(t *testing.T) {
	config := InitializeCloudInitNetworkIPv4WithDefaults()

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

// TestCloudInitNetworkIPv4Getters tests the getters of the CloudInitNetworkIPv4 configuration.
func TestCloudInitNetworkIPv4Getters(t *testing.T) {
	config := &CloudInitNetworkIPv4{
		AutoConfigure: false,
		Address:       "10.10.0.10/22",
		Gateway:       "10.10.0.1",
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

// TestCloudInitNetworkIPv4IsConfiguredWithoutAutoConfigure tests the IsConfigured function of the CloudInitNetworkIPv4 configuration without AutoConfigure.
func TestCloudInitNetworkIPv4IsConfiguredWithoutAutoConfigure(t *testing.T) {
	config := &CloudInitNetworkIPv4{
		AutoConfigure: false,
		Address:       "10.10.0.10/22",
		Gateway:       "10.10.0.1",
	}

	if config.IsConfigured() == false {
		t.Errorf("IsConfigured() = %v; want %v", config.IsConfigured(), false)
	}
}

// TestCloudInitNetworkIPv4IsConfiguredWithAutoConfigure tests the IsConfigured function of the CloudInitNetworkIPv4 configuration with AutoConfigure.
func TestCloudInitNetworkIPv4IsConfiguredWithAutoConfigure(t *testing.T) {
	config := &CloudInitNetworkIPv4{
		AutoConfigure: true,
		Address:       "",
		Gateway:       "",
	}

	if config.IsConfigured() != true {
		t.Errorf("IsConfigured() = %v; want %v", config.IsConfigured(), true)
	}
}
