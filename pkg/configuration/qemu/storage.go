package qemu

// QemuStorage is a structure that holds information for QEMU storage.
type QemuStorage struct {
	// Name is the name of the storage.
	Name string `json:"name" yaml:"name" toml:"name" mapstructure:"name"`
	// Resize is the value to which we should resize image.
	Resize string `json:"resize" yaml:"resize" toml:"resize" mapstructure:"resize"`
}

// InitializeQemuStorageWithDefaults initializes the storage with defaults.
func InitializeQemuStorageWithDefaults() *QemuStorage {
	return &QemuStorage{
		Name:   "",
		Resize: "",
	}
}

// GetStorage returns the storage to use.
func (qemuStorage *QemuStorage) GetStorage() string {
	return qemuStorage.Name
}

// GetResize returns the value to which we should resize image.
func (qemuStorage *QemuStorage) GetResize() string {
	return qemuStorage.Resize
}

// IsSupposedToResize returns true if the image is supposed to be resized.
func (qemuStorage *QemuStorage) IsSupposedToResize() bool {
	return qemuStorage.Resize != ""
}

// IsConfigured returns true if the storage is configured.
func (qemuStorage *QemuStorage) IsConfigured() bool {
	if qemuStorage.Name == "" {
		return false
	}

	return true
}
