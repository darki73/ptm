package utils

import "testing"

// TestIsNumeric tests the IsNumeric function.
func TestIsNumeric(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"Numeric", "1", true},
		{"NonNumeric", "a", false},
		{"SemiNumeric", "1a", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsNumeric(test.value)
			if result != test.expected {
				t.Errorf("IsNumeric(%s) = %v; want %v", test.value, result, test.expected)
			}
		})
	}
}
