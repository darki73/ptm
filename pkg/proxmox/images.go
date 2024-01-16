package proxmox

import (
	"os"
	"path"
)

// Images represents the structure containing the list of available images.
type Images struct {
	// path is the path where downloaded images are stored.
	path string
	// isos is the list of available images.
	isos []*Image
}

// NewImages creates a new Images instance.
func NewImages(saveTo string) (*Images, error) {
	images := &Images{
		path: saveTo,
		isos: make([]*Image, 0),
	}

	if err := images.listAvailableISOs(); err != nil {
		return nil, err
	}

	return images, nil
}

// GetPath returns the path where downloaded images are stored.
func (images *Images) GetPath() string {
	return images.path
}

// GetISOs returns the list of available images.
func (images *Images) GetISOs() []*Image {
	return images.isos
}

// FindISOByFullPath finds an image by its full path.
func (images *Images) FindISOByFullPath(fullPath string) *Image {
	for _, image := range images.isos {
		if image.GetFullPath() == fullPath {
			return image
		}
	}
	return nil
}

// listAvailableISOs lists the available images.
func (images *Images) listAvailableISOs() error {
	items, err := os.ReadDir(images.GetPath())
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}

		fileInfo, err := os.Stat(path.Join(images.GetPath(), item.Name()))
		if err != nil {
			return err
		}

		image, err := NewImage(item.Name(), images.GetPath(), fileInfo.Size())
		if err != nil {
			return err
		}

		images.isos = append(images.isos, image)
	}

	return nil
}
