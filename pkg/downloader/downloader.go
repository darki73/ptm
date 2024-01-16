package downloader

import (
	"fmt"
	config "github.com/darki73/ptm/pkg/configuration"
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
	"github.com/darki73/ptm/pkg/distributions"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path"
)

// Downloader is the structure that holds the downloader configuration.
type Downloader struct {
	// saveTo is the path to save the downloaded file to.
	saveTo string
	// image is the reference to the image configuration.
	image *bi.Configuration
	// distributions is the reference to the distributions configuration.
	distributions *distributions.Distributions
	// distribution is the reference to the distribution configuration.
	distribution distributions.Distribution
}

// NewDownloader creates a new Downloader instance.
func NewDownloader(configuration *config.Configuration) (*Downloader, error) {
	distros, err := distributions.NewDistributions(configuration.GetBaseImage())
	if err != nil {
		return nil, err
	}
	return &Downloader{
		saveTo:        configuration.GetDownloader().GetSaveTo(),
		image:         configuration.GetBaseImage(),
		distributions: distros,
		distribution:  distros.GetActiveDistribution(),
	}, nil
}

// Download downloads the image.
func (downloader *Downloader) Download() error {
	alreadyDownloaded, err := downloader.IsAlreadyDownloaded()
	if err != nil {
		return err
	}

	if alreadyDownloaded {
		return nil
	}

	url, err := downloader.distribution.GetUrl()
	if err != nil {
		return err
	}

	savePath, err := downloader.GetFullImagePath()
	if err != nil {
		return err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()

	if err != nil {
		return err
	}

	handle, err := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Downloading %s\n", url)
	bar := progressbar.DefaultBytes(
		response.ContentLength,
	)

	if _, err = io.Copy(io.MultiWriter(handle, bar), response.Body); err != nil {
		return err
	}

	return nil
}

// IsAlreadyDownloaded returns true if the image is already downloaded.
func (downloader *Downloader) IsAlreadyDownloaded() (bool, error) {
	imageFullPath, err := downloader.GetFullImagePath()
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(imageFullPath); os.IsNotExist(err) {
		return false, nil
	}

	fmt.Printf("Image already downloaded at %s\n", imageFullPath)
	return true, nil
}

// GetFullImagePath returns the full image path.
func (downloader *Downloader) GetFullImagePath() (string, error) {
	imageName, err := downloader.distribution.GetImageName()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(downloader.saveTo); os.IsNotExist(err) {
		if err := os.MkdirAll(downloader.saveTo, 0755); err != nil {
			return "", err
		}
	}

	return path.Join(
		downloader.saveTo,
		imageName,
	), nil
}
