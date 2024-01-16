package distributions

import bi "github.com/darki73/ptm/pkg/configuration/base-image"

// Distribution is the interface that defines the methods that must be implemented by an underlying distribution.
type Distribution interface {
	// Initialize initializes the distribution.
	Initialize(baseImage *bi.Configuration) Distribution
	// GetCompleteVersionBaseUrl returns the base URL to the complete version of the distribution.
	GetCompleteVersionBaseUrl() string
	// GetMinimalVersionBaseUrl returns the base URL to the minimal version of the distribution.
	GetMinimalVersionBaseUrl() string
	// GetVersionToRelease returns a map of distribution versions to releases.
	GetVersionToRelease() map[string]string
	// GetReleaseToVersion returns a map of distribution releases to versions.
	GetReleaseToVersion() map[string]string
	// GetSupportedVersions returns a list of supported versions of the distribution.
	GetSupportedVersions() []string
	// GetSupportedReleases returns a list of supported releases of the distribution.
	GetSupportedReleases() []string
	// IsVersionSupported returns true if the version is supported by the distribution.
	IsVersionSupported(version string) bool
	// IsReleaseSupported returns true if the release is supported by the distribution.
	IsReleaseSupported(release string) bool
	// GetCompleteSupportedArchitectures returns a list of supported architectures for complete type of the distribution.
	GetCompleteSupportedArchitectures() []string
	// GetCompleteSupportedImageFormats returns a list of supported image formats for complete type of the distribution.
	GetCompleteSupportedImageFormats() []string
	// GetMinimalSupportedArchitectures returns a list of supported architectures for minimal type of the distribution.
	GetMinimalSupportedArchitectures() []string
	// GetMinimalSupportedImageFormats returns a list of supported image formats for minimal type of the distribution.
	GetMinimalSupportedImageFormats() []string
	// IsArchitectureSupported returns true if the architecture is supported by the distribution.
	IsArchitectureSupported(architecture string) bool
	// IsImageFormatSupported returns true if the image format is supported by the distribution.
	IsImageFormatSupported(imageFormat string) bool
	// GetVersionFromRelease returns the version of the distribution from the release.
	GetVersionFromRelease(release string) (string, error)
	// GetReleaseFromVersion returns the release of the distribution from the version.
	GetReleaseFromVersion(version string) (string, error)
	// GetReleaseFromReleaseOrVersion returns the release of the distribution from the release or version.
	GetReleaseFromReleaseOrVersion(releaseOrVersion string) (string, error)
	// GetVersionFromReleaseOrVersion returns the version of the distribution from the release or version.
	GetVersionFromReleaseOrVersion(releaseOrVersion string) (string, error)
	// GetImageName returns the image name.
	GetImageName() (string, error)
	// GetCompleteVersionUrl returns the complete version URL of the distribution.
	GetCompleteVersionUrl() (string, error)
	// GetMinimalVersionUrl returns the minimal version URL of the distribution.
	GetMinimalVersionUrl() (string, error)
	// GetUrl returns the URL of the distribution.
	GetUrl() (string, error)
}
