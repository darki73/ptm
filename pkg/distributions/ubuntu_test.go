package distributions

import (
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
	"reflect"
	"sort"
	"testing"
)

// ubuntuTestInitialBaseImage is the initial base image configuration for Ubuntu for testing.
var ubuntuTestInitialBaseImage = &bi.Configuration{
	Distribution: "ubuntu",
	Release:      "jammy",
	Minimal:      true,
	Architecture: "amd64",
	Format:       "img",
}

// TestNewUbuntu tests the NewUbuntu function (and its default values).
func TestNewUbuntu(t *testing.T) {
	ubuntu := NewUbuntu()

	if ubuntu == nil {
		t.Error("NewUbuntu() returned nil")
	}

	if ubuntu.GetCompleteVersionBaseUrl() != "https://cloud-images.ubuntu.com/releases" {
		t.Errorf("Expected completeVersionBaseUrl to be 'https://cloud-images.ubuntu.com/releases', got %s", ubuntu.GetCompleteVersionBaseUrl())
	}

	if ubuntu.GetMinimalVersionBaseUrl() != "https://cloud-images.ubuntu.com/minimal/releases" {
		t.Errorf("Expected minimalVersionBaseUrl to be 'https://cloud-images.ubuntu.com/minimal/releases', got %s", ubuntu.GetMinimalVersionBaseUrl())
	}

	result := ubuntu.GetCompleteSupportedArchitectures()
	expected := []string{"amd64", "arm64", "armhf"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedArchitectures returned %v, want %v", result, expected)
	}

	result = ubuntu.GetCompleteSupportedImageFormats()
	expected = []string{"img", "vmdk"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedImageFormats returned %v, want %v", result, expected)
	}

	result = ubuntu.GetMinimalSupportedArchitectures()
	expected = []string{"amd64"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedArchitectures returned %v, want %v", result, expected)
	}

	result = ubuntu.GetMinimalSupportedImageFormats()
	expected = []string{"img"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedImageFormats returned %v, want %v", result, expected)
	}
}

// TestUbuntuInitialize tests the Initialize function.
func TestUbuntuInitialize(t *testing.T) {
	ubuntu := NewUbuntu()
	initializationResult := ubuntu.Initialize(ubuntuTestInitialBaseImage)

	if initializationResult != ubuntu {
		t.Error("Initialize() did not return the Ubuntu instance")
	}

	if ubuntu.baseImage != ubuntuTestInitialBaseImage {
		t.Error("Initialize() did not set the baseImage")
	}

	resultMapStringString := ubuntu.GetReleaseToVersion()
	expectedMapStringString := map[string]string{
		"mantic":  "23.10",
		"lunar":   "23.04",
		"kinetic": "22.10",
		"jammy":   "22.04",
		"impish":  "21.10",
		"hirsute": "21.04",
		"groovy":  "20.10",
		"focal":   "20.04",
		"bionic":  "18.04",
		"xenial":  "16.04",
		"trusty":  "14.04",
	}

	if !reflect.DeepEqual(resultMapStringString, expectedMapStringString) {
		t.Errorf("GetReleaseToVersion returned %v, want %v", resultMapStringString, expectedMapStringString)
	}

	resultMapStringString = ubuntu.GetVersionToRelease()
	expectedMapStringString = map[string]string{
		"23.10": "mantic",
		"23.04": "lunar",
		"22.10": "kinetic",
		"22.04": "jammy",
		"21.10": "impish",
		"21.04": "hirsute",
		"20.10": "groovy",
		"20.04": "focal",
		"18.04": "bionic",
		"16.04": "xenial",
		"14.04": "trusty",
	}

	if !reflect.DeepEqual(resultMapStringString, expectedMapStringString) {
		t.Errorf("GetVersionToRelease returned %v, want %v", resultMapStringString, expectedMapStringString)
	}

	resultSliceString := ubuntu.GetSupportedVersions()
	sort.Strings(resultSliceString)
	expectedSliceString := []string{
		"23.10",
		"23.04",
		"22.10",
		"22.04",
		"21.10",
		"21.04",
		"20.10",
		"20.04",
		"18.04",
		"16.04",
		"14.04",
	}
	sort.Strings(expectedSliceString)

	if !reflect.DeepEqual(resultSliceString, expectedSliceString) {
		t.Errorf("GetSupportedVersions returned %v, want %v", resultSliceString, expectedSliceString)
	}

	resultSliceString = ubuntu.GetSupportedReleases()
	sort.Strings(resultSliceString)
	expectedSliceString = []string{
		"mantic",
		"lunar",
		"kinetic",
		"jammy",
		"impish",
		"hirsute",
		"groovy",
		"focal",
		"bionic",
		"xenial",
		"trusty",
	}
	sort.Strings(expectedSliceString)

	if !reflect.DeepEqual(resultSliceString, expectedSliceString) {
		t.Errorf("GetSupportedReleases returned %v, want %v", resultSliceString, expectedSliceString)
	}
}

// TestUbuntuGetCompleteVersionBaseUrl tests the GetCompleteVersionBaseUrl function.
func TestUbuntuGetCompleteVersionBaseUrl(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetCompleteVersionBaseUrl()
	expected := "https://cloud-images.ubuntu.com/releases"

	if result != expected {
		t.Errorf("GetCompleteVersionBaseUrl returned %s, want %s", result, expected)
	}
}

// TestUbuntuGetMinimalVersionBaseUrl tests the GetMinimalVersionBaseUrl function.
func TestUbuntuGetMinimalVersionBaseUrl(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetMinimalVersionBaseUrl()
	expected := "https://cloud-images.ubuntu.com/minimal/releases"

	if result != expected {
		t.Errorf("GetMinimalVersionBaseUrl returned %s, want %s", result, expected)
	}
}

// TestUbuntuGetCompleteSupportedArchitectures tests the GetCompleteSupportedArchitectures function.
func TestUbuntuGetCompleteSupportedArchitectures(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetCompleteSupportedArchitectures()
	expected := []string{"amd64", "arm64", "armhf"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedArchitectures returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetCompleteSupportedImageFormats tests the GetCompleteSupportedImageFormats function.
func TestUbuntuGetCompleteSupportedImageFormats(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetCompleteSupportedImageFormats()
	expected := []string{"img", "vmdk"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetCompleteSupportedImageFormats returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetMinimalSupportedArchitectures tests the GetMinimalSupportedArchitectures function.
func TestUbuntuGetMinimalSupportedArchitectures(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetMinimalSupportedArchitectures()
	expected := []string{"amd64"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedArchitectures returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetMinimalSupportedImageFormats tests the GetMinimalSupportedImageFormats function.
func TestUbuntuGetMinimalSupportedImageFormats(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetMinimalSupportedImageFormats()
	expected := []string{"img"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetMinimalSupportedImageFormats returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetVersionToRelease tests the GetVersionToRelease function.
func TestUbuntuGetVersionToRelease(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetVersionToRelease()
	expected := map[string]string{
		"23.10": "mantic",
		"23.04": "lunar",
		"22.10": "kinetic",
		"22.04": "jammy",
		"21.10": "impish",
		"21.04": "hirsute",
		"20.10": "groovy",
		"20.04": "focal",
		"18.04": "bionic",
		"16.04": "xenial",
		"14.04": "trusty",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetVersionToRelease returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetReleaseToVersion tests the GetReleaseToVersion function.
func TestUbuntuGetReleaseToVersion(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetReleaseToVersion()
	expected := map[string]string{
		"mantic":  "23.10",
		"lunar":   "23.04",
		"kinetic": "22.10",
		"jammy":   "22.04",
		"impish":  "21.10",
		"hirsute": "21.04",
		"groovy":  "20.10",
		"focal":   "20.04",
		"bionic":  "18.04",
		"xenial":  "16.04",
		"trusty":  "14.04",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetReleaseToVersion returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetSupportedVersions tests the GetSupportedVersions function.
func TestUbuntuGetSupportedVersions(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetSupportedVersions()
	sort.Strings(result)
	expected := []string{
		"23.10",
		"23.04",
		"22.10",
		"22.04",
		"21.10",
		"21.04",
		"20.10",
		"20.04",
		"18.04",
		"16.04",
		"14.04",
	}
	sort.Strings(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetSupportedVersions returned %v, want %v", result, expected)
	}
}

// TestUbuntuGetSupportedReleases tests the GetSupportedReleases function.
func TestUbuntuGetSupportedReleases(t *testing.T) {
	ubuntu := NewUbuntu()
	ubuntu.Initialize(ubuntuTestInitialBaseImage)

	result := ubuntu.GetSupportedReleases()
	sort.Strings(result)
	expected := []string{
		"mantic",
		"lunar",
		"kinetic",
		"jammy",
		"impish",
		"hirsute",
		"groovy",
		"focal",
		"bionic",
		"xenial",
		"trusty",
	}
	sort.Strings(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetSupportedReleases returned %v, want %v", result, expected)
	}
}

// TestUbuntuIsArchitectureSupported tests the IsArchitectureSupported function.
func TestUbuntuIsArchitectureSupported(t *testing.T) {
	minimalBaseImage := &bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "img",
	}

	completeBaseImage := &bi.Configuration{
		Distribution: "ubuntu",
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
		{"MinimalARMHF", "armhf", false, minimalBaseImage},
		{"CompleteAMD64", "amd64", true, completeBaseImage},
		{"CompleteARM64", "arm64", true, completeBaseImage},
		{"CompleteARMHF", "armhf", true, completeBaseImage},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ubuntu := NewUbuntu().Initialize(test.baseImage)
			result := ubuntu.IsArchitectureSupported(test.arch)

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

// TestUbuntuIsImageFormatSupported tests the IsImageFormatSupported function.
func TestUbuntuIsImageFormatSupported(t *testing.T) {
	minimalBaseImage := &bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "img",
	}

	completeBaseImage := &bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "img",
	}

	tests := []struct {
		name      string
		format    string
		expected  bool
		baseImage *bi.Configuration
	}{
		{"MinimalIMG", "img", true, minimalBaseImage},
		{"MinimalVMDK", "vmdk", false, minimalBaseImage},
		{"CompleteIMG", "img", true, completeBaseImage},
		{"CompleteVMDK", "vmdk", true, completeBaseImage},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ubuntu := NewUbuntu().Initialize(test.baseImage)
			result := ubuntu.IsImageFormatSupported(test.format)

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

// TestUbuntuIsReleaseSupported tests the IsReleaseSupported function.
func TestUbuntuIsReleaseSupported(t *testing.T) {
	ubuntu := NewUbuntu().Initialize(ubuntuTestInitialBaseImage)

	tests := []struct {
		name     string
		release  string
		expected bool
	}{
		{"Supported", "jammy", true},
		{"Unsupported", "unsupported", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ubuntu.IsReleaseSupported(test.release)

			if result != test.expected {
				t.Errorf("IsReleaseSupported(%s) returned %v, want %v", test.release, result, test.expected)
			}
		})
	}
}

// TestUbuntuIsVersionSupported tests the IsVersionSupported function.
func TestUbuntuIsVersionSupported(t *testing.T) {
	ubuntu := NewUbuntu().Initialize(ubuntuTestInitialBaseImage)

	tests := []struct {
		name     string
		version  string
		expected bool
	}{
		{"Supported", "22.04", true},
		{"Unsupported", "unsupported", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ubuntu.IsVersionSupported(test.version)

			if result != test.expected {
				t.Errorf("IsVersionSupported(%s) returned %v, want %v", test.version, result, test.expected)
			}
		})
	}
}

// TestUbuntuGetCompleteVersionUrl tests the GetCompleteVersionUrl function.
func TestUbuntuGetCompleteVersionUrl(t *testing.T) {
	ubuntu := NewUbuntu().Initialize(&bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "img",
	})

	result, err := ubuntu.GetCompleteVersionUrl()
	expected := "https://cloud-images.ubuntu.com/releases/jammy/release/ubuntu-22.04-cloudimg-amd64.img"

	if err != nil {
		t.Errorf("GetCompleteVersionUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetCompleteVersionUrl returned %s, want %s", result, expected)
	}
}

// TestUbuntuGetMinimalVersionUrl tests the GetMinimalVersionUrl function.
func TestUbuntuGetMinimalVersionUrl(t *testing.T) {
	ubuntu := NewUbuntu().Initialize(&bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "img",
	})

	result, err := ubuntu.GetMinimalVersionUrl()
	expected := "https://cloud-images.ubuntu.com/minimal/releases/jammy/release/ubuntu-22.04-minimal-cloudimg-amd64.img"

	if err != nil {
		t.Errorf("GetMinimalVersionUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetMinimalVersionUrl returned %s, want %s", result, expected)
	}
}

// TestUbuntuGetUrlForComplete tests the GetUrl function.
func TestUbuntuGetUrlForComplete(t *testing.T) {
	ubuntu := NewUbuntu().Initialize(&bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      false,
		Architecture: "amd64",
		Format:       "img",
	})

	result, err := ubuntu.GetUrl()
	expected := "https://cloud-images.ubuntu.com/releases/jammy/release/ubuntu-22.04-cloudimg-amd64.img"

	if err != nil {
		t.Errorf("GetUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetUrl returned %s, want %s", result, expected)
	}
}

// TestUbuntuGetUrlForMinimal tests the GetUrl function.
func TestUbuntuGetUrlForMinimal(t *testing.T) {
	ubuntu := NewUbuntu().Initialize(&bi.Configuration{
		Distribution: "ubuntu",
		Release:      "jammy",
		Minimal:      true,
		Architecture: "amd64",
		Format:       "img",
	})

	result, err := ubuntu.GetUrl()
	expected := "https://cloud-images.ubuntu.com/minimal/releases/jammy/release/ubuntu-22.04-minimal-cloudimg-amd64.img"

	if err != nil {
		t.Errorf("GetUrl returned an error: %s", err)
	}

	if result != expected {
		t.Errorf("GetUrl returned %s, want %s", result, expected)
	}
}
