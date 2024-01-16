package cloud_init

// CloudInitNetworkIPv4 is a struct that represents the IPv4 network configuration of a cloud-init configuration.
type CloudInitNetworkIPv4 struct {
	// AutoConfigure is a boolean value that indicates whether the network configuration should be automatically configured.
	AutoConfigure bool `json:"auto_configure" yaml:"auto_configure" toml:"auto_configure" mapstructure:"auto_configure"`
	// Address is the IPv4 address of the network interface.
	Address string `json:"ip" yaml:"ip" toml:"ip" mapstructure:"ip"`
	// Gateway is the IPv4 gateway of the network interface.
	Gateway string `json:"gateway" yaml:"gateway" toml:"gateway" mapstructure:"gateway"`
}

// InitializeCloudInitNetworkIPv4WithDefaults initializes a CloudInitNetworkIPv4 struct with default values.
func InitializeCloudInitNetworkIPv4WithDefaults() *CloudInitNetworkIPv4 {
	return &CloudInitNetworkIPv4{
		AutoConfigure: true,
		Address:       "",
		Gateway:       "",
	}
}

// GetAutoConfigure returns the AutoConfigure field value.
func (cinipv4 *CloudInitNetworkIPv4) GetAutoConfigure() bool {
	return cinipv4.AutoConfigure
}

// GetAddress returns the Address field value.
func (cinipv4 *CloudInitNetworkIPv4) GetAddress() string {
	return cinipv4.Address
}

// GetGateway returns the Gateway field value.
func (cinipv4 *CloudInitNetworkIPv4) GetGateway() string {
	return cinipv4.Gateway
}

// IsConfigured returns true if the configuration is configured.
func (cinipv4 *CloudInitNetworkIPv4) IsConfigured() bool {
	if cinipv4.AutoConfigure {
		return true
	}
	return cinipv4.Address != "" && cinipv4.Gateway != ""
}
