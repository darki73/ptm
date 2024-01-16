package utils

import (
	"runtime"
	"testing"
)

// TestGetCoreCount tests the GetCoreCount function
func TestGetCoreCount(t *testing.T) {
	expectedCores := runtime.NumCPU()
	actualCores := GetCoreCount()

	if actualCores != expectedCores {
		t.Errorf("TestGetCoreCount failed: expected %d cores, got %d", expectedCores, actualCores)
	}
}

// TestGetTotalMemory tests the GetTotalMemory function
func TestGetTotalMemory(t *testing.T) {
	totalMemory := GetTotalMemory()

	if totalMemory <= 0 {
		t.Errorf("TestGetTotalMemory failed: expected total memory to be greater than zero")
	}
}
