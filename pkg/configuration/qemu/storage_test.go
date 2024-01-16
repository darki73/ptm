package qemu

import (
	"testing"
)

// TestInitializeQemuStorageWithDefaults tests if QemuStorage is correctly initialized with default values.
func TestInitializeQemuStorageWithDefaults(t *testing.T) {
	qemuStorage := InitializeQemuStorageWithDefaults()

	if qemuStorage.Name != "" {
		t.Errorf("Expected Name to be empty, got %s", qemuStorage.Name)
	}
	if qemuStorage.Resize != "" {
		t.Errorf("Expected Resize to be empty, got %s", qemuStorage.Resize)
	}
}

// TestGetStorage tests the GetStorage method.
func TestGetStorage(t *testing.T) {
	qemuStorage := &QemuStorage{Name: "test-storage"}
	if storage := qemuStorage.GetStorage(); storage != "test-storage" {
		t.Errorf("GetStorage() = %s, want %s", storage, "test-storage")
	}
}

// TestGetResize tests the GetResize method.
func TestGetResize(t *testing.T) {
	qemuStorage := &QemuStorage{Resize: "10G"}
	if resize := qemuStorage.GetResize(); resize != "10G" {
		t.Errorf("GetResize() = %s, want %s", resize, "10G")
	}
}

// TestIsSupposedToResize tests the IsSupposedToResize method.
func TestIsSupposedToResize(t *testing.T) {
	testCases := []struct {
		resize   string
		expected bool
	}{
		{"", false},
		{"10G", true},
		{"100M", true},
	}
	for _, tc := range testCases {
		t.Run(tc.resize, func(t *testing.T) {
			qemuStorage := &QemuStorage{Resize: tc.resize}
			if result := qemuStorage.IsSupposedToResize(); result != tc.expected {
				t.Errorf("IsSupposedToResize() = %v, want %v for resize value '%s'", result, tc.expected, tc.resize)
			}
		})
	}
}

// TestStorageIsConfigured tests the IsConfigured method.
func TestStorageIsConfigured(t *testing.T) {
	testCases := []struct {
		name     string
		resize   string
		expected bool
	}{
		{"", "", false},
		{"test-storage", "", true},
		{"", "10G", false},
		{"test-storage", "10G", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			qemuStorage := &QemuStorage{Name: tc.name, Resize: tc.resize}
			if result := qemuStorage.IsConfigured(); result != tc.expected {
				t.Errorf("IsConfigured() = %v, want %v for name '%s' and resize '%s'", result, tc.expected, tc.name, tc.resize)
			}
		})
	}
}
