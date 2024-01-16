package cloud_init

import "testing"

// TestNewCloudInitConfiguration tests the NewCloudInitConfiguration function.
func TestNewCloudInitConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()

	if cloudInit == nil {
		t.Error("NewCloudInitConfiguration returned nil")
	}

	// Test default values
	if cloudInit.GetIPv4() != "dhcp" || cloudInit.GetIPv6() != "auto" || cloudInit.GetSSHKeysTemporaryFilePath() != "/tmp/ptm-ssh-keys" {
		t.Error("Default values for ipv4, ipv6 or sshKeysTemporaryFilePath are not set correctly")
	}
}

// TestSetAndGetUsername tests the SetUsername and GetUsername functions.
func TestSetAndGetUsername(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	username := "testuser"

	cloudInit.SetUsername(username)

	if cloudInit.GetUsername() != username {
		t.Errorf("GetUsername returned %v, want %v", cloudInit.GetUsername(), username)
	}
}

// TestSetAndGetPassword tests the SetPassword and GetPassword functions.
func TestSetAndGetPassword(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	password := "testpassword"

	cloudInit.SetPassword(password)

	if cloudInit.GetPassword() != password {
		t.Errorf("GetPassword returned %v, want %v", cloudInit.GetPassword(), password)
	}
}

// TestIsValidSshKey tests the IsValidSshKey function.
func TestIsValidSshKey(t *testing.T) {
	cloudInit := &CloudInit{} // Assuming CloudInit has a default constructor

	// Test cases
	testCases := []struct {
		name     string
		key      string
		expected bool
	}{
		// Valid cases
		{"Valid RSA", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCYID8wzSmmNYIyweQLk4TCJCGCCDAza+wl3PxpKslPV7LdZwdpfxo0Q2vlXsv4SN5wmJXJwwwE6ekGgMUFP4TtsGuyzXKxa1+jnBzc1+e/BtJ1AFdVAGzlUcqbZeHlnXfTlF+pnDrnXqG1jdJDsJZcgrZOK6uUQKMfbU/KoZM56w== user@example.com", true},
		{"Valid DSA", "ssh-dss AAAAB3NzaC1kc3MAAABBAPymgs6OEsq6Ju/M9xEOUm2weLBe3svNHrSiCPOuFheuAfNbkaR+bfY0E8XhLtCJm80TKs1Q2ZFRvcQ+5zdZLhcAAAAVAJYu3cw2nLqOuyYO5rahJtk0bjjFAAAAQGeEcbJ6nPRO6RpJxRR9samq8kTwWkNNZIaTHS0UJxueNQMLcf1z2heQabMuKTVjDhwgYjVNDaIKbEFuUL55TKQAAABBAN+ERyjUsXa0aO4tjZT+6HlAzW99K8yGEqp7rMGDz70t1+snD7/epOWxGCfasVnRn9qvPXR8clpOOnDinlBzuT4= user@example.com", true},
		{"Valid ECDSA 256", "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGkgZSuqj4Y0iWQP3YJCwR7gBMubTQsBewNMscVxYVJ3EA5rUsOWFDJ1J7MSDf9Iz6isc6YFxclvzCWi1/YE74o= user@example.com", true},
		{"Valid ECDSA 384", "ecdsa-sha2-nistp384 AAAAE2VjZHNhLXNoYTItbmlzdHAzODQAAAAIbmlzdHAzODQAAABhBGtpp+CYbXFW6HpY0eBJmOqD2bpX9WGZv3yAZ4No9JA616xZ4H5fAtcwxNgrBDrlCk1ajCnjEnNGflSpqsKaki6GJldgAqws9UuXY4KNPmfictdi+YCuv8VQRBaa4VfTsQ== user@example.com", true},
		{"Valid ECDSA 521", "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBACH5RBUOVjmexlGTX0hzQs06/kOCDngLuxLaAbyGaD4KCdwKBJX45mdTPQQRkd8YQ7QHlkg28mlvKeUbEM39Y6S8QDnp8y3XYzOXKFtg9jp6sunUWDeL+9S42n6imAxnofVTkFhI35tgAM0QD+MvATnw6ZZe2aRKvU/49yLpLXEKCzoLQ== user@example.com", true},
		{"Valid ED25519", "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEI4F/3yw1Jgok9b52nCDrtVffYtVNK4yqegGzeQ/NgS user@example.com", true},

		// Invalid cases
		{"Invalid RSA", "ssh-rsa invalidkeydata user@example.com", false},
		{"Invalid DSA", "ssh-dss invalidkeydata user@example.com", false},
		{"Invalid ECDSA 256", "ecdsa-sha2-nistp256 invalidkeydata user@example.com", false},
		{"Invalid ECDSA 384", "ecdsa-sha2-nistp384 invalidkeydata user@example.com", false},
		{"Invalid ECDSA 521", "ecdsa-sha2-nistp521 invalidkeydata user@example.com", false},
		{"Invalid ED25519", "ssh-ed25519 invalidkeydata user@example.com", false},
		{"Invalid Key Type", "ssh-invalid AAAAB3NzaC1yc2EAAAADAQABAAABAQ... user@example.com", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cloudInit.IsValidSshKey(tc.key)
			if result != tc.expected {
				t.Errorf("IsValidSshKey(%s) = %v, expected %v", tc.key, result, tc.expected)
			}
		})
	}
}

// TestHasKeys tests the HasKeys function.
func TestHasKeys(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()

	if cloudInit.HasKeys() {
		t.Errorf("HasKeys returned true, want false")
	}

	cloudInit.SetKeys([]string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCYID8wzSmmNYIyweQLk4TCJCGCCDAza+wl3PxpKslPV7LdZwdpfxo0Q2vlXsv4SN5wmJXJwwwE6ekGgMUFP4TtsGuyzXKxa1+jnBzc1+e/BtJ1AFdVAGzlUcqbZeHlnXfTlF+pnDrnXqG1jdJDsJZcgrZOK6uUQKMfbU/KoZM56w== user@example.com"})
	if !cloudInit.HasKeys() {
		t.Errorf("HasKeys returned false, want true")
	}
}

// TestSetAndGetKeys tests the SetKeys and GetKeys functions.
func TestSetAndGetKeys(t *testing.T) {
	// Define test cases
	tests := []struct {
		name          string
		inputKeys     []string
		expectedKeys  []string
		expectedCount int
	}{
		{
			name: "Valid RSA and ED25519 keys",
			inputKeys: []string{
				"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCYID8wzSmmNYIyweQLk4TCJCGCCDAza+wl3PxpKslPV7LdZwdpfxo0Q2vlXsv4SN5wmJXJwwwE6ekGgMUFP4TtsGuyzXKxa1+jnBzc1+e/BtJ1AFdVAGzlUcqbZeHlnXfTlF+pnDrnXqG1jdJDsJZcgrZOK6uUQKMfbU/KoZM56w== user@example.com",
				"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEI4F/3yw1Jgok9b52nCDrtVffYtVNK4yqegGzeQ/NgS user@example.com",
			},
			expectedKeys: []string{
				"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCYID8wzSmmNYIyweQLk4TCJCGCCDAza+wl3PxpKslPV7LdZwdpfxo0Q2vlXsv4SN5wmJXJwwwE6ekGgMUFP4TtsGuyzXKxa1+jnBzc1+e/BtJ1AFdVAGzlUcqbZeHlnXfTlF+pnDrnXqG1jdJDsJZcgrZOK6uUQKMfbU/KoZM56w== user@example.com",
				"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEI4F/3yw1Jgok9b52nCDrtVffYtVNK4yqegGzeQ/NgS user@example.com",
			},
			expectedCount: 2,
		},
		{
			name: "Invalid key formats",
			inputKeys: []string{
				"invalid-ssh-key",
				"ssh-rsa invalidkeyformat",
			},
			expectedKeys:  nil,
			expectedCount: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cloudInit := NewCloudInitConfiguration()
			cloudInit.SetKeys(tc.inputKeys)

			retrievedKeys := cloudInit.GetKeys()
			if len(retrievedKeys) != tc.expectedCount {
				t.Errorf("%s: expected %d keys, got %d", tc.name, tc.expectedCount, len(retrievedKeys))
			}

			for i, key := range retrievedKeys {
				if key != tc.expectedKeys[i] {
					t.Errorf("%s: key %d does not match expected value. Got %s, want %s", tc.name, i, key, tc.expectedKeys[i])
				}
			}
		})
	}
}

// TestSetAndGetIPv4ManualConfiguration tests the SetIPv4 and GetIPv4 functions with manual configuration.
func TestSetAndGetIPv4ManualConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv4 := "127.0.0.1/24"

	cloudInit.SetIPv4(ipv4)

	if cloudInit.GetIPv4() != ipv4 {
		t.Errorf("GetIPv4 returned %v, want %v", cloudInit.GetIPv4(), ipv4)
	}
}

// TestSetAndGetIPv4AutoConfiguration tests the SetIPv4 and GetIPv4 functions with DHCP configuration.
func TestSetAndGetIPv4AutoConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv4 := "dhcp"

	cloudInit.AutoConfigureIPv4()

	if cloudInit.GetIPv4() != ipv4 {
		t.Errorf("GetIPv4 returned %v, want %v", cloudInit.GetIPv4(), ipv4)
	}
}

// TestSetAndGetIPv6ManualConfiguration tests the SetIPv6 and GetIPv6 functions with manual configuration.
func TestSetAndGetIPv6ManualConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv6 := "::1/64"

	cloudInit.SetIPv6(ipv6)

	if cloudInit.GetIPv6() != ipv6 {
		t.Errorf("GetIPv6 returned %v, want %v", cloudInit.GetIPv6(), ipv6)
	}
}

// TestSetAndGetIPv6AutoConfiguration tests the SetIPv6 and GetIPv6 functions with DHCP configuration.
func TestSetAndGetIPv6AutoConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv6 := "auto"

	cloudInit.AutoConfigureIPv6()

	if cloudInit.GetIPv6() != ipv6 {
		t.Errorf("GetIPv6 returned %v, want %v", cloudInit.GetIPv6(), ipv6)
	}
}

// TestSetAndGetIPv4Gateway tests the SetIPv4Gateway and GetIPv4Gateway functions.
func TestSetAndGetIPv4Gateway(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv4Gateway := "10.10.10.1"

	cloudInit.SetIPv4Gateway(ipv4Gateway)

	if cloudInit.GetIPv4Gateway() != ipv4Gateway {
		t.Errorf("GetIPv4Gateway returned %v, want %v", cloudInit.GetIPv4Gateway(), ipv4Gateway)
	}
}

// TestIsGateway4ConfigurationRequiredWithAutoConfiguration tests the IsGateway4ConfigurationRequired function with DHCP configuration.
func TestIsGateway4ConfigurationRequiredWithAutoConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()

	cloudInit.AutoConfigureIPv4()

	if cloudInit.IsGateway4ConfigurationRequired() {
		t.Errorf("IsGateway4ConfigurationRequired returned true, want false")
	}
}

// TestIsGateway4ConfigurationRequiredWithManualConfiguration tests the IsGateway4ConfigurationRequired function with manual configuration.
func TestIsGateway4ConfigurationRequiredWithManualConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv4 := "127.0.0.1/24"

	cloudInit.SetIPv4(ipv4)

	if !cloudInit.IsGateway4ConfigurationRequired() {
		t.Errorf("IsGateway4ConfigurationRequired returned false, want true")
	}
}

// TestSetAndGetIPv6Gateway tests the SetIPv6Gateway and GetIPv6Gateway functions.
func TestSetAndGetIPv6Gateway(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv6Gateway := "fe80::1"

	cloudInit.SetIPv6Gateway(ipv6Gateway)

	if cloudInit.GetIPv6Gateway() != ipv6Gateway {
		t.Errorf("GetIPv6Gateway returned %v, want %v", cloudInit.GetIPv6Gateway(), ipv6Gateway)
	}
}

// TestIsGateway6ConfigurationRequiredWithAutoConfiguration tests the IsGateway6ConfigurationRequired function with DHCP configuration.
func TestIsGateway6ConfigurationRequiredWithAutoConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()

	cloudInit.AutoConfigureIPv6()

	if cloudInit.IsGateway6ConfigurationRequired() {
		t.Errorf("IsGateway6ConfigurationRequired returned true, want false")
	}
}

// TestIsGateway6ConfigurationRequiredWithManualConfiguration tests the IsGateway6ConfigurationRequired function with manual configuration.
func TestIsGateway6ConfigurationRequiredWithManualConfiguration(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	ipv6 := "::1/64"

	cloudInit.SetIPv6(ipv6)

	if !cloudInit.IsGateway6ConfigurationRequired() {
		t.Errorf("IsGateway6ConfigurationRequired returned false, want true")
	}
}

// TestSetAndGetSSHKeysTemporaryFilePath tests the SetSSHKeysTemporaryFilePath and GetSSHKeysTemporaryFilePath functions.
func TestSetAndGetSSHKeysTemporaryFilePath(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	sshKeysTemporaryFilePath := "/tmp/test-ssh-keys"

	cloudInit.SetSSHKeysTemporaryFilePath(sshKeysTemporaryFilePath)

	if cloudInit.GetSSHKeysTemporaryFilePath() != sshKeysTemporaryFilePath {
		t.Errorf("GetSSHKeysTemporaryFilePath returned %v, want %v", cloudInit.GetSSHKeysTemporaryFilePath(), sshKeysTemporaryFilePath)
	}
}

// TestSetAndGetConfigurationSource tests the SetConfigurationSource and GetConfigurationSource functions.
func TestSetAndGetConfigurationSource(t *testing.T) {
	cloudInit := NewCloudInitConfiguration()
	configurationSource := "test"

	cloudInit.SetConfigurationSource(configurationSource)

	if cloudInit.GetConfigurationSource() != configurationSource {
		t.Errorf("GetConfigurationSource returned %v, want %v", cloudInit.GetConfigurationSource(), configurationSource)
	}
}

// TestIsConfigurationValid tests the IsConfigurationValid function.
func TestIsConfigurationValid(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*CloudInit)
		wantErr  bool
		errorMsg string
	}{
		{
			name: "Valid auto IPv4 and IPv6 configuration",
			setup: func(ci *CloudInit) {
				ci.AutoConfigureIPv4()
				ci.AutoConfigureIPv6()
			},
			wantErr: false,
		},
		{
			name: "Valid manual IPv4 configuration and auto IPv6",
			setup: func(ci *CloudInit) {
				ci.SetIPv4("10.10.10.10/24")
				ci.SetIPv4Gateway("10.10.10.1")
				ci.AutoConfigureIPv6()
			},
			wantErr: false,
		},
		{
			name: "Valid manual IPv6 configuration and auto IPv4",
			setup: func(ci *CloudInit) {
				ci.SetIPv6("fe80::1/64")
				ci.SetIPv6Gateway("fe80::1")
				ci.AutoConfigureIPv4()
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cloudInit := NewCloudInitConfiguration()
			tc.setup(cloudInit)

			valid, err := cloudInit.IsConfigurationValid()
			if (err != nil) != tc.wantErr {
				t.Errorf("IsConfigurationValid() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !valid && !tc.wantErr {
				t.Errorf("IsConfigurationValid() = %v, want %v", valid, true)
			}
			if tc.wantErr && err.Error() != tc.errorMsg {
				t.Errorf("Expected error message '%s', got '%s'", tc.errorMsg, err.Error())
			}
		})
	}
}
