package version

import "runtime"

var (
	// Version is the current version of the app
	Version = "dev"
	// Commit is the current git commit of the app
	Commit = "none"
	// Date is the date the binary was built
	Date = "0000-00-00T00:00:00Z"
	// Builder is the user that built the binary
	Builder = "unknown"
	// goVersion is the version of the go compiler used to build the binary
	goVersion = runtime.Version()
	// osArch is the os and arch the binary was built for
	osArch = runtime.GOARCH
	// osName is the os name the binary was built for
	osName = runtime.GOOS
	// versionTemplate is the template for the version command
	versionTemplate = `Version:      {{ .Version }}
SHA Commit:   {{ .Commit }}
Go version:   {{ .GoVersion }}
Built On:     {{ .Date }}
Built By:     {{ .Builder }}
OS/Arch:      {{ .OsName }}/{{ .OsArch }}
`
)

// BuildInfo holds the build information
type BuildInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
	Builder   string `json:"builder"`
	GoVersion string `json:"go_version"`
	OsArch    string `json:"os_arch"`
	OsName    string `json:"os_name"`
}

// GetVersion returns the version of the app
func GetVersion() string {
	return Version
}

// SetVersion sets the version of the app
func SetVersion(v string) {
	Version = v
}

// GetCommit returns the commit of the app
func GetCommit() string {
	return Commit
}

// SetCommit sets the commit of the app
func SetCommit(c string) {
	Commit = c
}

// GetDate returns the date the binary was built
func GetDate() string {
	return Date
}

// SetDate sets the date the binary was built
func SetDate(d string) {
	Date = d
}

// GetBuilder returns the user that built the binary
func GetBuilder() string {
	return Builder
}

// SetBuilder sets the user that built the binary
func SetBuilder(b string) {
	Builder = b
}

// GetGoVersion returns the version of the go compiler used to build the binary
func GetGoVersion() string {
	return goVersion
}

// GetOsArch returns the os and arch the binary was built for
func GetOsArch() string {
	return osArch
}

// GetOsName returns the os name the binary was built for
func GetOsName() string {
	return osName
}

// GetFullVersion returns the full version string
func GetFullVersion() string {
	return Version + "-" + Commit + " (" + Date + ") " + Builder + " " + goVersion + " " + osArch + " " + osName
}

// GetBuildInfo returns the build information
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   Version,
		Commit:    Commit,
		Date:      Date,
		Builder:   Builder,
		GoVersion: goVersion,
		OsArch:    osArch,
		OsName:    osName,
	}
}

// GetVersionTemplate returns the template for the version command
func GetVersionTemplate() string {
	return versionTemplate
}
