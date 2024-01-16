package distributions

import (
	"testing"
)

// TestNewDistributions tests the NewDistributions function.
func TestNewDistributions(t *testing.T) {
	distributions, err := NewDistributions(ubuntuTestInitialBaseImage)

	if err != nil {
		t.Errorf("NewDistributions() returned an error: %v", err)
	}

	if distributions == nil {
		t.Error("NewDistributions() returned nil")
	}

	if distributions.GetActiveDistribution() == nil {
		t.Error("Active distribution is nil")
	}
}

// TestIsDistributionSupported tests the IsDistributionSupported function.
func TestIsDistributionSupported(t *testing.T) {
	distributions, err := NewDistributions(ubuntuTestInitialBaseImage)

	if err != nil {
		t.Errorf("NewDistributions() returned an error: %v", err)
	}

	if !distributions.IsDistributionSupported("ubuntu") {
		t.Error("Expected 'ubuntu' to be supported, but it's not.")
	}

	if distributions.IsDistributionSupported("nonexistent") {
		t.Error("Expected 'nonexistent' to not be supported, but it is.")
	}
}

// TestGetDistributionByName tests the GetDistributionByName function.
func TestGetDistributionByName(t *testing.T) {
	distributions, err := NewDistributions(ubuntuTestInitialBaseImage)

	if err != nil {
		t.Errorf("NewDistributions() returned an error: %v", err)
	}

	ubuntu := distributions.GetDistributionByName("ubuntu")
	if ubuntu == nil {
		t.Error("Expected 'ubuntu' to be a valid distribution, but it's nil.")
	}

	nonexistent := distributions.GetDistributionByName("nonexistent")
	if nonexistent != nil {
		t.Error("Expected 'nonexistent' to be nil, but it's not.")
	}
}

// TestGetDistribution tests the GetDistribution function.
func TestGetDistribution(t *testing.T) {
	distributions, err := NewDistributions(ubuntuTestInitialBaseImage)

	if err != nil {
		t.Errorf("NewDistributions() returned an error: %v", err)
	}

	ubuntu := distributions.GetDistribution("ubuntu")
	if ubuntu == nil {
		t.Error("Expected 'ubuntu' to be a valid distribution, but it's nil.")
	}

	nonexistent := distributions.GetDistribution("nonexistent")
	if nonexistent != nil {
		t.Error("Expected 'nonexistent' to be nil, but it's not.")
	}
}

// TestGetActiveDistribution tests the GetActiveDistribution function.
func TestGetActiveDistribution(t *testing.T) {
	distributions, err := NewDistributions(ubuntuTestInitialBaseImage)

	if err != nil {
		t.Errorf("NewDistributions() returned an error: %v", err)
	}

	activeDistribution := distributions.GetActiveDistribution()
	if activeDistribution == nil {
		t.Error("Active distribution is nil")
	}
}
