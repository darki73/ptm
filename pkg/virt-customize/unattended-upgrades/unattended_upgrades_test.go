package unattended_upgrades

import (
	uu "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
	"testing"
)

// TestGetUnattendedUpgradesTemplate tests the GetUnattendedUpgradesTemplate function.
func TestGetUnattendedUpgradesTemplate(t *testing.T) {
	expectedTemplate := `
Unattended-Upgrade::Allowed-Origins {
	{{- range .Whitelist }}
	"{{ . }}";
	{{- end }}
};

Unattended-Upgrade::Package-Blacklist {
	{{- range .Blacklist }}
	"{{ . }}";
	{{- end }}
};

Unattended-Upgrade::DevRelease "{{.DevRelease}}";
Unattended-Upgrade::AutoFixInterruptedDpkg "{{.FixInterrupted}}";
Unattended-Upgrade::MinimalSteps "{{.MinimalSteps}}";
Unattended-Upgrade::InstallOnShutdown "{{.InstallOnShutdown}}";
Unattended-Upgrade::Remove-Unused-Kernel-Packages "{{.RemoveUnusedKernel}}";
Unattended-Upgrade::Remove-New-Unused-Dependencies "{{.RemoveUnusedAutoDepend}}";
Unattended-Upgrade::Remove-Unused-Dependencies "{{.RemoveUnusedDependencies}}";
Unattended-Upgrade::Automatic-Reboot "{{.AutomaticReboot}}";
Unattended-Upgrade::Automatic-Reboot-WithUsers "{{.AutomaticRebootWithUsers}}";
Unattended-Upgrade::Automatic-Reboot-Time "{{.AutomaticRebootTime}}";
`
	result := GetUnattendedUpgradesTemplate()
	if result != expectedTemplate {
		t.Errorf("GetUnattendedUpgradesTemplate generated incorrect template.\nExpected:\n%s\n\nActual:\n%s", expectedTemplate, result)
	}
}

// TestGetUnattendedUpgradesConfigurationPath tests the GetUnattendedUpgradesConfigurationPath function.
func TestGetUnattendedUpgradesConfigurationPath(t *testing.T) {
	expectedPath := "/etc/apt/apt.conf.d/50unattended-upgrades"
	result := GetUnattendedUpgradesConfigurationPath()
	if result != expectedPath {
		t.Errorf("GetUnattendedUpgradesConfigurationPath generated incorrect path.\nExpected:\n%s\n\nActual:\n%s", expectedPath, result)
	}
}

// TestGetUnattendedUpgradesTemporaryPath tests the GetUnattendedUpgradesTemporaryPath function.
func TestGetUnattendedUpgradesTemporaryPath(t *testing.T) {
	expectedPath := "/tmp/ptm-50unattended-upgrades"
	result := GetUnattendedUpgradesTemporaryPath()
	if result != expectedPath {
		t.Errorf("GetUnattendedUpgradesTemporaryPath generated incorrect path.\nExpected:\n%s\n\nActual:\n%s", expectedPath, result)
	}
}

// TestBuildUnattendedUpgradesConfiguration tests the BuildUnattendedUpgradesConfiguration function.
func TestBuildUnattendedUpgradesConfiguration(t *testing.T) {
	config := &uu.Configuration{
		Whitelist:                []string{"origin1", "origin2"},
		Blacklist:                []string{"package1", "package2"},
		DevRelease:               "auto",
		FixInterrupted:           true,
		MinimalSteps:             true,
		InstallOnShutdown:        true,
		RemoveUnusedKernel:       true,
		RemoveUnusedAutoDepend:   true,
		RemoveUnusedDependencies: true,
		AutomaticReboot:          true,
		AutomaticRebootWithUsers: true,
		AutomaticRebootTime:      "02:00",
	}

	result, err := BuildUnattendedUpgradesConfiguration(config)
	if err != nil {
		t.Fatalf("BuildUnattendedUpgradesConfiguration returned an error: %v", err)
	}

	expectedConfig := `
Unattended-Upgrade::Allowed-Origins {
	"origin1";
	"origin2";
};

Unattended-Upgrade::Package-Blacklist {
	"package1";
	"package2";
};

Unattended-Upgrade::DevRelease "auto";
Unattended-Upgrade::AutoFixInterruptedDpkg "true";
Unattended-Upgrade::MinimalSteps "true";
Unattended-Upgrade::InstallOnShutdown "true";
Unattended-Upgrade::Remove-Unused-Kernel-Packages "true";
Unattended-Upgrade::Remove-New-Unused-Dependencies "true";
Unattended-Upgrade::Remove-Unused-Dependencies "true";
Unattended-Upgrade::Automatic-Reboot "true";
Unattended-Upgrade::Automatic-Reboot-WithUsers "true";
Unattended-Upgrade::Automatic-Reboot-Time "02:00";
`

	if result != expectedConfig {
		t.Errorf("BuildUnattendedUpgradesConfiguration generated incorrect configuration.\nExpected:\n%s\n\nActual:\n%s", expectedConfig, result)
	}
}
