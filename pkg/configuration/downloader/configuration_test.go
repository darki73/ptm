package downloader

import (
	"testing"
)

// TestInitializeWithDefaults tests if InitializeWithDefaults function sets the default values correctly.
func TestInitializeWithDefaults(t *testing.T) {
	expectedSaveTo := "/tmp"
	config := InitializeWithDefaults()
	if config.SaveTo != expectedSaveTo {
		t.Errorf("InitializeWithDefaults() SaveTo = %v, want %v", config.SaveTo, expectedSaveTo)
	}
}

// TestGetSaveTo tests if GetSaveTo method returns the correct value.
func TestGetSaveTo(t *testing.T) {
	expectedSaveTo := "/tmp"
	config := InitializeWithDefaults()
	if saveTo := config.GetSaveTo(); saveTo != expectedSaveTo {
		t.Errorf("GetSaveTo() = %v, want %v", saveTo, expectedSaveTo)
	}
}
