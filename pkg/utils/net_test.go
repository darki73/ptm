package utils

import (
	"testing"
)

func TestIsValidIPWithSubnet(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"192.168.1.1/24", true},
		{"10.10.0.20/22", true},
		{"invalid_ip", false},
		{"192.168.1.1", false},
	}

	for _, tc := range testCases {
		result := IsValidIPWithSubnet(tc.input)
		if result != tc.expected {
			t.Errorf("IsValidIPWithSubnet(%s) = %v, want %v", tc.input, result, tc.expected)
		}
	}
}

func TestIsValidIP(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"192.168.1.1", true},
		{"10.10.0.20", true},
		{"invalid_ip", false},
		{"192.168.1.1/24", false},
	}

	for _, tc := range testCases {
		result := IsValidIP(tc.input)
		if result != tc.expected {
			t.Errorf("IsValidIP(%s) = %v, want %v", tc.input, result, tc.expected)
		}
	}
}

func TestIsInSameNetwork(t *testing.T) {
	testCases := []struct {
		ip       string
		network  string
		expected bool
	}{
		{"10.10.1.1", "10.10.0.2/22", true},
		{"192.168.1.1", "192.168.1.0/24", true},
		{"192.168.2.1", "192.168.1.0/24", false},
		{"invalid_ip", "192.168.1.0/24", false},
		{"192.168.1.1", "invalid_network", false},
	}

	for _, tc := range testCases {
		result := IsInSameNetwork(tc.ip, tc.network)
		if result != tc.expected {
			t.Errorf("IsInSameNetwork(%s, %s) = %v, want %v", tc.ip, tc.network, result, tc.expected)
		}
	}
}
