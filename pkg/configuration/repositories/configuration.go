package repositories

import "fmt"

// Configuration is the structure that holds configuration for a particular configuration.
type Configuration struct {
	// Name is the name of the configuration.
	Name string `json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`
	// GPG is the URL from which to download the GPG key.
	GPG string `json:"gpg" yaml:"gpg" xml:"gpg" toml:"gpg" mapstructure:"gpg"`
	// URL is the URL from which to download the configuration.
	URL string `json:"url" yaml:"url" xml:"url" toml:"url" mapstructure:"url"`
	// Release is the release of the configuration.
	Release string `json:"release" yaml:"release" xml:"release" toml:"release" mapstructure:"release"`
	// Component is the component of the configuration.
	Component string `json:"component" yaml:"component" xml:"component" toml:"component" mapstructure:"component"`
	// KeyName is the name of the key.
	KeyName string `json:"key_name" yaml:"key_name" xml:"key_name" toml:"key_name" mapstructure:"key_name"`
}

// GetName returns the name of the configuration.
func (configuration *Configuration) GetName() string {
	return configuration.Name
}

// GetGPG returns the URL from which to download the GPG key.
func (configuration *Configuration) GetGPG() string {
	return configuration.GPG
}

// GetURL returns the URL from which to download the configuration.
func (configuration *Configuration) GetURL() string {
	return configuration.URL
}

// GetRelease returns the release of the configuration.
func (configuration *Configuration) GetRelease() string {
	return configuration.Release
}

// GetComponent returns the component of the configuration.
func (configuration *Configuration) GetComponent() string {
	return configuration.Component
}

// GetKeyName returns the name of the key.
func (configuration *Configuration) GetKeyName() string {
	return configuration.KeyName
}

// GetKeyFullPath returns the full path to the key.
func (configuration *Configuration) GetKeyFullPath() string {
	return fmt.Sprintf(
		"/usr/share/keyrings/%s.gpg",
		configuration.GetKeyName(),
	)
}

// GetConfigurationFullPath returns the full path to the repository file.
func (configuration *Configuration) GetConfigurationFullPath() string {
	return fmt.Sprintf(
		"/etc/apt/sources.list.d/%s.list",
		configuration.GetName(),
	)
}

// GetConfigurationContents returns the contents of the repository file.
func (configuration *Configuration) GetConfigurationContents() string {
	return fmt.Sprintf(
		"deb [signed-by=%s] %s %s %s",
		configuration.GetKeyFullPath(),
		configuration.GetURL(),
		configuration.GetRelease(),
		configuration.GetComponent(),
	)
}
