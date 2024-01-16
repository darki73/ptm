package cloud_init

// Configuration is a struct that represents the configuration of cloud-init.
type Configuration struct {
	// Enabled is a flag that indicates whether cloud-init is enabled or not.
	Enabled bool `json:"enabled" yaml:"enabled" toml:"enabled" mapstructure:"enabled"`
	// Username is the username of the user that will be created by cloud-init.
	Username string `json:"username" yaml:"username" toml:"username" mapstructure:"username"`
	// Password is the password of the user that will be created by cloud-init.
	Password string `json:"password" yaml:"password" toml:"password" mapstructure:"password"`
	// Keys is a list of SSH keys that will be added to the user that will be created by cloud-init.
	Keys []string `json:"ssh_authorized_keys" yaml:"ssh_authorized_keys" toml:"ssh_authorized_keys" mapstructure:"ssh_authorized_keys"`
	// Network is a reference to the network configuration that will be created by cloud-init.
	Network *CloudInitNetwork `json:"network" yaml:"network" toml:"network" mapstructure:"network"`
}

// InitializeWithDefaults initializes the configuration with default values.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		Enabled:  false,
		Username: "",
		Password: "",
		Keys:     []string{},
		Network:  InitializeCloudInitNetworkWithDefaults(),
	}
}

// GetEnabled returns the Enabled field value.
func (configuration *Configuration) GetEnabled() bool {
	return configuration.Enabled
}

// GetUsername returns the Username field value.
func (configuration *Configuration) GetUsername() string {
	return configuration.Username
}

// GetPassword returns the Password field value.
func (configuration *Configuration) GetPassword() string {
	return configuration.Password
}

// GetKeys returns the Keys field value.
func (configuration *Configuration) GetKeys() []string {
	return configuration.Keys
}

// GetNetwork returns the Network field value.
func (configuration *Configuration) GetNetwork() *CloudInitNetwork {
	return configuration.Network
}

// IsConfigured returns true if the configuration is configured.
func (configuration *Configuration) IsConfigured() bool {
	if !configuration.Enabled {
		return false
	}

	if configuration.Username == "" {
		return false
	}

	if configuration.Password == "" {
		return false
	}

	if !configuration.Network.IsConfigured() {
		return false
	}

	return true
}
