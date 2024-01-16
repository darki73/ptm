package cloud_init

import (
	"testing"
)

// TestInitializeCloudInitNetworkWithDefaults tests the initialization of the CloudInitNetwork with default values.
func TestInitializeCloudInitNetworkWithDefaults(t *testing.T) {
	networkConfig := InitializeCloudInitNetworkWithDefaults()

	if networkConfig.IPv4 == nil {
		t.Error("Expected IPv4 configuration to be initialized, but got nil")
	}
	if networkConfig.IPv6 == nil {
		t.Error("Expected IPv6 configuration to be initialized, but got nil")
	}
}

// TestCloudInitNetworkGetters tests the getters of the CloudInitNetwork configuration.
func TestCloudInitNetworkGetters(t *testing.T) {
	ipv4Config := InitializeCloudInitNetworkIPv4WithDefaults()
	ipv6Config := InitializeCloudInitNetworkIPv6WithDefaults()

	networkConfig := &CloudInitNetwork{
		IPv4: ipv4Config,
		IPv6: ipv6Config,
	}

	if networkConfig.GetIPv4() != ipv4Config {
		t.Error("GetIPv4() did not return the expected IPv4 configuration")
	}
	if networkConfig.GetIPv6() != ipv6Config {
		t.Error("GetIPv6() did not return the expected IPv6 configuration")
	}
}
