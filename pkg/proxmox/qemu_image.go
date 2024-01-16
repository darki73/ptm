package proxmox

// QemuImage represents a QEMU image structure.
type QemuImage struct {
	// VirtualSize is the virtual size of the image (in bytes).
	VirtualSize int64 `json:"virtual-size"`
	// Filename is the filename of the image.
	Filename string `json:"filename"`
	// ClusterSize is the cluster size of the image (in bytes).
	ClusterSize int64 `json:"cluster-size"`
	// Format is the format of the image.
	Format string `json:"format"`
	// ActualSize is the actual size of the image (in bytes).
	ActualSize int64 `json:"actual-size"`
	// DirtyFlag is the dirty flag of the image.
	DirtyFlag bool `json:"dirty-flag"`
}

// GetVirtualSize returns the virtual size of the image (in bytes).
func (qemuImage *QemuImage) GetVirtualSize() int64 {
	return qemuImage.VirtualSize
}

// GetVirtualSizeInKilobytes returns the virtual size of the image (in kilobytes).
func (qemuImage *QemuImage) GetVirtualSizeInKilobytes() int64 {
	return qemuImage.VirtualSize / 1024
}

// GetFilename returns the filename of the image.
func (qemuImage *QemuImage) GetFilename() string {
	return qemuImage.Filename
}

// GetClusterSize returns the cluster size of the image (in bytes).
func (qemuImage *QemuImage) GetClusterSize() int64 {
	return qemuImage.ClusterSize
}

// GetFormat returns the format of the image.
func (qemuImage *QemuImage) GetFormat() string {
	return qemuImage.Format
}

// GetActualSize returns the actual size of the image (in bytes).
func (qemuImage *QemuImage) GetActualSize() int64 {
	return qemuImage.ActualSize
}

// GetActualSizeInKilobytes returns the actual size of the image (in kilobytes).
func (qemuImage *QemuImage) GetActualSizeInKilobytes() int64 {
	return qemuImage.ActualSize / 1024
}

// IsDirty returns true if the image is dirty, false otherwise.
func (qemuImage *QemuImage) IsDirty() bool {
	return qemuImage.DirtyFlag
}
