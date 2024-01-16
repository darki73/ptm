package qemu

// QemuNetwork is a structure that holds information for QEMU network configuration.
type QemuNetwork struct {
	// Driver is the network driver to use.
	Driver string `json:"driver" yaml:"driver" toml:"driver" mapstructure:"driver"`
	// Bridge is the network bridge to use.
	Bridge string `json:"bridge" yaml:"bridge" toml:"bridge" mapstructure:"bridge"`
}

// InitializeQemuNetworkWithDefaults initializes the QemuNetwork with default values.
func InitializeQemuNetworkWithDefaults() *QemuNetwork {
	return &QemuNetwork{
		Driver: "virtio",
		Bridge: "",
	}
}

// GetDriver returns the network driver to use.
func (qemuNetwork *QemuNetwork) GetDriver() string {
	return qemuNetwork.Driver
}

// GetBridge returns the network bridge to use.
func (qemuNetwork *QemuNetwork) GetBridge() string {
	return qemuNetwork.Bridge
}

// IsConfigured returns true if the configuration is configured.
func (qemuNetwork *QemuNetwork) IsConfigured() bool {
	if qemuNetwork.Driver == "" {
		return false
	}

	if qemuNetwork.Bridge == "" {
		return false
	}

	return true
}
