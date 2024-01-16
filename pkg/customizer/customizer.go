package customizer

import (
	"fmt"
	"github.com/cqroot/prompt/choose"
	config "github.com/darki73/ptm/pkg/configuration"
	"github.com/darki73/ptm/pkg/prompter"
	"github.com/darki73/ptm/pkg/proxmox"
	vc "github.com/darki73/ptm/pkg/virt-customize"
)

// Customizer represents the customizer.
type Customizer struct {
	// configuration represents the reference to the configuration.
	configuration *config.Configuration
	// virtCustomizeConfiguration represents the reference to the virt-customize configuration.
	virtCustomizeConfiguration *vc.VirtCustomize
	// selectedImage represents the selected image.
	selectedImage string
}

// NewCustomizer creates a new customizer instance.
func NewCustomizer(configuration *config.Configuration) *Customizer {
	return &Customizer{
		configuration:              configuration,
		virtCustomizeConfiguration: nil,
		selectedImage:              "",
	}
}

// Run runs the customizer.
func (customizer *Customizer) Run() error {
	if err := customizer.askForImageToCustomize(); err != nil {
		return err
	}

	customizer.virtCustomizeConfiguration = vc.NewVirtCustomizeConfiguration(
		customizer.selectedImage,
		customizer.configuration.GetBasePackages(),
		customizer.configuration.GetExtraPackages(),
		customizer.configuration.GetRepositories(),
		customizer.configuration.GetUnattendedUpgrades(),
	)

	cli := vc.NewCommandLineInterface(customizer.virtCustomizeConfiguration)

	return cli.Execute()
}

// askForImageToCustomize asks for the image to customize.
func (customizer *Customizer) askForImageToCustomize() error {
	choices := make([]choose.Choice, 0)

	images, err := proxmox.NewImages(customizer.configuration.GetDownloader().GetSaveTo())
	if err != nil {
		return err
	}

	for _, image := range images.GetISOs() {
		choice := choose.Choice{
			Text: image.GetFullPath(),
			Note: fmt.Sprintf(
				"Size: %d MB | Format: %s",
				image.GetVirtualSize(),
				image.GetQemuImageInformation().GetFormat(),
			),
		}
		choices = append(choices, choice)
	}

	result, err := prompter.PromptChoiceString(
		"Please select the target image for the virtual machine template:",
		choices,
	)

	if err != nil {
		return err
	}

	customizer.selectedImage = result

	return nil
}
