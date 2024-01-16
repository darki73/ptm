package qemu

// Configuration is the configuration for the Qemu.
type Configuration struct {
	// Identifier is the identifier of the virtual machine.
	Identifier int `json:"identifier" yaml:"identifier" toml:"identifier" mapstructure:"identifier"`
	// Name is the name of the virtual machine.
	Name string `json:"name" yaml:"name" toml:"name" mapstructure:"name"`
	// Image is the path to the image used to create the virtual machine.
	Image string `json:"image" yaml:"image" toml:"image" mapstructure:"image"`
	// Network is the reference to the network configuration.
	Network *QemuNetwork `json:"network" yaml:"network" toml:"network" mapstructure:"network"`
	// Resources is the reference to the resources configuration.
	Resources *QemuResources `json:"resources" yaml:"resources" toml:"resources" mapstructure:"resources"`
	// Storage is the reference to the storage configuration.
	Storage *QemuStorage `json:"storage" yaml:"storage" toml:"storage" mapstructure:"storage"`
}

// InitializeWithDefaults initializes the configuration with default values.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		Identifier: 0,
		Name:       "",
		Image:      "",
		Network:    InitializeQemuNetworkWithDefaults(),
		Resources:  InitializeQemuResourcesWithDefaults(),
		Storage:    InitializeQemuStorageWithDefaults(),
	}
}

// GetIdentifier returns the identifier of the virtual machine.
func (configuration *Configuration) GetIdentifier() int {
	return configuration.Identifier
}

// GetName returns the name of the virtual machine.
func (configuration *Configuration) GetName() string {
	return configuration.Name
}

// GetImage returns the path to the image used to create the virtual machine.
func (configuration *Configuration) GetImage() string {
	return configuration.Image
}

// GetNetwork returns the reference to the network configuration.
func (configuration *Configuration) GetNetwork() *QemuNetwork {
	return configuration.Network
}

// GetResources returns the reference to the resources configuration.
func (configuration *Configuration) GetResources() *QemuResources {
	return configuration.Resources
}

// GetStorage returns the reference to the storage configuration.
func (configuration *Configuration) GetStorage() *QemuStorage {
	return configuration.Storage
}

// IsConfigured returns true if the configuration is configured.
func (configuration *Configuration) IsConfigured() bool {
	if configuration.Identifier == 0 {
		return false
	}

	if configuration.Name == "" {
		return false
	}

	if !configuration.Network.IsConfigured() {
		return false
	}

	if !configuration.Resources.IsConfigured() {
		return false
	}

	if !configuration.Storage.IsConfigured() {
		return false
	}

	return true
}
