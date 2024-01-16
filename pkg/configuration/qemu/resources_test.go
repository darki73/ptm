package qemu

import (
	"testing"
)

// TestInitializeQemuResourcesWithDefaults tests if the QemuResources are correctly initialized with default values.
func TestInitializeQemuResourcesWithDefaults(t *testing.T) {
	qemuResources := InitializeQemuResourcesWithDefaults()

	if qemuResources.CpuType != "host" {
		t.Errorf("Expected CpuType to be 'host', got %s", qemuResources.CpuType)
	}
}

// TestGetCores tests the GetCores method.
func TestGetCores(t *testing.T) {
	qemuResources := &QemuResources{Cores: 4}
	if cores := qemuResources.GetCores(); cores != 4 {
		t.Errorf("GetCores() = %d, want %d", cores, 4)
	}
}

// TestGetMemory tests the GetMemory method with various memory sizes.
func TestGetMemory(t *testing.T) {
	testCases := []struct {
		name      string
		memory    string
		expected  int64
		expectErr bool
	}{
		{"Numeric Memory", "1024", 1024, false},
		{"Memory in G", "2G", 2048, false},
		{"Memory in M", "512M", 512, false},
		{"Memory in T", "1T", 1048576, false},
		{"Invalid Memory", "invalid", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qemuResources := &QemuResources{Memory: tc.memory}
			got, err := qemuResources.GetMemory()

			if (err != nil) != tc.expectErr {
				t.Fatalf("GetMemory() error = %v, expectErr %v", err, tc.expectErr)
			}

			if got != tc.expected {
				t.Errorf("GetMemory() = %d, want %d", got, tc.expected)
			}
		})
	}
}

// TestGetCpuType tests the GetCpuType method.
func TestGetCpuType(t *testing.T) {
	qemuResources := &QemuResources{CpuType: "custom-cpu"}
	if cpuType := qemuResources.GetCpuType(); cpuType != "custom-cpu" {
		t.Errorf("GetCpuType() = %s, want %s", cpuType, "custom-cpu")
	}
}

// TestResourcesIsConfigured tests the IsConfigured method.
func TestResourcesIsConfigured(t *testing.T) {
	testCases := []struct {
		name          string
		qemuResources *QemuResources
		expected      bool
	}{
		{"Not Configured", &QemuResources{}, false},
		{"Configured", &QemuResources{Cores: 1, Memory: "1G", CpuType: "host"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.qemuResources.IsConfigured(); got != tc.expected {
				t.Errorf("IsConfigured() = %v, want %v", got, tc.expected)
			}
		})
	}
}
