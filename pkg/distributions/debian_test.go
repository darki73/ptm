package distributions

import (
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
	"reflect"
	"sort"
	"testing"
)

// debianTestInitialBaseImage is the initial base image configuration for Debian for testing.
var debianTestInitialBaseImage = &bi.Configuration{
	Distribution: "debian",
	Release:      "12",
	Minimal:      true,
	Architecture: "amd64",
	Format:       "qcow2",
}

// TestNewDebian tests the NewDebian function (and its default values).
func TestNewDebian(t *testing.T) {
	debian := NewDebian()

	if debian == nil {
		t.Error("NewDebian() returned nil")
	}

	if debian.GetCompleteVersionBaseUrl() != "https://cloud.debian.org/images/cloud" {
		t.Errorf("Expected completeVersionBaseUrl to be 'https://cloud.debian.org/images/cloud', got %s", debian.GetCompleteVersionBaseUrl())
	}

	if debian.GetMinimalVersionBaseUrl() != "https://cloud.debian.org/images/cloud" {
		t.Errorf("Expected minimalVersionBaseUrl to be 'https://cloud.debian.org/images/cloud', got %s", debian.GetMinimalVersionBaseUrl())
	}

	result := debian.GetCompleteSupportedArchitectures()
	expected := []string{"amd64", "arm64"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedArchitectures returned %v, want %v", result, expected)
	}

	result = debian.GetCompleteSupportedImageFormats()
	expected = []string{"qcow2", "raw"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedImageFormats returned %v, want %v", result, expected)
	}

	result = debian.GetMinimalSupportedArchitectures()
	expected = []string{"amd64"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedArchitectures returned %v, want %v", result, expected)
	}

	result = debian.GetMinimalSupportedImageFormats()
	expected = []string{"qcow2", "raw"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedImageFormats returned %v, want %v", result, expected)
	}
}

// TestDebianInitialize tests the Initialize function.
func TestDebianInitialize(t *testing.T) {
	debian := NewDebian()
	initializationResult := debian.Initialize(debianTestInitialBaseImage)

	if initializationResult != debian {
		t.Error("Initialize() did not return the Debian instance")
	}

	if debian.baseImage != debianTestInitialBaseImage {
		t.Error("Initialize() did not set the baseImage")
	}

	resultMapStringString := debian.GetReleaseToVersion()
	expectedMapStringString := map[string]string{
		"bookworm": "12",
		"bullseye": "11",
		"buster":   "10",
		"stretch":  "9",
	}

	if !reflect.DeepEqual(resultMapStringString, expectedMapStringString) {
		t.Errorf("GetReleaseToVersion returned %v, want %v", resultMapStringString, expectedMapStringString)
	}

	resultMapStringString = debian.GetVersionToRelease()
	expectedMapStringString = map[string]string{
		"12": "bookworm",
		"11": "bullseye",
		"10": "buster",
		"9":  "stretch",
	}

	if !reflect.DeepEqual(resultMapStringString, expectedMapStringString) {
		t.Errorf("GetVersionToRelease returned %v, want %v", resultMapStringString, expectedMapStringString)
	}

	resultSliceString := debian.GetSupportedVersions()
	sort.Strings(resultSliceString)
	expectedSliceString := []string{
		"12",
		"11",
		"10",
		"9",
	}
	sort.Strings(expectedSliceString)

	if !reflect.DeepEqual(resultSliceString, expectedSliceString) {
		t.Errorf("GetSupportedVersions returned %v, want %v", resultSliceString, expectedSliceString)
	}

	resultSliceString = debian.GetSupportedReleases()
	sort.Strings(resultSliceString)
	expectedSliceString = []string{
		"bookworm",
		"bullseye",
		"buster",
		"stretch",
	}
	sort.Strings(expectedSliceString)

	if !reflect.DeepEqual(resultSliceString, expectedSliceString) {
		t.Errorf("GetSupportedReleases returned %v, want %v", resultSliceString, expectedSliceString)
	}
}

// TestDebianGetCompleteVersionBaseUrl tests the GetCompleteVersionBaseUrl function.
func TestDebianGetCompleteVersionBaseUrl(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetCompleteVersionBaseUrl()
	expected := "https://cloud.debian.org/images/cloud"

	if result != expected {
		t.Errorf("GetCompleteVersionBaseUrl returned %s, want %s", result, expected)
	}
}

// TestDebianGetMinimalVersionBaseUrl tests the GetMinimalVersionBaseUrl function.
func TestDebianGetMinimalVersionBaseUrl(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetMinimalVersionBaseUrl()
	expected := "https://cloud.debian.org/images/cloud"

	if result != expected {
		t.Errorf("GetMinimalVersionBaseUrl returned %s, want %s", result, expected)
	}
}

// TestDebianGetCompleteSupportedArchitectures tests the GetCompleteSupportedArchitectures function.
func TestDebianGetCompleteSupportedArchitectures(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetCompleteSupportedArchitectures()
	expected := []string{"amd64", "arm64"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedArchitectures returned %v, want %v", result, expected)
	}
}

// TestDebianGetCompleteSupportedImageFormats tests the GetCompleteSupportedImageFormats function.
func TestDebianGetCompleteSupportedImageFormats(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetCompleteSupportedImageFormats()
	expected := []string{"qcow2", "raw"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedImageFormats returned %v, want %v", result, expected)
	}
}

// TestDebianGetMinimalSupportedArchitectures tests the GetMinimalSupportedArchitectures function.
func TestDebianGetMinimalSupportedArchitectures(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetMinimalSupportedArchitectures()
	expected := []string{"amd64"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedArchitectures returned %v, want %v", result, expected)
	}
}

// TestDebianGetMinimalSupportedImageFormats tests the GetMinimalSupportedImageFormats function.
func TestDebianGetMinimalSupportedImageFormats(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetMinimalSupportedImageFormats()
	expected := []string{"qcow2", "raw"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedImageFormats returned %v, want %v", result, expected)
	}
}

// TestDebianGetVersionToRelease tests the GetVersionToRelease function.
func TestDebianGetVersionToRelease(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetVersionToRelease()
	expected := map[string]string{
		"12": "bookworm",
		"11": "bullseye",
		"10": "buster",
		"9":  "stretch",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetVersionToRelease returned %v, want %v", result, expected)
	}
}

// TestDebianGetReleaseToVersion tests the GetReleaseToVersion function.
func TestDebianGetReleaseToVersion(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetReleaseToVersion()
	expected := map[string]string{
		"bookworm": "12",
		"bullseye": "11",
		"buster":   "10",
		"stretch":  "9",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetReleaseToVersion returned %v, want %v", result, expected)
	}
}

// TestDebianGetSupportedVersions tests the GetSupportedVersions function.
func TestDebianGetSupportedVersions(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetSupportedVersions()
	sort.Strings(result)
	expected := []string{
		"12",
		"11",
		"10",
		"9",
	}
	sort.Strings(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetSupportedVersions returned %v, want %v", result, expected)
	}
}

// TestDebianGetSupportedReleases tests the GetSupportedReleases function.
func TestDebianGetSupportedReleases(t *testing.T) {
	debian := NewDebian()
	debian.Initialize(debianTestInitialBaseImage)

	result := debian.GetSupportedReleases()
	sort.Strings(result)
	expected := []string{
		"bookworm",
		"bullseye",
		"buster",
		"stretch",
	}
	sort.Strings(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetSupportedReleases returned %v, want %v", result, expected)
	}
}

// TestDebianIsArchitectureSupported tests the IsArchitectureSupported function.
func TestDebianIsArchitectureSupported(t *testing.T) {
	minimalBaseImage := &bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "qcow2",
	}

	completeBaseImage := &bi.Configuration{
		Distribution: "debian",
		Release:      "jammy",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "img",
	}

	tests := []struct {
		name      string
		arch      string
		expected  bool
		baseImage *bi.Configuration
	}{
		{"MinimalAMD64", "amd64", true, minimalBaseImage},
		{"MinimalARM64", "arm64", false, minimalBaseImage},
		{"CompleteAMD64", "amd64", true, completeBaseImage},
		{"CompleteARM64", "arm64", true, completeBaseImage},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			debian := NewDebian().Initialize(test.baseImage)
			result := debian.IsArchitectureSupported(test.arch)

			imageType := "complete"
			if test.baseImage.Minimal {
				imageType = "minimal"
			}

			if result != test.expected {
				t.Errorf("IsArchitectureSupported(%s) returned %v, want %v", imageType, result, test.expected)
			}
		})
	}
}

// TestDebianIsImageFormatSupported tests the IsImageFormatSupported function.
func TestDebianIsImageFormatSupported(t *testing.T) {
	minimalBaseImage := &bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "qcow2",
	}

	completeBaseImage := &bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "qcow2",
	}

	tests := []struct {
		name      string
		format    string
		expected  bool
		baseImage *bi.Configuration
	}{
		{"MinimalQcow2", "qcow2", true, minimalBaseImage},
		{"MinimalRaw", "raw", true, minimalBaseImage},
		{"MinimalImg", "img", false, minimalBaseImage},
		{"CompleteQcow2", "qcow2", true, completeBaseImage},
		{"CompleteRaw", "raw", true, completeBaseImage},
		{"CompleteImg", "img", false, completeBaseImage},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			debian := NewDebian().Initialize(test.baseImage)
			result := debian.IsImageFormatSupported(test.format)

			imageType := "complete"
			if test.baseImage.Minimal {
				imageType = "minimal"
			}

			if result != test.expected {
				t.Errorf("IsImageFormatSupported(%s) returned %v, want %v", imageType, result, test.expected)
			}
		})
	}
}

// TestDebianIsReleaseSupported tests the IsReleaseSupported function.
func TestDebianIsReleaseSupported(t *testing.T) {
	debian := NewDebian().Initialize(debianTestInitialBaseImage)

	tests := []struct {
		name     string
		release  string
		expected bool
	}{
		{"Supported", "bookworm", true},
		{"Unsupported", "unsupported", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := debian.IsReleaseSupported(test.release)

			if result != test.expected {
				t.Errorf("IsReleaseSupported(%s) returned %v, want %v", test.release, result, test.expected)
			}
		})
	}
}

// TestDebianIsVersionSupported tests the IsVersionSupported function.
func TestDebianIsVersionSupported(t *testing.T) {
	debian := NewDebian().Initialize(debianTestInitialBaseImage)

	tests := []struct {
		name     string
		version  string
		expected bool
	}{
		{"Supported", "12", true},
		{"Unsupported", "unsupported", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := debian.IsVersionSupported(test.version)

			if result != test.expected {
				t.Errorf("IsVersionSupported(%s) returned %v, want %v", test.version, result, test.expected)
			}
		})
	}
}

// TestDebianGetCompleteVersionUrl tests the GetCompleteVersionUrl function.
func TestDebianGetCompleteVersionUrl(t *testing.T) {
	debian := NewDebian().Initialize(&bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "qcow2",
	})

	result, err := debian.GetCompleteVersionUrl()
	expected := "https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-generic-amd64.qcow2"

	if err != nil {
		t.Errorf("GetCompleteVersionUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetCompleteVersionUrl returned %s, want %s", result, expected)
	}
}

// TestDebianGetMinimalVersionUrl tests the GetMinimalVersionUrl function.
func TestDebianGetMinimalVersionUrl(t *testing.T) {
	debian := NewDebian().Initialize(&bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "qcow2",
	})

	result, err := debian.GetMinimalVersionUrl()
	expected := "https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-genericcloud-amd64.qcow2"

	if err != nil {
		t.Errorf("GetMinimalVersionUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetMinimalVersionUrl returned %s, want %s", result, expected)
	}
}

// TestDebianGetUrlForComplete tests the GetUrl function.
func TestDebianGetUrlForComplete(t *testing.T) {
	debian := NewDebian().Initialize(&bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "qcow2",
	})

	result, err := debian.GetUrl()
	expected := "https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-generic-amd64.qcow2"

	if err != nil {
		t.Errorf("GetUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetUrl returned %s, want %s", result, expected)
	}
}

// TestDebianGetUrlForMinimal tests the GetUrl function.
func TestDebianGetUrlForMinimal(t *testing.T) {
	debian := NewDebian().Initialize(&bi.Configuration{
		Distribution: "debian",
		Release:      "12",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "qcow2",
	})

	result, err := debian.GetUrl()
	expected := "https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-genericcloud-amd64.qcow2"

	if err != nil {
		t.Errorf("GetUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetUrl returned %s, want %s", result, expected)
	}
}
