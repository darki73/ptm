package cloud_init

import (
	"fmt"
	"github.com/darki73/ptm/pkg/utils"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	// ConfigurationSourcePrompt indicates that the configuration source is the prompt.
	ConfigurationSourcePrompt = "prompt"
	// ConfigurationSourceFlags indicates that the configuration source is the flags.
	ConfigurationSourceFlags = "flags"
	// ConfigurationSourceConfigurationFile indicates that the configuration source is the configuration file.
	ConfigurationSourceConfigurationFile = "configuration-file"
)

// CloudInit is a structure that contains the configuration for cloud-init.
type CloudInit struct {
	// username is the username to use for the cloud-init configuration.
	username string
	// password is the password to use for the cloud-init configuration.
	password string
	// keys is a list of SSH keys to use for the cloud-init configuration.
	keys []string
	// ipv4 is the IPv4 address to use for the cloud-init configuration.
	ipv4 string
	// ipv6 is the IPv6 address to use for the cloud-init configuration.
	ipv6 string
	// gateway4 is the IPv4 gateway to use for the cloud-init configuration.
	gateway4 string
	// gateway6 is the IPv6 gateway to use for the cloud-init configuration.
	gateway6 string
	// sshKeysTemporaryFilePath is the temporary file path for the SSH keys.
	sshKeysTemporaryFilePath string
	// configurationSource is the source of the configuration.
	configurationSource string
}

// NewCloudInitConfiguration creates a new cloud-init configuration.
func NewCloudInitConfiguration() *CloudInit {
	return &CloudInit{
		username:                 "",
		password:                 "",
		keys:                     []string{},
		ipv4:                     "dhcp",
		ipv6:                     "auto",
		gateway4:                 "",
		gateway6:                 "",
		sshKeysTemporaryFilePath: "/tmp/ptm-ssh-keys",
		configurationSource:      ConfigurationSourcePrompt,
	}
}

// GetUsername returns the username to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetUsername() string {
	return cloudInit.username
}

// SetUsername sets the username to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetUsername(username string) *CloudInit {
	cloudInit.username = username
	return cloudInit
}

// GetPassword returns the password to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetPassword() string {
	return cloudInit.password
}

// SetPassword sets the password to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetPassword(password string) *CloudInit {
	cloudInit.password = password
	return cloudInit
}

// GetKeys returns the list of SSH keys to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetKeys() []string {
	return cloudInit.keys
}

// HasKeys returns true if the cloud-init configuration has SSH keys.
func (cloudInit *CloudInit) HasKeys() bool {
	return len(cloudInit.keys) > 0
}

// IsValidSshKey checks if a given SSH key string is valid.
func (cloudInit *CloudInit) IsValidSshKey(key string) bool {
	var patterns = []string{
		`^ssh-rsa [A-Za-z0-9+\/]{100,}(={0,2})?( .+)?$`,
		`^ssh-dss [A-Za-z0-9+\/]{100,}(={0,2})?( .+)?$`,
		`^ecdsa-sha2-nistp256 [A-Za-z0-9+\/]{60,}(={0,2})?( .+)?$`,
		`^ecdsa-sha2-nistp384 [A-Za-z0-9+\/]{80,}(={0,2})?( .+)?$`,
		`^ecdsa-sha2-nistp521 [A-Za-z0-9+\/]{100,}(={0,2})?( .+)?$`,
		`^ssh-ed25519 [A-Za-z0-9+\/]{30,}(={0,2})?( .+)?$`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(key) {
			return true
		}
	}

	return false
}

// SetKeys sets the list of SSH keys to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetKeys(keys []string) *CloudInit {
	var keysSlice []string

	for _, key := range keys {
		if strings.HasSuffix(key, ".pub") {
			path, err := utils.ExpandHomeDir(key)
			if err != nil {
				log.Println(fmt.Sprintf("Failed to expand home directory: %s, error: %v", key, err))
				continue
			}

			fmt.Println(path, key)
			pubKey, err := os.ReadFile(path)
			if err != nil {
				if os.IsNotExist(err) {
					log.Println(fmt.Sprintf("File does not exist: %s", key))
				} else {
					log.Println(fmt.Sprintf("Failed to read file: %s, error: %v", key, err))
				}
				continue
			}
			key = strings.TrimSpace(string(pubKey))
		}

		if !cloudInit.IsValidSshKey(key) {
			log.Println(fmt.Sprintf("Invalid SSH key: %s, ignoring", key))
			continue
		}

		keysSlice = append(keysSlice, key)
	}
	cloudInit.keys = keysSlice
	return cloudInit
}

// GetIPv4 returns the IPv4 address to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetIPv4() string {
	return cloudInit.ipv4
}

// SetIPv4 sets the IPv4 address to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetIPv4(ipv4 string) *CloudInit {
	cloudInit.ipv4 = ipv4
	return cloudInit
}

// AutoConfigureIPv4 configures the IPv4 address to use for the cloud-init configuration.
func (cloudInit *CloudInit) AutoConfigureIPv4() *CloudInit {
	cloudInit.ipv4 = "dhcp"
	return cloudInit
}

// GetIPv6 returns the IPv6 address to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetIPv6() string {
	return cloudInit.ipv6
}

// SetIPv6 sets the IPv6 address to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetIPv6(ipv6 string) *CloudInit {
	cloudInit.ipv6 = ipv6
	return cloudInit
}

// AutoConfigureIPv6 configures the IPv6 address to use for the cloud-init configuration.
func (cloudInit *CloudInit) AutoConfigureIPv6() *CloudInit {
	cloudInit.ipv6 = "auto"
	return cloudInit
}

// GetIPv4Gateway returns the IPv4 gateway to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetIPv4Gateway() string {
	return cloudInit.gateway4
}

// SetIPv4Gateway sets the IPv4 gateway to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetIPv4Gateway(gateway4 string) *CloudInit {
	cloudInit.gateway4 = gateway4
	return cloudInit
}

// GetIPv6Gateway returns the IPv6 gateway to use for the cloud-init configuration.
func (cloudInit *CloudInit) GetIPv6Gateway() string {
	return cloudInit.gateway6
}

// SetIPv6Gateway sets the IPv6 gateway to use for the cloud-init configuration.
func (cloudInit *CloudInit) SetIPv6Gateway(gateway6 string) *CloudInit {
	cloudInit.gateway6 = gateway6
	return cloudInit
}

// GetSSHKeysTemporaryFilePath returns the temporary file path for the SSH keys.
func (cloudInit *CloudInit) GetSSHKeysTemporaryFilePath() string {
	return cloudInit.sshKeysTemporaryFilePath
}

// SetSSHKeysTemporaryFilePath sets the temporary file path for the SSH keys.
func (cloudInit *CloudInit) SetSSHKeysTemporaryFilePath(sshKeysTemporaryFilePath string) *CloudInit {
	cloudInit.sshKeysTemporaryFilePath = sshKeysTemporaryFilePath
	return cloudInit
}

// IsGateway4ConfigurationRequired returns true if the IPv4 gateway is required for the cloud-init configuration.
func (cloudInit *CloudInit) IsGateway4ConfigurationRequired() bool {
	return cloudInit.ipv4 != "dhcp"
}

// IsGateway6ConfigurationRequired returns true if the IPv6 gateway is required for the cloud-init configuration.
func (cloudInit *CloudInit) IsGateway6ConfigurationRequired() bool {
	return cloudInit.ipv6 != "auto"
}

// GetConfigurationSource returns the source of the configuration.
func (cloudInit *CloudInit) GetConfigurationSource() string {
	return cloudInit.configurationSource
}

// SetConfigurationSource sets the source of the configuration.
func (cloudInit *CloudInit) SetConfigurationSource(configurationSource string) *CloudInit {
	cloudInit.configurationSource = configurationSource
	return cloudInit
}

// IsConfigurationValid returns true if the configuration is valid.
func (cloudInit *CloudInit) IsConfigurationValid() (bool, error) {
	if cloudInit.IsGateway4ConfigurationRequired() {
		if cloudInit.gateway4 == "" && cloudInit.ipv4 != "dhcp" {
			return false, fmt.Errorf("cloud-init configuration requires IPv4 gateway when IPv4 address is not DHCP")
		}

		if cloudInit.ipv4 == "" && cloudInit.gateway4 != "" {
			return false, fmt.Errorf("cloud-init configuration requires IPv4 address when IPv4 gateway is set")
		}

	}

	if cloudInit.IsGateway6ConfigurationRequired() {
		if cloudInit.gateway6 == "" && cloudInit.ipv6 != "auto" {
			return false, fmt.Errorf("cloud-init configuration requires IPv6 gateway when IPv6 address is not auto")
		}

		if cloudInit.ipv6 == "" && cloudInit.gateway6 != "" {
			return false, fmt.Errorf("cloud-init configuration requires IPv6 address when IPv6 gateway is set")
		}
	}

	return true, nil
}
