package utils

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Byte     = 1
	Kilobyte = 1024 * Byte
	Megabyte = 1024 * Kilobyte
	Gigabyte = 1024 * Megabyte
	Terabyte = 1024 * Gigabyte
)

// BooleanToInteger converts a boolean value to an integer.
func BooleanToInteger(value bool) int {
	if value {
		return 1
	}
	return 0
}

// IntegerToBoolean converts an integer value to a boolean.
func IntegerToBoolean(value int) bool {
	if value != 0 {
		return true
	}
	return false
}

// ParseSizeForFromStringToInteger parses a size string to an integer.
func ParseSizeForFromStringToInteger(sizeStr string) (int64, string, error) {
	numStr := strings.TrimRight(sizeStr, "TtGgMmKkBb")
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0, "", err
	}

	unit := sizeStr[len(numStr):]
	return num, strings.ToUpper(unit), nil
}

// ConvertFromStringToInteger converts a size string to an integer.
func ConvertFromStringToInteger(sizeStr string, targetUnit string) (int64, error) {
	num, unit, err := ParseSizeForFromStringToInteger(sizeStr)
	if err != nil {
		return 0, err
	}

	switch targetUnit {
	case "T":
		switch unit {
		case "T":
			return num, nil
		default:
			return 0, fmt.Errorf("conversion to units which will produce floating point number is not possible")
		}
	case "G":
		switch unit {
		case "T":
			return num * Terabyte / Gigabyte, nil
		case "G":
			return num, nil
		default:
			return 0, fmt.Errorf("conversion to units which will produce floating point number is not possible")
		}
	case "M":
		switch unit {
		case "T":
			return num * Terabyte / Megabyte, nil
		case "G":
			return num * Gigabyte / Megabyte, nil
		case "M":
			return num, nil
		default:
			return 0, fmt.Errorf("conversion to units which will produce floating point number is not possible")
		}
	case "K":
		switch unit {
		case "T":
			return num * Terabyte / Kilobyte, nil
		case "G":
			return num * Gigabyte / Kilobyte, nil
		case "M":
			return num * Megabyte / Kilobyte, nil
		case "K":
			return num, nil
		default:
			return 0, fmt.Errorf("conversion to units which will produce floating point number is not possible")
		}
	case "B":
		switch unit {
		case "T":
			return num * Terabyte, nil
		case "G":
			return num * Gigabyte, nil
		case "M":
			return num * Megabyte, nil
		case "K":
			return num * Kilobyte, nil
		case "B":
			return num, nil
		}
	}

	return 0, fmt.Errorf("invalid conversion")
}

// ConvertToBytes converts a size string to bytes.
func ConvertToBytes(sizeStr string) (int64, error) {
	return ConvertFromStringToInteger(sizeStr, "B")
}

// ConvertToKilobytes converts a size string to kilobytes.
func ConvertToKilobytes(sizeStr string) (int64, error) {
	return ConvertFromStringToInteger(sizeStr, "K")
}

// ConvertToMegabytes converts a size string to megabytes.
func ConvertToMegabytes(sizeStr string) (int64, error) {
	return ConvertFromStringToInteger(sizeStr, "M")
}

// ConvertToGigabytes converts a size string to gigabytes.
func ConvertToGigabytes(sizeStr string) (int64, error) {
	return ConvertFromStringToInteger(sizeStr, "G")
}

// ConvertToTerabytes converts a size string to terabytes.
func ConvertToTerabytes(sizeStr string) (int64, error) {
	return ConvertFromStringToInteger(sizeStr, "T")
}
