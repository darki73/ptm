package version

import (
	"runtime"
	"testing"
)

func TestGetVersion(t *testing.T) {
	expectedVersion := "dev"
	actualVersion := GetVersion()

	if actualVersion != expectedVersion {
		t.Errorf("Expected version: %s, but got: %s", expectedVersion, actualVersion)
	}
}

func TestGetCommit(t *testing.T) {
	expectedCommit := "none"
	actualCommit := GetCommit()

	if actualCommit != expectedCommit {
		t.Errorf("Expected commit: %s, but got: %s", expectedCommit, actualCommit)
	}
}

func TestGetDate(t *testing.T) {
	expectedDate := "0000-00-00T00:00:00Z"
	actualDate := GetDate()

	if actualDate != expectedDate {
		t.Errorf("Expected date: %s, but got: %s", expectedDate, actualDate)
	}
}

func TestGetBuilder(t *testing.T) {
	expectedBuilder := "unknown"
	actualBuilder := GetBuilder()

	if actualBuilder != expectedBuilder {
		t.Errorf("Expected builder: %s, but got: %s", expectedBuilder, actualBuilder)
	}
}

func TestGetGoVersion(t *testing.T) {
	expectedGoVersion := runtime.Version()
	actualGoVersion := GetGoVersion()

	if actualGoVersion != expectedGoVersion {
		t.Errorf("Expected Go version: %s, but got: %s", expectedGoVersion, actualGoVersion)
	}
}

func TestGetOsArch(t *testing.T) {
	expectedOsArch := runtime.GOARCH
	actualOsArch := GetOsArch()

	if actualOsArch != expectedOsArch {
		t.Errorf("Expected OS/Arch: %s, but got: %s", expectedOsArch, actualOsArch)
	}
}

func TestGetOsName(t *testing.T) {
	expectedOsName := runtime.GOOS
	actualOsName := GetOsName()

	if actualOsName != expectedOsName {
		t.Errorf("Expected OS name: %s, but got: %s", expectedOsName, actualOsName)
	}
}

func TestGetFullVersion(t *testing.T) {
	expectedFullVersion := "dev-none (0000-00-00T00:00:00Z) unknown " + runtime.Version() + " " + runtime.GOARCH + " " + runtime.GOOS
	actualFullVersion := GetFullVersion()

	if actualFullVersion != expectedFullVersion {
		t.Errorf("Expected full version: %s, but got: %s", expectedFullVersion, actualFullVersion)
	}
}

func TestGetBuildInfo(t *testing.T) {
	expectedBuildInfo := BuildInfo{
		Version:   "dev",
		Commit:    "none",
		Date:      "0000-00-00T00:00:00Z",
		Builder:   "unknown",
		GoVersion: runtime.Version(),
		OsArch:    runtime.GOARCH,
		OsName:    runtime.GOOS,
	}
	actualBuildInfo := GetBuildInfo()

	if actualBuildInfo != expectedBuildInfo {
		t.Errorf("Expected build info: %+v, but got: %+v", expectedBuildInfo, actualBuildInfo)
	}
}

func TestSetVersion(t *testing.T) {
	expectedVersion := "1.0.0"
	SetVersion(expectedVersion)
	actualVersion := GetVersion()

	if actualVersion != expectedVersion {
		t.Errorf("Expected version: %s, but got: %s", expectedVersion, actualVersion)
	}
}

func TestSetCommit(t *testing.T) {
	expectedCommit := "abcd1234"
	SetCommit(expectedCommit)
	actualCommit := GetCommit()

	if actualCommit != expectedCommit {
		t.Errorf("Expected commit: %s, but got: %s", expectedCommit, actualCommit)
	}
}

func TestSetDate(t *testing.T) {
	expectedDate := "2023-07-13T00:00:00Z"
	SetDate(expectedDate)
	actualDate := GetDate()

	if actualDate != expectedDate {
		t.Errorf("Expected date: %s, but got: %s", expectedDate, actualDate)
	}
}

func TestSetBuilder(t *testing.T) {
	expectedBuilder := "jenkins"
	SetBuilder(expectedBuilder)
	actualBuilder := GetBuilder()

	if actualBuilder != expectedBuilder {
		t.Errorf("Expected builder: %s, but got: %s", expectedBuilder, actualBuilder)
	}
}

func TestGetVersionTemplate(t *testing.T) {
	expectedVersionTemplate := `Version:      {{ .Version }}
SHA Commit:   {{ .Commit }}
Go version:   {{ .GoVersion }}
Built On:     {{ .Date }}
Built By:     {{ .Builder }}
OS/Arch:      {{ .OsName }}/{{ .OsArch }}
`
	actualVersionTemplate := GetVersionTemplate()

	if actualVersionTemplate != expectedVersionTemplate {
		t.Errorf("Expected version template: %s, but got: %s", expectedVersionTemplate, actualVersionTemplate)
	}
}
