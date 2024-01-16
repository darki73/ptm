package base_image

// Configuration represents the base image configuration.
type Configuration struct {
	// Distribution is the name of the image distribution.
	Distribution string `json:"distribution" yaml:"distribution" toml:"distribution" mapstructure:"distribution"`
	// Release is the name of the image release.
	Release string `json:"release" yaml:"release" toml:"release" mapstructure:"release"`
	// Minimal is a boolean value that indicates if the image is minimal.
	Minimal bool `json:"minimal" yaml:"minimal" toml:"minimal" mapstructure:"minimal"`
	// Architecture is the name of the image architecture.
	Architecture string `json:"architecture" yaml:"architecture" toml:"architecture" mapstructure:"architecture"`
	// Format is the name of the image format.
	Format string `json:"format" yaml:"format" toml:"format" mapstructure:"format"`
}

// InitializeWithDefaults initializes the base image configuration with default values.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "img",
	}
}

// GetDistribution returns the name of the image distribution.
func (configuration *Configuration) GetDistribution() string {
	return configuration.Distribution
}

// GetRelease returns the name of the image release.
func (configuration *Configuration) GetRelease() string {
	return configuration.Release
}

// GetMinimal returns a boolean value that indicates if the image is minimal.
func (configuration *Configuration) GetMinimal() bool {
	return configuration.Minimal
}

// GetArchitecture returns the name of the image architecture.
func (configuration *Configuration) GetArchitecture() string {
	return configuration.Architecture
}

// GetFormat returns the name of the image format.
func (configuration *Configuration) GetFormat() string {
	return configuration.Format
}
