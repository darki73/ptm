package distributions

import (
	"fmt"
	bi "github.com/darki73/ptm/pkg/configuration/base-image"
)

// Ubuntu is the structure that holds configuration for Ubuntu distributions.
type Ubuntu struct {
	// baseImage is the user configuration for base image.
	baseImage *bi.Configuration
	// completeVersionBaseUrl is the base URL for complete version of Ubuntu distributions.
	completeVersionBaseUrl string
	// minimalVersionBaseUrl is the base URL for minimal version of Ubuntu distributions.
	minimalVersionBaseUrl string
	// versionToRelease is a map of Ubuntu versions to releases.
	versionToRelease map[string]string
	// releaseToVersion is a map of Ubuntu releases to versions.
	releaseToVersion map[string]string
	// supportedVersions is a list of supported versions of Debian.
	supportedVersions []string
	// supportedReleases is a list of supported releases of Debian.
	supportedReleases []string
	// completeSupportedArchitectures is a list of supported architectures for complete type of Ubuntu.
	completeSupportedArchitectures []string
	// minimalSupportedArchitectures is a list of supported architectures for minimal type of Ubuntu.
	minimalSupportedArchitectures []string
	// completeSupportedImageFormats is a list of supported image formats for complete type of Ubuntu.
	completeSupportedImageFormats []string
	// minimalSupportedImageFormats is a list of supported image formats for minimal type of Ubuntu.
	minimalSupportedImageFormats []string
}

// NewUbuntu returns a new instance of Ubuntu distribution configuration.
func NewUbuntu() *Ubuntu {
	return &Ubuntu{
		baseImage:              nil,
		completeVersionBaseUrl: "https://cloud-images.ubuntu.com/releases",
		minimalVersionBaseUrl:  "https://cloud-images.ubuntu.com/minimal/releases",
		versionToRelease:       map[string]string{},
		releaseToVersion:       map[string]string{},
		supportedVersions:      []string{},
		supportedReleases:      []string{},
		completeSupportedArchitectures: []string{
			"amd64",
			"arm64",
			"armhf",
		},
		completeSupportedImageFormats: []string{
			"img",
			"vmdk",
		},
		minimalSupportedArchitectures: []string{
			"amd64",
		},
		minimalSupportedImageFormats: []string{
			"img",
		},
	}
}

// Initialize initializes the Ubuntu distribution.
func (ubuntu *Ubuntu) Initialize(baseImage *bi.Configuration) Distribution {
	ubuntu.baseImage = baseImage

	ubuntu.releaseToVersion = map[string]string{
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

	for key, value := range ubuntu.releaseToVersion {
		ubuntu.supportedReleases = append(ubuntu.supportedReleases, key)
		ubuntu.supportedVersions = append(ubuntu.supportedVersions, value)
		ubuntu.versionToRelease[value] = key
	}

	return ubuntu
}

// GetCompleteVersionBaseUrl returns the base URL to the complete version of Ubuntu.
func (ubuntu *Ubuntu) GetCompleteVersionBaseUrl() string {
	return ubuntu.completeVersionBaseUrl
}

// GetMinimalVersionBaseUrl returns the base URL to the minimal version of Ubuntu.
func (ubuntu *Ubuntu) GetMinimalVersionBaseUrl() string {
	return ubuntu.minimalVersionBaseUrl
}

// GetVersionToRelease returns a map of Ubuntu versions to releases.
func (ubuntu *Ubuntu) GetVersionToRelease() map[string]string {
	return ubuntu.versionToRelease
}

// GetReleaseToVersion returns a map of Ubuntu releases to versions.
func (ubuntu *Ubuntu) GetReleaseToVersion() map[string]string {
	return ubuntu.releaseToVersion
}

// GetSupportedVersions returns a list of supported versions of Ubuntu.
func (ubuntu *Ubuntu) GetSupportedVersions() []string {
	return ubuntu.supportedVersions
}

// GetSupportedReleases returns a list of supported releases of Ubuntu.
func (ubuntu *Ubuntu) GetSupportedReleases() []string {
	return ubuntu.supportedReleases
}

// IsVersionSupported returns true if the version is supported by the Ubuntu.
func (ubuntu *Ubuntu) IsVersionSupported(version string) bool {
	for _, supportedVersion := range ubuntu.supportedVersions {
		if supportedVersion == version {
			return true
		}
	}
	return false
}

// IsReleaseSupported returns true if the release is supported by the Ubuntu.
func (ubuntu *Ubuntu) IsReleaseSupported(release string) bool {
	for _, supportedRelease := range ubuntu.supportedReleases {
		if supportedRelease == release {
			return true
		}
	}
	return false
}

// GetCompleteSupportedArchitectures returns a list of supported architectures for complete type of Ubuntu.
func (ubuntu *Ubuntu) GetCompleteSupportedArchitectures() []string {
	return ubuntu.completeSupportedArchitectures
}

// GetCompleteSupportedImageFormats returns a list of supported image formats for complete type of Ubuntu.
func (ubuntu *Ubuntu) GetCompleteSupportedImageFormats() []string {
	return ubuntu.completeSupportedImageFormats
}

// GetMinimalSupportedArchitectures returns a list of supported architectures for minimal type of Ubuntu.
func (ubuntu *Ubuntu) GetMinimalSupportedArchitectures() []string {
	return ubuntu.minimalSupportedArchitectures
}

// GetMinimalSupportedImageFormats returns a list of supported image formats for minimal type of Ubuntu.
func (ubuntu *Ubuntu) GetMinimalSupportedImageFormats() []string {
	return ubuntu.minimalSupportedImageFormats
}

// IsArchitectureSupported returns true if the architecture is supported by the Ubuntu.
func (ubuntu *Ubuntu) IsArchitectureSupported(architecture string) bool {
	architectureRange := ubuntu.completeSupportedArchitectures
	if ubuntu.baseImage.GetMinimal() {
		architectureRange = ubuntu.minimalSupportedArchitectures
	}

	for _, supportedArchitecture := range architectureRange {
		if supportedArchitecture == architecture {
			return true
		}
	}
	return false
}

// IsImageFormatSupported returns true if the image format is supported by the Ubuntu.
func (ubuntu *Ubuntu) IsImageFormatSupported(imageFormat string) bool {
	imageFormatRange := ubuntu.completeSupportedImageFormats
	if ubuntu.baseImage.GetMinimal() {
		imageFormatRange = ubuntu.minimalSupportedImageFormats
	}

	for _, supportedImageFormat := range imageFormatRange {
		if supportedImageFormat == imageFormat {
			return true
		}
	}
	return false
}

// GetVersionFromRelease returns the version of the Ubuntu from the release.
func (ubuntu *Ubuntu) GetVersionFromRelease(release string) (string, error) {
	if ubuntu.IsReleaseSupported(release) {
		return ubuntu.releaseToVersion[release], nil
	}

	return "", fmt.Errorf("release %s is not supported", release)
}

// GetReleaseFromVersion returns the release of the Ubuntu from the version.
func (ubuntu *Ubuntu) GetReleaseFromVersion(version string) (string, error) {
	if ubuntu.IsVersionSupported(version) {
		return ubuntu.versionToRelease[version], nil
	}

	return "", fmt.Errorf("version %s is not supported", version)
}

// GetReleaseFromReleaseOrVersion returns the release of the Ubuntu from the release or version.
func (ubuntu *Ubuntu) GetReleaseFromReleaseOrVersion(releaseOrVersion string) (string, error) {
	if ubuntu.IsReleaseSupported(releaseOrVersion) {
		return releaseOrVersion, nil
	}

	if ubuntu.IsVersionSupported(releaseOrVersion) {
		return ubuntu.GetReleaseFromVersion(releaseOrVersion)
	}

	return "", fmt.Errorf("release or version %s is not supported", releaseOrVersion)
}

// GetVersionFromReleaseOrVersion returns the version of the Ubuntu from the release or version.
func (ubuntu *Ubuntu) GetVersionFromReleaseOrVersion(releaseOrVersion string) (string, error) {
	if ubuntu.IsReleaseSupported(releaseOrVersion) {
		return ubuntu.GetVersionFromRelease(releaseOrVersion)
	}

	if ubuntu.IsVersionSupported(releaseOrVersion) {
		return releaseOrVersion, nil
	}

	return "", fmt.Errorf("release or version %s is not supported", releaseOrVersion)
}

// GetImageName returns the image name.
func (ubuntu *Ubuntu) GetImageName() (string, error) {
	version, err := ubuntu.GetVersionFromReleaseOrVersion(ubuntu.baseImage.GetRelease())
	if err != nil {
		return "", err
	}

	if !ubuntu.IsArchitectureSupported(ubuntu.baseImage.GetArchitecture()) {
		return "", fmt.Errorf("architecture %s is not supported", ubuntu.baseImage.GetArchitecture())
	}

	if !ubuntu.IsImageFormatSupported(ubuntu.baseImage.GetFormat()) {
		return "", fmt.Errorf("image format %s is not supported", ubuntu.baseImage.GetFormat())
	}

	if ubuntu.baseImage.GetMinimal() {
		return fmt.Sprintf(
			"ubuntu-%s-minimal-cloudimg-%s.%s",
			version,
			ubuntu.baseImage.GetArchitecture(),
			ubuntu.baseImage.GetFormat(),
		), nil
	}

	return fmt.Sprintf(
		"ubuntu-%s-cloudimg-%s.%s",
		version,
		ubuntu.baseImage.GetArchitecture(),
		ubuntu.baseImage.GetFormat(),
	), nil
}

// GetCompleteVersionUrl returns the complete version URL of the Ubuntu.
func (ubuntu *Ubuntu) GetCompleteVersionUrl() (string, error) {
	release, err := ubuntu.GetReleaseFromReleaseOrVersion(ubuntu.baseImage.GetRelease())
	if err != nil {
		return "", err
	}

	imageName, err := ubuntu.GetImageName()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s/%s/release/%s",
		ubuntu.GetCompleteVersionBaseUrl(),
		release,
		imageName,
	), nil
}

// GetMinimalVersionUrl returns the minimal version URL of the Ubuntu.
func (ubuntu *Ubuntu) GetMinimalVersionUrl() (string, error) {
	release, err := ubuntu.GetReleaseFromReleaseOrVersion(ubuntu.baseImage.GetRelease())
	if err != nil {
		return "", err
	}

	imageName, err := ubuntu.GetImageName()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s/%s/release/%s",
		ubuntu.GetMinimalVersionBaseUrl(),
		release,
		imageName,
	), nil
}

// GetUrl returns the URL of the Ubuntu.
func (ubuntu *Ubuntu) GetUrl() (string, error) {
	if ubuntu.baseImage.GetMinimal() {
		return ubuntu.GetMinimalVersionUrl()
	}

	return ubuntu.GetCompleteVersionUrl()
}
