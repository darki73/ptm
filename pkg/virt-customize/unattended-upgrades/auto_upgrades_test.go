package unattended_upgrades

import "testing"

// TestGetAutoUpgradesTemplate tests the GetAutoUpgradesTemplate function.
func TestGetAutoUpgradesTemplate(t *testing.T) {
	expectedTemplate := `
APT::Periodic::Update-Package-Lists "1";
APT::Periodic::Download-Upgradeable-Packages "1";
APT::Periodic::AutocleanInterval "7";
APT::Periodic::Unattended-Upgrade "1";
`
	result := GetAutoUpgradesTemplate()
	if result != expectedTemplate {
		t.Errorf("GetAutoUpgradesTemplate generated incorrect template.\nExpected:\n%s\n\nActual:\n%s", expectedTemplate, result)
	}
}

// TestGetAutoUpgradesConfigurationPath tests the GetAutoUpgradesConfigurationPath function.
func TestGetAutoUpgradesConfigurationPath(t *testing.T) {
	expectedPath := "/etc/apt/apt.conf.d/20auto-upgrades"
	result := GetAutoUpgradesConfigurationPath()
	if result != expectedPath {
		t.Errorf("GetAutoUpgradesConfigurationPath generated incorrect path.\nExpected:\n%s\n\nActual:\n%s", expectedPath, result)
	}
}

// TestGetAutoUpgradesTemporaryPath tests the GetAutoUpgradesTemporaryPath function.
func TestGetAutoUpgradesTemporaryPath(t *testing.T) {
	expectedPath := "/tmp/ptm-20auto-upgrades"
	result := GetAutoUpgradesTemporaryPath()
	if result != expectedPath {
		t.Errorf("GetAutoUpgradesTemporaryPath generated incorrect path.\nExpected:\n%s\n\nActual:\n%s", expectedPath, result)
	}
}
