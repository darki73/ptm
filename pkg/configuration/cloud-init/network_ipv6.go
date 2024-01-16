package cloud_init

// CloudInitNetworkIPv6 is a struct that represents the IPv6 network configuration of a cloud-init configuration.
type CloudInitNetworkIPv6 struct {
	// AutoConfigure is a boolean value that indicates whether the network configuration should be automatically configured.
	AutoConfigure bool `json:"auto_configure" yaml:"auto_configure" toml:"auto_configure" mapstructure:"auto_configure"`
	// Address is the IPv6 address of the network interface.
	Address string `json:"ip" yaml:"ip" toml:"ip" mapstructure:"ip"`
	// Gateway is the IPv6 gateway of the network interface.
	Gateway string `json:"gateway" yaml:"gateway" toml:"gateway" mapstructure:"gateway"`
}

// InitializeCloudInitNetworkIPv6WithDefaults initializes a CloudInitNetworkIPv6 struct with default values.
func InitializeCloudInitNetworkIPv6WithDefaults() *CloudInitNetworkIPv6 {
	return &CloudInitNetworkIPv6{
		AutoConfigure: true,
		Address:       "",
		Gateway:       "",
	}
}

// GetAutoConfigure returns the AutoConfigure field value.
func (cinipv6 *CloudInitNetworkIPv6) GetAutoConfigure() bool {
	return cinipv6.AutoConfigure
}

// GetAddress returns the Address field value.
func (cinipv6 *CloudInitNetworkIPv6) GetAddress() string {
	return cinipv6.Address
}

// GetGateway returns the Gateway field value.
func (cinipv6 *CloudInitNetworkIPv6) GetGateway() string {
	return cinipv6.Gateway
}

// IsConfigured returns true if the configuration is configured.
func (cinipv6 *CloudInitNetworkIPv6) IsConfigured() bool {
	if cinipv6.AutoConfigure {
		return true
	}
	return cinipv6.Address != "" && cinipv6.Gateway != ""
}
