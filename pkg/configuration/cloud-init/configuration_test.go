package cloud_init

import (
	"reflect"
	"testing"
)

// TestInitializeWithDefaults tests the initialization of the Configuration with default values.
func TestInitializeWithDefaults(t *testing.T) {
	config := InitializeWithDefaults()

	if config.Enabled {
		t.Errorf("Expected Enabled to be false, got %v", config.Enabled)
	}
	if config.Username != "" {
		t.Errorf("Expected Username to be empty, got %s", config.Username)
	}
	if config.Password != "" {
		t.Errorf("Expected Password to be empty, got %s", config.Password)
	}
	if len(config.Keys) != 0 {
		t.Errorf("Expected Keys to be empty, got %v", config.Keys)
	}
	if config.Network == nil {
		t.Error("Expected Network configuration to be initialized, but got nil")
	}
}

// TestConfigurationGetters tests the getters of the Configuration.
func TestConfigurationGetters(t *testing.T) {
	keys := []string{"key1", "key2"}
	networkConfig := InitializeCloudInitNetworkWithDefaults()

	config := &Configuration{
		Enabled:  true,
		Username: "testuser",
		Password: "testpass",
		Keys:     keys,
		Network:  networkConfig,
	}

	if config.GetEnabled() != config.Enabled {
		t.Errorf("GetEnabled() = %v; want %v", config.GetEnabled(), config.Enabled)
	}
	if config.GetUsername() != config.Username {
		t.Errorf("GetUsername() = %s; want %s", config.GetUsername(), config.Username)
	}
	if config.GetPassword() != config.Password {
		t.Errorf("GetPassword() = %s; want %s", config.GetPassword(), config.Password)
	}
	if !reflect.DeepEqual(config.GetKeys(), config.Keys) {
		t.Errorf("GetKeys() = %v; want %v", config.GetKeys(), config.Keys)
	}
	if config.GetNetwork() != config.Network {
		t.Error("GetNetwork() did not return the expected Network configuration")
	}
}
