package downloader

// Configuration represents the configuration for the downloader.
type Configuration struct {
	// SaveTo is the path to save the downloaded image.
	SaveTo string `json:"save_to" yaml:"save_to" xml:"save_to" toml:"save_to" mapstructure:"save_to"`
}

// InitializeWithDefaults initializes the configuration struct with default values.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		SaveTo: "/etc/ptm/images",
	}
}

// GetSaveTo returns the path to save the downloaded image.
func (configuration *Configuration) GetSaveTo() string {
	return configuration.SaveTo
}
