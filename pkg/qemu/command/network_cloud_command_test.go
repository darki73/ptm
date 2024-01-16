package command

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"reflect"
	"strconv"
	"testing"
)

// TestNewNetworkCloudCommandWithAutoConfigurationForIPv4andIPv6 tests the NewNetworkCloudCommand function with auto configuration for IPv4 and IPv6.
func TestNewNetworkCloudCommandWithAutoConfigurationForIPv4andIPv6(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.AutoConfigureIPv4().AutoConfigureIPv6()

	cmd := NewNetworkCloudCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewNetworkCloudCommandWithAutoConfigurationForIPv4andIPv6 did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--ipconfig0", "ip6=auto,ip=dhcp"}) {
		t.Errorf("TestNewNetworkCloudCommandWithAutoConfigurationForIPv4andIPv6 did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--ipconfig0", "ip6=auto,ip=dhcp"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}

// TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndAutoConfigurationForIPv6 tests the NewNetworkCloudCommand function with manual configuration for IPv4 and auto configuration for IPv6.
func TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndAutoConfigurationForIPv6(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.SetIPv4("127.0.0.1/24").SetIPv4Gateway("127.0.0.1").AutoConfigureIPv6()

	cmd := NewNetworkCloudCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndAutoConfigurationForIPv6 did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--ipconfig0", "ip6=auto,gw4=127.0.0.1,ip=127.0.0.1/24"}) {
		t.Errorf("TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndAutoConfigurationForIPv6 did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--ipconfig0", "ip6=auto,gw4=127.0.0.1,ip=127.0.0.1/24"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}

// TestNewNetworkCloudCommandWithAutoConfigurationForIPv4AndManualConfigurationForIPv6 tests the NewNetworkCloudCommand function with auto configuration for IPv4 and manual configuration for IPv6.
func TestNewNetworkCloudCommandWithAutoConfigurationForIPv4AndManualConfigurationForIPv6(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.AutoConfigureIPv4().SetIPv6("::1/64").SetIPv6Gateway("::1")

	cmd := NewNetworkCloudCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewNetworkCloudCommandWithAutoConfigurationForIPv4AndManualConfigurationForIPv6 did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--ipconfig0", "gw6=::1,ip6=::1/64,ip=dhcp"}) {
		t.Errorf("TestNewNetworkCloudCommandWithAutoConfigurationForIPv4AndManualConfigurationForIPv6 did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--ipconfig0", "gw6=::1,ip6=::1/64,ip=dhcp"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}

// TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndIPv6 tests the NewNetworkCloudCommand function with manual configuration for IPv4 and IPv6.
func TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndIPv6(t *testing.T) {
	identifier := 1
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.SetIPv4("127.0.0.1/24").SetIPv4Gateway("127.0.0.1").SetIPv6("::1/64").SetIPv6Gateway("::1")

	cmd := NewNetworkCloudCommand(identifier, cloudInit)

	if cmd.GetCommand() != qemuCommandSet || cmd.GetIdentifier() != identifier {
		t.Errorf("TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndIPv6 did not set command and identifier correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"--ipconfig0", "gw6=::1,ip6=::1/64,gw4=127.0.0.1,ip=127.0.0.1/24"}) {
		t.Errorf("TestNewNetworkCloudCommandWithManualConfigurationForIPv4AndIPv6 did not set arguments correctly")
	}

	expected := []string{qemuCommandSet, strconv.Itoa(identifier), "--ipconfig0", "gw6=::1,ip6=::1/64,gw4=127.0.0.1,ip=127.0.0.1/24"}
	result := cmd.BuildExecutionerCommand()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BuildExecutionerCommand returned %v, want %v", result, expected)
	}
}
