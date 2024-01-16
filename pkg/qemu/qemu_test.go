package qemu

import (
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"testing"
)

// TestNewQemuConfiguration tests the NewQemuConfiguration method.
func TestNewQemuConfiguration(t *testing.T) {
	qemu := NewQemuConfiguration()

	if qemu == nil {
		t.Error("NewQemuConfiguration returned nil")
	}
}

// TestSetAndGetIdentifier tests the SetIdentifier and GetIdentifier methods.
func TestSetAndGetIdentifier(t *testing.T) {
	qemu := NewQemuConfiguration()
	identifier := 123

	qemu.SetIdentifier(identifier)

	if qemu.GetIdentifier() != identifier {
		t.Errorf("GetIdentifier returned %v, want %v", qemu.GetIdentifier(), identifier)
	}
}

// TestSetAndGetName tests the SetName and GetName methods.
func TestSetAndGetName(t *testing.T) {
	qemu := NewQemuConfiguration()
	name := "testVM"

	qemu.SetName(name)

	if qemu.GetName() != name {
		t.Errorf("GetName returned %v, want %v", qemu.GetName(), name)
	}
}

// TestSetAndGetCores tests the SetCores and GetCores methods.
func TestSetAndGetCores(t *testing.T) {
	qemu := NewQemuConfiguration()
	cores := 4

	qemu.SetCores(cores)

	if qemu.GetCores() != cores {
		t.Errorf("GetCores returned %v, want %v", qemu.GetCores(), cores)
	}
}

// TestSetAndGetMemory tests the SetMemory and GetMemory methods.
func TestSetAndGetMemory(t *testing.T) {
	qemu := NewQemuConfiguration()
	memory := 1024

	qemu.SetMemory(memory)

	if qemu.GetMemory() != memory {
		t.Errorf("GetMemory returned %v, want %v", qemu.GetMemory(), memory)
	}
}

// TestSetAndGetCpuType tests the SetCPUType and GetCPUType methods.
func TestSetAndGetCpuType(t *testing.T) {
	qemu := NewQemuConfiguration()
	cpuType := "example"

	qemu.SetCpuType(cpuType)

	if qemu.GetCpuType() != cpuType {
		t.Errorf("GetCpuType returned %v, want %v", qemu.GetCpuType(), cpuType)
	}
}

// TestSetAndGetNetworkDriver tests the SetNetworkDriver and GetNetworkDriver methods.
func TestSetAndGetNetworkDriver(t *testing.T) {
	qemu := NewQemuConfiguration()
	networkDriver := "example"

	qemu.SetNetworkDriver(networkDriver)

	if qemu.GetNetworkDriver() != networkDriver {
		t.Errorf("GetNetworkDriver returned %v, want %v", qemu.GetNetworkDriver(), networkDriver)
	}
}

// TestSetAndGetNetworkBridge tests the SetNetworkBridge and GetNetworkBridge methods.
func TestSetAndGetNetworkBridge(t *testing.T) {
	qemu := NewQemuConfiguration()
	networkBridge := "example"

	qemu.SetNetworkBridge(networkBridge)

	if qemu.GetNetworkBridge() != networkBridge {
		t.Errorf("GetNetworkBridge returned %v, want %v", qemu.GetNetworkBridge(), networkBridge)
	}
}

// TestSetAndGetStorage tests the SetStorage and GetStorage methods.
func TestSetAndGetStorage(t *testing.T) {
	qemu := NewQemuConfiguration()
	storage := "local-lvm"

	qemu.SetStorage(storage)

	if qemu.GetStorage() != storage {
		t.Errorf("GetStorage returned %v, want %v", qemu.GetStorage(), storage)
	}
}

// TestSetAndGetStorageSize tests the SetStorageSize and GetStorageSize methods.
func TestSetAndGetStorageSize(t *testing.T) {
	qemu := NewQemuConfiguration()
	storageSize := int64(1024)

	qemu.SetStorageSize(storageSize)

	if qemu.GetStorageSize() != storageSize {
		t.Errorf("GetStorageSize returned %v, want %v", qemu.GetStorageSize(), storageSize)
	}
}

// TestSetAndGetImage tests the SetImage and GetImage methods.
func TestSetAndGetImage(t *testing.T) {
	qemu := NewQemuConfiguration()
	image := "/etc/ptm/images/ubuntu-22.04-cloudimg-amd64.img"

	qemu.SetImage(image)

	if qemu.GetImage() != image {
		t.Errorf("GetImage returned %v, want %v", qemu.GetImage(), image)
	}
}

// TestSetAndGetNewImageSize tests the SetNewImageSize and GetNewImageSize methods.
func TestSetAndGetNewImageSize(t *testing.T) {
	qemu := NewQemuConfiguration()
	newImageSize := int64(10 * 1024 * 1024)

	qemu.SetNewImageSize(newImageSize)

	if qemu.GetNewImageSize() != newImageSize {
		t.Errorf("GetNewImageSize returned %v, want %v", qemu.GetNewImageSize(), newImageSize)
	}
}

// TestSetAndGetNewImageSizeAsString tests the SetNewImageSizeAsString and GetNewImageSizeAsString methods.
func TestSetAndGetNewImageSizeAsString(t *testing.T) {
	qemu := NewQemuConfiguration()
	newImageSize := "10G"

	qemu.SetNewImageSizeAsString(newImageSize)

	if qemu.GetNewImageSizeAsString() != newImageSize {
		t.Errorf("GetNewImageSizeAsString returned %v, want %v", qemu.GetNewImageSizeAsString(), newImageSize)
	}
}

// TestIsResizingRequiredWhenEmpty tests the IsResizingRequired method when the new image size is empty.
func TestIsResizingRequiredWhenEmpty(t *testing.T) {
	var imageSize float64 = 2.2 * 1024 * 1024
	qemu := NewQemuConfiguration().SetImageSize(int64(imageSize))

	if qemu.IsResizingRequired() {
		t.Errorf("IsResizingRequired returned true, want false")
	}
}

// TestIsResizingRequiredWhenSmaller tests the IsResizingRequired method when the new image size is smaller than the current image size.
func TestIsResizingRequiredWhenSmaller(t *testing.T) {
	var imageSize float64 = 2.2 * 1024 * 1024

	qemu := NewQemuConfiguration()
	qemu.
		SetImage("/etc/ptm/images/ubuntu-22.04-cloudimg-amd64.img").
		SetImageSize(int64(imageSize)).
		SetNewImageSize(2 * 1024 * 1024)

	if qemu.IsResizingRequired() {
		t.Errorf("IsResizingRequired returned true, want false")
	}
}

// TestIsResizingRequiredWhenGreater tests the IsResizingRequired method when the new image size is greater than the current image size.
func TestIsResizingRequiredWhenGreater(t *testing.T) {
	var imageSize float64 = 2.2 * 1024 * 1024

	qemu := NewQemuConfiguration()
	qemu.
		SetImage("/etc/ptm/images/ubuntu-22.04-cloudimg-amd64.img").
		SetImageSize(int64(imageSize)).
		SetNewImageSize(3 * 1024 * 1024)

	if !qemu.IsResizingRequired() {
		t.Errorf("IsResizingRequired returned false, want true")
	}
}

// TestSetAndGetCloudInit tests the SetCloudInit and GetCloudInit methods.
func TestSetAndGetCloudInit(t *testing.T) {
	qemu := NewQemuConfiguration()
	cloudInit := ci.NewCloudInitConfiguration()
	cloudInit.SetUsername("user").SetPassword("pass")

	// TODO: add tests for the other CloudInit configuration options

	qemu.SetCloudInit(cloudInit)

	if qemu.GetCloudInit().GetUsername() != "user" || qemu.GetCloudInit().GetPassword() != "pass" {
		t.Errorf("GetCloudInit did not return the expected CloudInit configuration")
	}
}

// TestSetAndGetConfigurationSource tests the SetConfigurationSource and GetConfigurationSource methods.
func TestSetAndGetConfigurationSource(t *testing.T) {
	qemu := NewQemuConfiguration()
	configurationSource := "test"

	qemu.SetConfigurationSource(configurationSource)

	if qemu.GetConfigurationSource() != configurationSource {
		t.Errorf("GetConfigurationSource returned %v, want %v", qemu.GetConfigurationSource(), configurationSource)
	}
}

func TestIsConfigurationValid(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(*Qemu)
		wantValid bool
		wantErr   bool
		errorMsg  string
	}{
		{
			name:      "Qemu::IsConfigurationValid - Missing identifier",
			setup:     func(q *Qemu) {},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing virtual machine identifier",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing name",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing virtual machine name",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing cores",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "invalid number of cores assigned to virtual machine",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing memory",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "invalid amount of memory assigned to virtual machine",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing CPU type",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
				q.SetMemory(1024)
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing cpu type for virtual machine",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing network driver",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
				q.SetMemory(1024)
				q.SetCpuType("host")
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing network driver for virtual machine",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing network bridge",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
				q.SetMemory(1024)
				q.SetCpuType("host")
				q.SetNetworkDriver("virtio")
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing network bridge for virtual machine",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing storage",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
				q.SetMemory(1024)
				q.SetCpuType("host")
				q.SetNetworkDriver("virtio")
				q.SetNetworkBridge("vmbr0")
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing storage for virtual machine",
		},
		{
			name: "Qemu::IsConfigurationValid - Missing image",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
				q.SetMemory(1024)
				q.SetCpuType("host")
				q.SetNetworkDriver("virtio")
				q.SetNetworkBridge("vmbr0")
				q.SetStorage("local-lvm")
			},
			wantValid: false,
			wantErr:   true,
			errorMsg:  "missing image for virtual machine",
		},
		{
			name: "Complete configuration",
			setup: func(q *Qemu) {
				q.SetIdentifier(123)
				q.SetName("test")
				q.SetCores(4)
				q.SetMemory(1024)
				q.SetCpuType("host")
				q.SetNetworkDriver("virtio")
				q.SetNetworkBridge("vmbr0")
				q.SetStorage("local-lvm")
				q.SetImage("/etc/ptm/images/ubuntu-22.04-cloudimg-amd64.img")
				cloudInit := ci.NewCloudInitConfiguration()
				cloudInit.SetIPv4("10.10.10.10/24")
				cloudInit.SetIPv4Gateway("10.10.10.1")
				cloudInit.SetIPv6("2001:db8::1/64")
				cloudInit.SetIPv6Gateway("2001:db8::1")
				q.SetCloudInit(cloudInit)
			},
			wantValid: true,
			wantErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			qemu := NewQemuConfiguration()
			tc.setup(qemu)

			valid, err := qemu.IsConfigurationValid()
			if (err != nil) != tc.wantErr {
				t.Errorf("%s: IsConfigurationValid() error = %v, wantErr %v", tc.name, err, tc.wantErr)
				return
			}
			if valid != tc.wantValid {
				t.Errorf("%s: IsConfigurationValid() = %v, want %v", tc.name, valid, tc.wantValid)
			}
			if tc.wantErr && err.Error() != tc.errorMsg {
				t.Errorf("%s: Expected error message '%s', got '%s'", tc.name, tc.errorMsg, err.Error())
			}
		})
	}
}
