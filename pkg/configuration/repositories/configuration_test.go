package repositories

import (
	"testing"
)

// TestConfiguration tests all methods of the Configuration struct.
func TestConfiguration(t *testing.T) {
	// Sample configuration for testing
	config := Configuration{
		Name:      "test-config",
		GPG:       "https://example.com/gpg",
		URL:       "https://example.com",
		Release:   "test-release",
		Component: "test-component",
		KeyName:   "test-key",
	}

	if name := config.GetName(); name != config.Name {
		t.Errorf("GetName() = %v, want %v", name, config.Name)
	}

	if gpg := config.GetGPG(); gpg != config.GPG {
		t.Errorf("GetGPG() = %v, want %v", gpg, config.GPG)
	}

	if url := config.GetURL(); url != config.URL {
		t.Errorf("GetURL() = %v, want %v", url, config.URL)
	}

	if release := config.GetRelease(); release != config.Release {
		t.Errorf("GetRelease() = %v, want %v", release, config.Release)
	}

	if component := config.GetComponent(); component != config.Component {
		t.Errorf("GetComponent() = %v, want %v", component, config.Component)
	}

	if keyName := config.GetKeyName(); keyName != config.KeyName {
		t.Errorf("GetKeyName() = %v, want %v", keyName, config.KeyName)
	}

	expectedKeyPath := "/usr/share/keyrings/test-key.gpg"
	if keyPath := config.GetKeyFullPath(); keyPath != expectedKeyPath {
		t.Errorf("GetKeyFullPath() = %v, want %v", keyPath, expectedKeyPath)
	}

	expectedConfigPath := "/etc/apt/sources.list.d/test-config.list"
	if configPath := config.GetConfigurationFullPath(); configPath != expectedConfigPath {
		t.Errorf("GetConfigurationFullPath() = %v, want %v", configPath, expectedConfigPath)
	}

	expectedConfigContents := "deb [signed-by=/usr/share/keyrings/test-key.gpg] https://example.com test-release test-component"
	if configContents := config.GetConfigurationContents(); configContents != expectedConfigContents {
		t.Errorf("GetConfigurationContents() = %v, want %v", configContents, expectedConfigContents)
	}
}
