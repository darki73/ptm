package unattended_upgrades

import (
	"bytes"
	uu "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
	"text/template"
)

// unattendedUpgradesTemplate represents the template for the unattended-upgrades configuration.
// /etc/apt/apt.conf.d/50unattended-upgrades
const unattendedUpgradesTemplate = `
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

// GetUnattendedUpgradesTemplate returns the unattended-upgrades template.
func GetUnattendedUpgradesTemplate() string {
	return unattendedUpgradesTemplate
}

// GetUnattendedUpgradesConfigurationPath returns the path to the unattended-upgrades configuration.
func GetUnattendedUpgradesConfigurationPath() string {
	return "/etc/apt/apt.conf.d/50unattended-upgrades"
}

// GetUnattendedUpgradesTemporaryPath returns the path to the unattended-upgrades temporary file.
func GetUnattendedUpgradesTemporaryPath() string {
	return "/tmp/ptm-50unattended-upgrades"
}

// BuildUnattendedUpgradesConfiguration builds the unattended-upgrades configuration.
func BuildUnattendedUpgradesConfiguration(configuration *uu.Configuration) (string, error) {
	tmpl, err := template.New("unattended_upgrades").Parse(GetUnattendedUpgradesTemplate())
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	if err := tmpl.Execute(&buffer, configuration); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
