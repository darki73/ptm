package utils

// RemoveEmptyStringsFromSlice removes empty strings from a slice.
func RemoveEmptyStringsFromSlice(slice []string) []string {
	var newSlice []string

	for _, s := range slice {
		if s != "" {
			newSlice = append(newSlice, s)
		}
	}

	return newSlice
}
