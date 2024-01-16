package configuration

import (
	"testing"
)

// TestNewOptions tests the NewOptions function for correct default and custom value assignments.
func TestNewOptions(t *testing.T) {
	defaultOptions := NewDefaultOptions()

	// Test with custom values
	customName, customPath, customExtension := "custom-config", "/custom/path", "json"
	options := NewOptions(customName, customPath, customExtension)
	if options.Name != customName {
		t.Errorf("Expected Name to be %s, got %s", customName, options.Name)
	}
	if options.Path != customPath {
		t.Errorf("Expected Path to be %s, got %s", customPath, options.Path)
	}
	if options.Extension != customExtension {
		t.Errorf("Expected Extension to be %s, got %s", customExtension, options.Extension)
	}

	// Test with empty values (should use defaults)
	options = NewOptions("", "", "")
	if options.Name != defaultOptions.GetName() {
		t.Errorf("Expected default Name %s, got %s", defaultOptions.GetName(), options.Name)
	}
	if options.Path != defaultOptions.GetPath() {
		t.Errorf("Expected default Path %s, got %s", defaultOptions.GetPath(), options.Path)
	}
	if options.Extension != defaultOptions.GetExtension() {
		t.Errorf("Expected default Extension %s, got %s", defaultOptions.GetExtension(), options.Extension)
	}
}

// TestNewCustomOptions tests the creation of an Options struct with custom values.
func TestNewCustomOptions(t *testing.T) {
	name, path, extension := "myconfig", "/my/path", "toml"
	options := NewCustomOptions(name, path, extension)

	if options.GetName() != name {
		t.Errorf("GetName() = %s; want %s", options.GetName(), name)
	}
	if options.GetPath() != path {
		t.Errorf("GetPath() = %s; want %s", options.GetPath(), path)
	}
	if options.GetExtension() != extension {
		t.Errorf("GetExtension() = %s; want %s", options.GetExtension(), extension)
	}
}

// TestNewDefaultOptions tests the creation of an Options struct with default values.
func TestNewDefaultOptions(t *testing.T) {
	options := NewDefaultOptions()

	expectedName, expectedPath, expectedExtension := "config", "/etc/ptm", "yaml"
	if options.GetName() != expectedName {
		t.Errorf("Expected default Name %s, got %s", expectedName, options.GetName())
	}
	if options.GetPath() != expectedPath {
		t.Errorf("Expected default Path %s, got %s", expectedPath, options.GetPath())
	}
	if options.GetExtension() != expectedExtension {
		t.Errorf("Expected default Extension %s, got %s", expectedExtension, options.GetExtension())
	}
}

// TestOptionsGetters tests the getter methods of the Options struct.
func TestOptionsGetters(t *testing.T) {
	name, path, extension := "test-config", "/test/path", "ini"
	options := Options{
		Name:      name,
		Path:      path,
		Extension: extension,
	}

	if options.GetName() != name {
		t.Errorf("GetName() = %s; want %s", options.GetName(), name)
	}
	if options.GetPath() != path {
		t.Errorf("GetPath() = %s; want %s", options.GetPath(), path)
	}
	if options.GetExtension() != extension {
		t.Errorf("GetExtension() = %s; want %s", options.GetExtension(), extension)
	}
}
