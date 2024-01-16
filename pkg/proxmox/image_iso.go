package proxmox

import (
	"encoding/json"
	"github.com/darki73/ptm/pkg/utils"
	"path"
)

// Image represents the downloaded image.
type Image struct {
	// name is the name of the image.
	name string
	// path is the path of the image.
	path string
	// size is the size of the image (in kilobytes).
	size int64
	// virtualSize is the virtual size of the image (in kilobytes).
	virtualSize int64
	// qemuImage is the QemuImage struct.
	qemuImage *QemuImage
}

// NewImage creates a new Image instance.
func NewImage(name string, path string, size int64) (*Image, error) {
	image := &Image{
		name:        name,
		path:        path,
		size:        size / 1024,
		virtualSize: 0,
		qemuImage:   &QemuImage{},
	}

	if err := image.loadQemuInfo(); err != nil {
		return nil, err
	}

	return image, nil
}

// GetName returns the name of the image.
func (image *Image) GetName() string {
	return image.name
}

// GetPath returns the path of the image.
func (image *Image) GetPath() string {
	return image.path
}

// GetFullPath returns the full path of the image.
func (image *Image) GetFullPath() string {
	return path.Join(image.GetPath(), image.GetName())
}

// GetSize returns the size of the image (in kilobytes).
func (image *Image) GetSize() int64 {
	return image.size
}

// GetVirtualSize returns the virtual size of the image (in kilobytes).
func (image *Image) GetVirtualSize() int64 {
	return image.virtualSize
}

// GetQemuImageInformation returns the Qemu image information.
func (image *Image) GetQemuImageInformation() *QemuImage {
	return image.qemuImage
}

// loadQemuInfo loads the Qemu image information.
func (image *Image) loadQemuInfo() error {
	output, err := utils.ExecuteCommand("qemu-img", "info", "--output=json", path.Join(image.GetPath(), image.GetName()))
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(output), image.qemuImage); err != nil {
		return err
	}

	image.virtualSize = image.qemuImage.GetVirtualSizeInKilobytes()

	return nil
}
