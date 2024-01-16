package cloud_init

// CloudInitNetwork represents the network configuration for cloud-init.
type CloudInitNetwork struct {
	// IPv4 is a reference to the IPv4 configuration.
	IPv4 *CloudInitNetworkIPv4 `json:"ipv4" yaml:"ipv4" toml:"ipv4" mapstructure:"ipv4"`
	// IPv6 is a reference to the IPv6 configuration.
	IPv6 *CloudInitNetworkIPv6 `json:"ipv6" yaml:"ipv6" toml:"ipv6" mapstructure:"ipv6"`
}

// InitializeCloudInitNetworkWithDefaults initializes CloudInitNetwork with default values.
func InitializeCloudInitNetworkWithDefaults() *CloudInitNetwork {
	return &CloudInitNetwork{
		IPv4: InitializeCloudInitNetworkIPv4WithDefaults(),
		IPv6: InitializeCloudInitNetworkIPv6WithDefaults(),
	}
}

// GetIPv4 returns the IPv4 configuration.
func (cin *CloudInitNetwork) GetIPv4() *CloudInitNetworkIPv4 {
	return cin.IPv4
}

// GetIPv6 returns the IPv6 configuration.
func (cin *CloudInitNetwork) GetIPv6() *CloudInitNetworkIPv6 {
	return cin.IPv6
}

// IsConfigured returns true if the configuration is configured.
func (cin *CloudInitNetwork) IsConfigured() bool {
	return cin.IPv4.IsConfigured() || cin.IPv6.IsConfigured()
}
