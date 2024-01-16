package utils

import (
	"testing"
)

// TestBooleanToInteger tests the BooleanToInteger function.
func TestBooleanToInteger(t *testing.T) {
	tests := []struct {
		name     string
		value    bool
		expected int
	}{
		{"TrueToInteger", true, 1},
		{"FalseToInteger", false, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BooleanToInteger(test.value)
			if result != test.expected {
				t.Errorf("BooleanToInteger(%v) = %d; want %d", test.value, result, test.expected)
			}
		})
	}
}

// TestIntegerToBoolean tests the IntegerToBoolean function.
func TestIntegerToBoolean(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"ZeroToBoolean", 0, false},
		{"NonZeroToBoolean", 1, true},
		{"NegativeToBoolean", -1, true}, // you might want to test negative numbers as well
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntegerToBoolean(test.value)
			if result != test.expected {
				t.Errorf("IntegerToBoolean(%d) = %v; want %v", test.value, result, test.expected)
			}
		})
	}
}

// TestParseSizeForFromStringToInteger tests the ParseSizeForFromStringToInteger function.
func TestParseSizeForFromStringToInteger(t *testing.T) {
	tests := []struct {
		name         string
		sizeStr      string
		expectedNum  int64
		expectedUnit string
		expectError  bool
	}{
		{"ValidTerabytes", "1T", 1, "T", false},
		{"ValidGigabytes", "10G", 10, "G", false},
		{"ValidMegabytes", "512M", 512, "M", false},
		{"ValidKilobytes", "1024K", 1024, "K", false},
		{"ValidBytes", "8B", 8, "B", false},
		{"InvalidSizeString", "100Z", 0, "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			num, unit, err := ParseSizeForFromStringToInteger(test.sizeStr)
			if (err != nil) != test.expectError {
				t.Errorf("ParseSizeForFromStringToInteger(%s) unexpected error status: %v", test.sizeStr, err)
			}
			if num != test.expectedNum || unit != test.expectedUnit {
				t.Errorf("ParseSizeForFromStringToInteger(%s) = %v, %s; want %v, %s", test.sizeStr, num, unit, test.expectedNum, test.expectedUnit)
			}
		})
	}
}

// TestConvertFromStringToInteger tests the ConvertFromStringToInteger function.
func TestConvertFromStringToInteger(t *testing.T) {
	tests := []struct {
		name        string
		sizeStr     string
		targetUnit  string
		expected    int64
		expectError bool
	}{
		{"BytesToBytes", "1B", "B", 1, false},
		{"KilobytesToBytes", "1K", "B", 1024, false},
		{"KilobytesToKilobytes", "1K", "K", 1, false},
		{"MegabytesToBytes", "1M", "B", 1024 * 1024, false},
		{"MegabytesToKilobytes", "1M", "K", 1024, false},
		{"MegabytesToMegabytes", "1M", "M", 1, false},
		{"GigabytesToBytes", "1G", "B", 1024 * 1024 * 1024, false},
		{"GigabytesToKilobytes", "1G", "K", 1024 * 1024, false},
		{"GigabytesToMegabytes", "1G", "M", 1024, false},
		{"GigabytesToGigabytes", "1G", "G", 1, false},
		{"TerabytesToBytes", "1T", "B", 1024 * 1024 * 1024 * 1024, false},
		{"TerabytesToKilobytes", "1T", "K", 1024 * 1024 * 1024, false},
		{"TerabytesToMegabytes", "1T", "M", 1024 * 1024, false},
		{"TerabytesToGigabytes", "1T", "G", 1024, false},
		{"TerabytesToTerabytes", "1T", "T", 1, false},
		{"InvalidConversionBytesToKilobytes", "1B", "K", 0, true},
		{"InvalidConversionBytesToMegabytes", "1B", "M", 0, true},
		{"InvalidConversionBytesToGigabytes", "1B", "G", 0, true},
		{"InvalidConversionBytesToTerabytes", "1B", "T", 0, true},
		{"InvalidConversionKilobytesToMegabytes", "1K", "M", 0, true},
		{"InvalidConversionKilobytesToGigabytes", "1K", "G", 0, true},
		{"InvalidConversionKilobytesToTerabytes", "1K", "T", 0, true},
		{"InvalidConversionMegabytesToGigabytes", "1M", "G", 0, true},
		{"InvalidConversionMegabytesToTerabytes", "1M", "T", 0, true},
		{"InvalidConversionGigabytesToTerabytes", "1G", "T", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ConvertFromStringToInteger(test.sizeStr, test.targetUnit)
			if (err != nil) != test.expectError {
				t.Errorf("ConvertFromStringToInteger(%s, %s) unexpected error status: %v", test.sizeStr, test.targetUnit, err)
			}
			if result != test.expected {
				t.Errorf("ConvertFromStringToInteger(%s, %s) = %v; want %v", test.sizeStr, test.targetUnit, result, test.expected)
			}
		})
	}
}

// TestConvertToBytes tests the ConvertToBytes function.
func TestConvertToBytes(t *testing.T) {
	tests := []struct {
		name        string
		sizeStr     string
		expected    int64
		expectError bool
	}{
		{"BytesToBytes", "1B", 1, false},
		{"KilobytesToBytes", "1K", 1024, false},
		{"MegabytesToBytes", "1M", 1024 * 1024, false},
		{"GigabytesToBytes", "1G", 1024 * 1024 * 1024, false},
		{"TerabytesToBytes", "1T", 1024 * 1024 * 1024 * 1024, false},
		{"InvalidConversion", "1Z", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ConvertToBytes(test.sizeStr)
			if (err != nil) != test.expectError {
				t.Errorf("ConvertToBytes(%s) unexpected error status: %v", test.sizeStr, err)
			}
			if result != test.expected {
				t.Errorf("ConvertToBytes(%s) = %v; want %v", test.sizeStr, result, test.expected)
			}
		})
	}
}

// TestConvertToKilobytes tests the ConvertToKilobytes function.
func TestConvertToKilobytes(t *testing.T) {
	tests := []struct {
		name        string
		sizeStr     string
		expected    int64
		expectError bool
	}{
		{"BytesToKilobytes", "1B", 0, true},
		{"KilobytesToKilobytes", "1K", 1, false},
		{"MegabytesToKilobytes", "1M", 1024, false},
		{"GigabytesToKilobytes", "1G", 1024 * 1024, false},
		{"TerabytesToKilobytes", "1T", 1024 * 1024 * 1024, false},
		{"InvalidConversion", "1Z", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ConvertToKilobytes(test.sizeStr)
			if (err != nil) != test.expectError {
				t.Errorf("ConvertToKilobytes(%s) unexpected error status: %v", test.sizeStr, err)
			}
			if result != test.expected {
				t.Errorf("ConvertToKilobytes(%s) = %v; want %v", test.sizeStr, result, test.expected)
			}
		})
	}
}

// TestConvertToMegabytes tests the ConvertToMegabytes function.
func TestConvertToMegabytes(t *testing.T) {
	tests := []struct {
		name        string
		sizeStr     string
		expected    int64
		expectError bool
	}{
		{"BytesToMegabytes", "1B", 0, true},
		{"KilobytesToMegabytes", "1K", 0, true},
		{"MegabytesToMegabytes", "1M", 1, false},
		{"GigabytesToMegabytes", "1G", 1024, false},
		{"TerabytesToMegabytes", "1T", 1024 * 1024, false},
		{"InvalidConversion", "1Z", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ConvertToMegabytes(test.sizeStr)
			if (err != nil) != test.expectError {
				t.Errorf("ConvertToMegabytes(%s) unexpected error status: %v", test.sizeStr, err)
			}
			if result != test.expected {
				t.Errorf("ConvertToMegabytes(%s) = %v; want %v", test.sizeStr, result, test.expected)
			}
		})
	}
}

// TestConvertToGigabytes tests the ConvertToGigabytes function.
func TestConvertToGigabytes(t *testing.T) {
	tests := []struct {
		name        string
		sizeStr     string
		expected    int64
		expectError bool
	}{
		{"BytesToGigabytes", "1B", 0, true},
		{"KilobytesToGigabytes", "1K", 0, true},
		{"MegabytesToGigabytes", "1M", 0, true},
		{"GigabytesToGigabytes", "1G", 1, false},
		{"TerabytesToGigabytes", "1T", 1024, false},
		{"InvalidConversion", "1Z", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ConvertToGigabytes(test.sizeStr)
			if (err != nil) != test.expectError {
				t.Errorf("ConvertToGigabytes(%s) unexpected error status: %v", test.sizeStr, err)
			}
			if result != test.expected {
				t.Errorf("ConvertToGigabytes(%s) = %v; want %v", test.sizeStr, result, test.expected)
			}
		})
	}
}

// TestConvertToTerabytes tests the ConvertToTerabytes function.
func TestConvertToTerabytes(t *testing.T) {
	tests := []struct {
		name        string
		sizeStr     string
		expected    int64
		expectError bool
	}{
		{"BytesToTerabytes", "1B", 0, true},
		{"KilobytesToTerabytes", "1K", 0, true},
		{"MegabytesToTerabytes", "1M", 0, true},
		{"GigabytesToTerabytes", "1G", 0, true},
		{"TerabytesToTerabytes", "1T", 1, false},
		{"InvalidConversion", "1Z", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ConvertToTerabytes(test.sizeStr)
			if (err != nil) != test.expectError {
				t.Errorf("ConvertToTerabytes(%s) unexpected error status: %v", test.sizeStr, err)
			}
			if result != test.expected {
				t.Errorf("ConvertToTerabytes(%s) = %v; want %v", test.sizeStr, result, test.expected)
			}
		})
	}
}
