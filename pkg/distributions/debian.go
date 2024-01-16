package distributions

import (
	"fmt"
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
)

// Debian is the structure that holds configuration for Debian distributions.
type Debian struct {
	// baseImage is the user configuration for base image.
	baseImage *bi.Configuration
	// completeVersionBaseUrl is the base URL for complete version of Debian distributions.
	completeVersionBaseUrl string
	// minimalVersionBaseUrl is the base URL for minimal version of Debian distributions.
	minimalVersionBaseUrl string
	// versionToRelease is a map of Debian versions to releases.
	versionToRelease map[string]string
	// releaseToVersion is a map of Debian releases to versions.
	releaseToVersion map[string]string
	// supportedVersions is a list of supported versions of Debian.
	supportedVersions []string
	// supportedReleases is a list of supported releases of Debian.
	supportedReleases []string
	// completeSupportedArchitectures is a list of supported architectures for complete type of Debian.
	completeSupportedArchitectures []string
	// minimalSupportedArchitectures is a list of supported architectures for minimal type of Debian.
	minimalSupportedArchitectures []string
	// completeSupportedImageFormats is a list of supported image formats for complete type of Debian.
	completeSupportedImageFormats []string
	// minimalSupportedImageFormats is a list of supported image formats for minimal type of Debian.
	minimalSupportedImageFormats []string
}

// NewDebian returns a new instance of Debian distribution configuration.
func NewDebian() *Debian {
	return &Debian{
		baseImage:              nil,
		completeVersionBaseUrl: "https://cloud.debian.org/images/cloud",
		minimalVersionBaseUrl:  "https://cloud.debian.org/images/cloud",
		versionToRelease:       map[string]string{},
		releaseToVersion:       map[string]string{},
		supportedVersions:      []string{},
		supportedReleases:      []string{},
		completeSupportedArchitectures: []string{
			"amd64",
			"arm64",
		},
		completeSupportedImageFormats: []string{
			"qcow2",
			"raw",
		},
		minimalSupportedArchitectures: []string{
			"amd64",
		},
		minimalSupportedImageFormats: []string{
			"qcow2",
			"raw",
		},
	}
}

// Initialize initializes the Debian distribution.
func (debian *Debian) Initialize(baseImage *bi.Configuration) Distribution {
	debian.baseImage = baseImage

	debian.releaseToVersion = map[string]string{
		"bookworm": "12",
		"bullseye": "11",
		"buster":   "10",
		"stretch":  "9",
	}

	for key, value := range debian.releaseToVersion {
		debian.supportedReleases = append(debian.supportedReleases, key)
		debian.supportedVersions = append(debian.supportedVersions, value)
		debian.versionToRelease[value] = key
	}

	return debian
}

// GetVersionToRelease returns a map of Debian versions to releases.
func (debian *Debian) GetVersionToRelease() map[string]string {
	return debian.versionToRelease
}

// GetReleaseToVersion returns a map of Debian releases to versions.
func (debian *Debian) GetReleaseToVersion() map[string]string {
	return debian.releaseToVersion
}

// GetCompleteVersionBaseUrl returns the base URL to the complete version of Debian distributions.
func (debian *Debian) GetCompleteVersionBaseUrl() string {
	return debian.completeVersionBaseUrl
}

// GetMinimalVersionBaseUrl returns the base URL to the minimal version of Debian distributions.
func (debian *Debian) GetMinimalVersionBaseUrl() string {
	return debian.minimalVersionBaseUrl
}

// GetSupportedVersions returns a list of supported versions of Debian.
func (debian *Debian) GetSupportedVersions() []string {
	return debian.supportedVersions
}

// GetSupportedReleases returns a list of supported releases of Debian.
func (debian *Debian) GetSupportedReleases() []string {
	return debian.supportedReleases
}

// IsVersionSupported returns true if the version is supported by the distribution.
func (debian *Debian) IsVersionSupported(version string) bool {
	for _, supportedVersion := range debian.supportedVersions {
		if supportedVersion == version {
			return true
		}
	}
	return false
}

// IsReleaseSupported returns true if the release is supported by the distribution.
func (debian *Debian) IsReleaseSupported(release string) bool {
	for _, supportedRelease := range debian.supportedReleases {
		if supportedRelease == release {
			return true
		}
	}
	return false
}

// GetCompleteSupportedArchitectures returns a list of supported architectures for complete type of Debian.
func (debian *Debian) GetCompleteSupportedArchitectures() []string {
	return debian.completeSupportedArchitectures
}

// GetCompleteSupportedImageFormats returns a list of supported image formats for complete type of Debian.
func (debian *Debian) GetCompleteSupportedImageFormats() []string {
	return debian.completeSupportedImageFormats
}

// GetMinimalSupportedArchitectures returns a list of supported architectures for minimal type of Debian.
func (debian *Debian) GetMinimalSupportedArchitectures() []string {
	return debian.minimalSupportedArchitectures
}

// GetMinimalSupportedImageFormats returns a list of supported image formats for minimal type of Debian.
func (debian *Debian) GetMinimalSupportedImageFormats() []string {
	return debian.minimalSupportedImageFormats
}

// IsArchitectureSupported returns true if the architecture is supported by the distribution.
func (debian *Debian) IsArchitectureSupported(architecture string) bool {
	architectureRange := debian.completeSupportedArchitectures
	if debian.baseImage.GetMinimal() {
		architectureRange = debian.minimalSupportedArchitectures
	}

	for _, supportedArchitecture := range architectureRange {
		if supportedArchitecture == architecture {
			return true
		}
	}
	return false
}

// IsImageFormatSupported returns true if the image format is supported by the distribution.
func (debian *Debian) IsImageFormatSupported(imageFormat string) bool {
	imageFormatRange := debian.completeSupportedImageFormats
	if debian.baseImage.GetMinimal() {
		imageFormatRange = debian.minimalSupportedImageFormats
	}

	for _, supportedImageFormat := range imageFormatRange {
		if supportedImageFormat == imageFormat {
			return true
		}
	}
	return false
}

// GetVersionFromRelease returns the version of the Debian from the release.
func (debian *Debian) GetVersionFromRelease(release string) (string, error) {
	if debian.IsReleaseSupported(release) {
		return debian.releaseToVersion[release], nil
	}

	return "", fmt.Errorf("release %s is not supported", release)
}

// GetReleaseFromVersion returns the release of the Debian from the version.
func (debian *Debian) GetReleaseFromVersion(version string) (string, error) {
	if debian.IsVersionSupported(version) {
		return debian.versionToRelease[version], nil
	}

	return "", fmt.Errorf("version %s is not supported", version)
}

// GetReleaseFromReleaseOrVersion returns the release of the Debian from the release or version.
func (debian *Debian) GetReleaseFromReleaseOrVersion(releaseOrVersion string) (string, error) {
	if debian.IsReleaseSupported(releaseOrVersion) {
		return releaseOrVersion, nil
	}

	if debian.IsVersionSupported(releaseOrVersion) {
		return debian.versionToRelease[releaseOrVersion], nil
	}

	return "", fmt.Errorf("release or version %s is not supported", releaseOrVersion)
}

// GetVersionFromReleaseOrVersion returns the version of the Debian from the release or version.
func (debian *Debian) GetVersionFromReleaseOrVersion(releaseOrVersion string) (string, error) {
	if debian.IsReleaseSupported(releaseOrVersion) {
		return debian.releaseToVersion[releaseOrVersion], nil
	}

	if debian.IsVersionSupported(releaseOrVersion) {
		return releaseOrVersion, nil
	}

	return "", fmt.Errorf("release or version %s is not supported", releaseOrVersion)
}

// GetImageName returns the image name.
func (debian *Debian) GetImageName() (string, error) {
	version, err := debian.GetVersionFromReleaseOrVersion(debian.baseImage.GetRelease())
	if err != nil {
		return "", err
	}

	if !debian.IsArchitectureSupported(debian.baseImage.GetArchitecture()) {
		return "", fmt.Errorf("architecture %s is not supported", debian.baseImage.GetArchitecture())
	}

	if !debian.IsImageFormatSupported(debian.baseImage.GetFormat()) {
		return "", fmt.Errorf("image format %s is not supported", debian.baseImage.GetFormat())
	}

	if debian.baseImage.GetMinimal() {
		return fmt.Sprintf(
			"debian-%s-genericcloud-%s.%s",
			version,
			debian.baseImage.GetArchitecture(),
			debian.baseImage.GetFormat(),
		), nil
	}

	return fmt.Sprintf(
		"debian-%s-generic-%s.%s",
		version,
		debian.baseImage.GetArchitecture(),
		debian.baseImage.GetFormat(),
	), nil
}

// GetCompleteVersionUrl returns the complete version URL of the Debian.
func (debian *Debian) GetCompleteVersionUrl() (string, error) {
	release, err := debian.GetReleaseFromReleaseOrVersion(debian.baseImage.GetRelease())
	if err != nil {
		return "", err
	}

	imageName, err := debian.GetImageName()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s/%s/latest/%s",
		debian.GetCompleteVersionBaseUrl(),
		release,
		imageName,
	), nil
}

// GetMinimalVersionUrl returns the minimal version URL of the Debian.
func (debian *Debian) GetMinimalVersionUrl() (string, error) {
	release, err := debian.GetReleaseFromReleaseOrVersion(debian.baseImage.GetRelease())
	if err != nil {
		return "", err
	}

	imageName, err := debian.GetImageName()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s/%s/latest/%s",
		debian.GetMinimalVersionBaseUrl(),
		release,
		imageName,
	), nil
}

// GetUrl returns the URL of the Debian.
func (debian *Debian) GetUrl() (string, error) {
	if debian.baseImage.GetMinimal() {
		return debian.GetMinimalVersionUrl()
	}

	return debian.GetCompleteVersionUrl()
}
