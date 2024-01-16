package maker

import (
	"fmt"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
	config "github.com/darki73/ptm/pkg/configuration"
	"github.com/darki73/ptm/pkg/prompter"
	"github.com/darki73/ptm/pkg/proxmox"
	"github.com/darki73/ptm/pkg/qemu"
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"github.com/darki73/ptm/pkg/utils"
)

// Maker represents the maker struct.
type Maker struct {
	// configuration represents the reference to the configuration.
	configuration *config.Configuration
	// qemuConfiguration represents the reference to the QEMU configuration.
	qemuConfiguration *qemu.Qemu
	// cloudInitConfiguration represents the reference to the cloud-init configuration.
	cloudInitConfiguration *ci.CloudInit
	// availableCoreCount represents the number of available cores.
	availableCoreCount int
	// availableMemory represents the amount of available memory.
	availableMemory uint64
	// images represents the reference to the images.
	images *proxmox.Images
	// storage represents the reference to the storage.
	storage *proxmox.Storage
	// keys represents the reference to the shell keys.
	keys *proxmox.ShellKeys
}

// NewMaker creates a new maker instance
func NewMaker(configuration *config.Configuration, qemuConfiguration *qemu.Qemu) (*Maker, error) {
	var cloudInitConfiguration *ci.CloudInit

	images, err := proxmox.NewImages(configuration.GetDownloader().GetSaveTo())
	if err != nil {
		return nil, err
	}

	storage, err := proxmox.NewStorage()
	if err != nil {
		return nil, err
	}

	keys, err := proxmox.NewShellKeys()
	if err != nil {
		return nil, err
	}

	if qemuConfiguration != nil {
		if qemuConfiguration.GetCloudInit() != nil {
			cloudInitConfiguration = qemuConfiguration.GetCloudInit()
		} else {
			cloudInitConfiguration = nil
		}
	}

	return &Maker{
		configuration:          configuration,
		qemuConfiguration:      qemuConfiguration,
		cloudInitConfiguration: cloudInitConfiguration,
		availableCoreCount:     utils.GetCoreCount(),
		availableMemory:        utils.GetTotalMemory(),
		images:                 images,
		storage:                storage,
		keys:                   keys,
	}, nil
}

// Run runs the maker.
func (maker *Maker) Run() error {
	if err := maker.handleQemuConfigurationLogic(); err != nil {
		return err
	}

	if err := maker.handleCloudInitConfigurationLogic(); err != nil {
		return err
	}

	cli := qemu.NewCommandLineInterface(maker.qemuConfiguration)

	return cli.Execute()
}

// askForTemplateIdentifier asks for the template identifier.
func (maker *Maker) askForTemplateIdentifier() error {
	result, err := prompter.PromptInteger(
		"Please enter the identifier for the virtual machine template",
		9000,
	)

	if err != nil {
		return err
	}

	maker.qemuConfiguration.SetIdentifier(result)

	return nil
}

// askForTemplateName asks for the template name.
func (maker *Maker) askForTemplateName() error {
	result, err := prompter.PromptString(
		"Please enter the name for the virtual machine template",
		"ptm-template",
	)

	if err != nil {
		return err
	}

	maker.qemuConfiguration.SetName(result)

	return nil
}

// askForCoreCount asks for the core count.
func (maker *Maker) askForCoreCount() error {
	result, err := prompter.PromptInteger(
		fmt.Sprintf(
			"Please enter the core count for the virtual machine template [1-%d]",
			maker.availableCoreCount,
		),
		1,
	)

	if err != nil {
		return err
	}

	return maker.handleCoreCountSelectionLogic(result)
}

// askForMemory asks for the memory.
func (maker *Maker) askForMemory() error {
	resultString, err := prompter.PromptString(
		fmt.Sprintf(
			"Please enter the memory for the virtual machine template [1-%d MB]",
			maker.availableMemory,
		),
		"512M",
	)

	if err != nil {
		return err
	}

	result, err := utils.ConvertToMegabytes(resultString)
	if err != nil {
		return err
	}

	return maker.handleMemorySelectionLogicFromInt(int(result))
}

// askForCpuType asks for the CPU type.
func (maker *Maker) askForCpuType() error {
	result, err := prompter.PromptString(
		"Please enter the CPU type for the virtual machine template",
		"host",
	)

	if err != nil {
		return err
	}

	maker.qemuConfiguration.SetCpuType(result)

	return nil
}

// askForNetworkDriver asks for the network driver.
func (maker *Maker) askForNetworkDriver() error {
	result, err := prompter.PromptString(
		"Please enter the network driver for the virtual machine template",
		"virtio",
	)

	if err != nil {
		return err
	}

	maker.qemuConfiguration.SetNetworkDriver(result)

	return nil
}

// askForNetworkBridge asks for the network bridge.
func (maker *Maker) askForNetworkBridge() error {
	result, err := prompter.PromptString(
		"Please enter the network bridge for the virtual machine template",
		"vmbr0",
	)

	if err != nil {
		return err
	}

	maker.qemuConfiguration.SetNetworkBridge(result)

	return nil
}

// askForStorage asks for the storage.
func (maker *Maker) askForStorage() error {
	choices := make([]choose.Choice, 0)

	for _, storage := range maker.storage.GetTargets() {
		choice := choose.Choice{
			Text: storage.GetName(),
			Note: fmt.Sprintf(
				"Available: %d MB | Used: %s",
				storage.GetAvailable(),
				storage.GetPercentUsed(),
			),
		}
		choices = append(choices, choice)
	}

	result, err := prompter.PromptChoiceString(
		"Please select the storage for the virtual machine template",
		choices,
	)

	if err != nil {
		return err
	}

	return maker.handleStorageSelectionLogic(result)
}

// askForTargetImage asks for the target image.
func (maker *Maker) askForTargetImage() error {
	choices := make([]choose.Choice, 0)

	for _, image := range maker.images.GetISOs() {
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
		"Please select the target image for the virtual machine template",
		choices,
	)

	if err != nil {
		return err
	}

	return maker.handleImageSelectionLogic(result)
}

// askWhetherToResizeImage asks whether to resize the image.
func (maker *Maker) askWhetherToResizeImage() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to resize the image?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForImageNewSize(); err != nil {
			return err
		}
	}

	return nil
}

// askForImageNewSize asks for the image new size.
func (maker *Maker) askForImageNewSize() error {
	result, err := prompter.PromptString(
		"Please enter the new size for the image",
		"4G",
	)

	if err != nil {
		return err
	}

	return maker.handleImageResizeLogic(maker.qemuConfiguration.GetStorage(), result)
}

// askWhetherToConfigureCloudInit asks whether to configure the cloud-init.
func (maker *Maker) askWhetherToConfigureCloudInit() (bool, error) {
	return prompter.PromptChoiceYesNo(
		"Would you like to configure the cloud-init?",
	)
}

// askWhetherToSetCloudInitUsername asks whether to set the cloud-init username.
func (maker *Maker) askWhetherToSetCloudInitUsername() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to set the cloud-init username?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitUsername(); err != nil {
			return err
		}
	}

	return nil
}

// askForCloudInitUsername asks for the cloud-init username.
func (maker *Maker) askForCloudInitUsername() error {
	result, err := prompter.PromptString(
		"Please enter the username for the cloud-init configuration",
		"",
	)

	if err != nil {
		return err
	}

	maker.cloudInitConfiguration.SetUsername(result)

	return nil
}

// askWhetherToSetCloudInitPassword asks whether to set the cloud-init password.
func (maker *Maker) askWhetherToSetCloudInitPassword() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to set the cloud-init password?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitPassword(); err != nil {
			return err
		}
	}

	return nil
}

// askForCloudInitPassword asks for the cloud-init password.
func (maker *Maker) askForCloudInitPassword() error {
	result, err := prompter.PromptString(
		"Please enter the password for the cloud-init configuration",
		"",
		input.WithEchoMode(input.EchoPassword),
	)

	if err != nil {
		return err
	}

	maker.cloudInitConfiguration.SetPassword(result)

	return nil
}

// askWhetherToSetCloudInitSSHKeys asks whether to set the cloud-init SSH keys.
func (maker *Maker) askWhetherToSetCloudInitSSHKeys() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to set the cloud-init SSH keys?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitSSHKeys(); err != nil {
			return err
		}
	}

	return nil
}

// askForCloudInitSSHKeys asks for the cloud-init SSH keys.
func (maker *Maker) askForCloudInitSSHKeys() error {
	choices := make([]string, 0)

	for _, key := range maker.keys.GetKeys() {
		choices = append(choices, key.GetName())
	}

	result, err := prompter.PromptMultiChoice(
		"Please select the shell keys for the virtual machine template:",
		choices,
	)

	if err != nil {
		return err
	}

	var shellKeys []string

	for _, selectedShellKey := range result {
		keyReference := maker.keys.FindKeyByName(selectedShellKey)
		shellKeys = append(shellKeys, keyReference.GetFullPath())
	}

	maker.cloudInitConfiguration.SetKeys(shellKeys)

	return nil
}

// askWhetherToConfigureCloudInitIPv4 asks whether to configure the cloud-init IPv4.
func (maker *Maker) askWhetherToConfigureCloudInitIPv4() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to configure the cloud-init IPv4?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitIPv4Address(); err != nil {
			return err
		}
		if err := maker.askForCloudInitIPv4Gateway(); err != nil {
			return err
		}
	} else {
		maker.cloudInitConfiguration.AutoConfigureIPv4()
	}

	return nil
}

// askWhetherToReconfigureCloudInitIPv4 asks whether to reconfigure the cloud-init IPv4.
func (maker *Maker) askWhetherToReconfigureCloudInitIPv4() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to reconfigure the cloud-init IPv4?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitIPv4Address(); err != nil {
			return err
		}
		if err := maker.askForCloudInitIPv4Gateway(); err != nil {
			return err
		}
	}

	return nil
}

// askForCloudInitIPv4Address asks for the cloud-init IPv4 address.
func (maker *Maker) askForCloudInitIPv4Address() error {
	result, err := prompter.PromptString(
		"Please enter the IPv4 address for the cloud-init configuration",
		"",
	)

	if err != nil {
		return err
	}

	if !utils.IsValidIPWithSubnet(result) {
		fmt.Println("The specified IPv4 address is invalid.")
		return maker.askForCloudInitIPv4Address()
	}

	maker.cloudInitConfiguration.SetIPv4(result)

	return nil
}

// askForCloudInitIPv4Gateway asks for the cloud-init IPv4 gateway.
func (maker *Maker) askForCloudInitIPv4Gateway() error {
	result, err := prompter.PromptString(
		"Please enter the IPv4 gateway for the cloud-init configuration",
		"",
	)

	if err != nil {
		return err
	}

	if !utils.IsValidIP(result) {
		fmt.Println("The specified IPv4 gateway is invalid.")
		return maker.askForCloudInitIPv4Gateway()
	}

	if !utils.IsInSameNetwork(result, maker.cloudInitConfiguration.GetIPv4()) {
		fmt.Println("The specified IPv4 gateway is not in the same network as the IPv4 address.")
		return maker.askForCloudInitIPv4Address()
	}

	maker.cloudInitConfiguration.SetIPv4Gateway(result)

	return nil
}

// askWhetherToConfigureCloudInitIPv6 asks whether to configure the cloud-init IPv6.
func (maker *Maker) askWhetherToConfigureCloudInitIPv6() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to configure the cloud-init IPv6?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitIPv6Address(); err != nil {
			return err
		}
		if err := maker.askForCloudInitIPv6Gateway(); err != nil {
			return err
		}
	} else {
		maker.cloudInitConfiguration.AutoConfigureIPv6()
	}

	return nil
}

// askWhetherToReconfigureCloudInitIPv6 asks whether to reconfigure the cloud-init IPv6.
func (maker *Maker) askWhetherToReconfigureCloudInitIPv6() error {
	result, err := prompter.PromptChoiceYesNo(
		"Would you like to reconfigure the cloud-init IPv6?",
	)

	if err != nil {
		return err
	}

	if result {
		if err := maker.askForCloudInitIPv6Address(); err != nil {
			return err
		}
		if err := maker.askForCloudInitIPv6Gateway(); err != nil {
			return err
		}
	}

	return nil
}

// askForCloudInitIPv6Address asks for the cloud-init IPv6 address.
func (maker *Maker) askForCloudInitIPv6Address() error {
	result, err := prompter.PromptString(
		"Please enter the IPv6 address for the cloud-init configuration",
		"",
	)

	if err != nil {
		return err
	}

	if !utils.IsValidIPWithSubnet(result) {
		fmt.Println("The specified IPv6 address is invalid.")
		return maker.askForCloudInitIPv6Address()
	}

	maker.cloudInitConfiguration.SetIPv6(result)

	return nil
}

// askForCloudInitIPv6Gateway asks for the cloud-init IPv6 gateway.
func (maker *Maker) askForCloudInitIPv6Gateway() error {
	result, err := prompter.PromptString(
		"Please enter the IPv6 gateway for the cloud-init configuration",
		"",
	)

	if err != nil {
		return err
	}

	if !utils.IsValidIP(result) {
		fmt.Println("The specified IPv6 gateway is invalid.")
		return maker.askForCloudInitIPv6Gateway()
	}

	if !utils.IsInSameNetwork(result, maker.cloudInitConfiguration.GetIPv6()) {
		fmt.Println("The specified IPv6 gateway is not in the same network as the IPv6 address.")
		return maker.askForCloudInitIPv6Address()
	}

	maker.cloudInitConfiguration.SetIPv6Gateway(result)

	return nil
}

// handleQemuConfigurationLogic handles the QEMU configuration logic.
func (maker *Maker) handleQemuConfigurationLogic() error {
	if maker.qemuConfiguration == nil {
		return maker.handleQemuConfigurationFromPrompt()
	}

	isValidConfiguration, err := maker.qemuConfiguration.IsConfigurationValid()
	if err != nil || !isValidConfiguration {
		if err.Error() != "missing image for virtual machine" && maker.qemuConfiguration.GetConfigurationSource() != qemu.ConfigurationSourceConfigurationFile {
			return maker.handleQemuConfigurationFromPrompt()
		}
	}

	configurationSource := maker.qemuConfiguration.GetConfigurationSource()

	switch configurationSource {
	case qemu.ConfigurationSourceFlags:
		return maker.handleQemuConfigurationFromFlags()
	case qemu.ConfigurationSourceConfigurationFile:
		return maker.handleQemuConfigurationFromConfigurationFile()
	default:
		return fmt.Errorf("invalid configuration source: %s", configurationSource)
	}
}

// handleQemuConfigurationFromFlags handles the QEMU configuration from flags.
func (maker *Maker) handleQemuConfigurationFromFlags() error {
	if err := maker.handleCoreCountSelectionLogic(maker.qemuConfiguration.GetCores()); err != nil {
		return err
	}

	if err := maker.handleMemorySelectionLogicFromInt(maker.qemuConfiguration.GetMemory()); err != nil {
		return err
	}

	if err := maker.handleStorageSelectionLogic(maker.qemuConfiguration.GetStorage()); err != nil {
		return err
	}

	if err := maker.handleImageSelectionLogic(maker.qemuConfiguration.GetImage()); err != nil {
		return err
	}

	if maker.qemuConfiguration.GetNewImageSizeAsString() != "" {
		if err := maker.handleImageResizeLogic(
			maker.qemuConfiguration.GetStorage(),
			maker.qemuConfiguration.GetNewImageSizeAsString(),
		); err != nil {
			return err
		}
	}
	return nil
}

// handleQemuConfigurationFromConfigurationFile handles the QEMU configuration from configuration file.
func (maker *Maker) handleQemuConfigurationFromConfigurationFile() error {
	if err := maker.handleCoreCountSelectionLogic(maker.qemuConfiguration.GetCores()); err != nil {
		return err
	}

	if err := maker.handleMemorySelectionLogicFromInt(maker.qemuConfiguration.GetMemory()); err != nil {
		return err
	}

	if err := maker.handleStorageSelectionLogic(maker.qemuConfiguration.GetStorage()); err != nil {
		return err
	}

	if maker.qemuConfiguration.GetImage() == "" {
		if err := maker.askForTargetImage(); err != nil {
			return err
		}
	}

	if err := maker.handleImageSelectionLogic(maker.qemuConfiguration.GetImage()); err != nil {
		return err
	}

	if maker.qemuConfiguration.GetNewImageSizeAsString() != "" {
		if err := maker.handleImageResizeLogic(
			maker.qemuConfiguration.GetStorage(),
			maker.qemuConfiguration.GetNewImageSizeAsString(),
		); err != nil {
			return err
		}
	}
	return nil
}

// handleQemuConfigurationFromPrompt handles the QEMU configuration from prompt.
func (maker *Maker) handleQemuConfigurationFromPrompt() error {
	maker.qemuConfiguration = qemu.NewQemuConfiguration()

	if err := maker.askForTemplateIdentifier(); err != nil {
		return err
	}

	if err := maker.askForTemplateName(); err != nil {
		return err
	}

	if err := maker.askForCoreCount(); err != nil {
		return err
	}

	if err := maker.askForMemory(); err != nil {
		return err
	}

	if err := maker.askForCpuType(); err != nil {
		return err
	}

	if err := maker.askForNetworkDriver(); err != nil {
		return err
	}

	if err := maker.askForNetworkBridge(); err != nil {
		return err
	}

	if err := maker.askForStorage(); err != nil {
		return err
	}

	if err := maker.askForTargetImage(); err != nil {
		return err
	}

	if err := maker.askWhetherToResizeImage(); err != nil {
		return err
	}

	return nil
}

// handleCloudInitConfigurationLogic handles the cloud-init configuration logic.
func (maker *Maker) handleCloudInitConfigurationLogic() error {
	if maker.cloudInitConfiguration == nil {
		return maker.handleCloudInitConfigurationFromPrompt()
	}

	isValidConfiguration, err := maker.cloudInitConfiguration.IsConfigurationValid()
	if err != nil || !isValidConfiguration {
		return maker.handleCloudInitConfigurationFromPrompt()
	}

	configurationSource := maker.cloudInitConfiguration.GetConfigurationSource()

	switch configurationSource {
	case ci.ConfigurationSourceFlags:
		return maker.handleCloudInitConfigurationFromFlags()
	case ci.ConfigurationSourceConfigurationFile:
		return maker.handleCloudInitConfigurationFromConfigurationFile()
	default:
		return fmt.Errorf("cloud-init invalid configuration source: %s", configurationSource)
	}
}

// handleCloudInitConfigurationFromFlags handles the cloud-init configuration from flags.
func (maker *Maker) handleCloudInitConfigurationFromFlags() error {
	return nil
}

// handleCloudInitConfigurationFromConfigurationFile handles the cloud-init configuration from configuration file.
func (maker *Maker) handleCloudInitConfigurationFromConfigurationFile() error {
	cloudInitConfiguration := maker.cloudInitConfiguration

	if cloudInitConfiguration.GetUsername() == "" {
		if err := maker.askWhetherToSetCloudInitUsername(); err != nil {
			return err
		}
	}

	if cloudInitConfiguration.GetPassword() == "" {
		if err := maker.askWhetherToSetCloudInitPassword(); err != nil {
			return err
		}
	}

	if !cloudInitConfiguration.HasKeys() {
		if err := maker.askWhetherToSetCloudInitSSHKeys(); err != nil {
			return err
		}
	}

	if err := maker.askWhetherToReconfigureCloudInitIPv4(); err != nil {
		return err
	}

	if err := maker.askWhetherToReconfigureCloudInitIPv6(); err != nil {
		return err
	}

	return nil
}

// handleCloudInitConfigurationFromPrompt handles the cloud-init configuration from prompt.
func (maker *Maker) handleCloudInitConfigurationFromPrompt() error {
	maker.cloudInitConfiguration = ci.NewCloudInitConfiguration()

	configureCloudInit, err := maker.askWhetherToConfigureCloudInit()
	if err != nil {
		return err
	}

	if !configureCloudInit {
		maker.cloudInitConfiguration.AutoConfigureIPv4()
		maker.cloudInitConfiguration.AutoConfigureIPv6()
	} else {
		if err := maker.askWhetherToSetCloudInitUsername(); err != nil {
			return err
		}

		if err := maker.askWhetherToSetCloudInitPassword(); err != nil {
			return err
		}

		if err := maker.askWhetherToSetCloudInitSSHKeys(); err != nil {
			return err
		}

		if err := maker.askWhetherToConfigureCloudInitIPv4(); err != nil {
			return err
		}

		if err := maker.askWhetherToConfigureCloudInitIPv6(); err != nil {
			return err
		}
	}

	maker.qemuConfiguration.SetCloudInit(maker.cloudInitConfiguration)

	return nil
}

// handleStorageSelectionLogic handles the storage selection logic.
func (maker *Maker) handleStorageSelectionLogic(storageName string) error {
	storageReference := maker.storage.FindTargetByName(storageName)
	if storageReference == nil {
		return fmt.Errorf("storage `%s` could not be found", storageName)
	}

	maker.qemuConfiguration.SetStorage(storageName)
	maker.qemuConfiguration.SetStorageSize(storageReference.GetAvailable())

	return nil
}

// handleImageSelectionLogic handles the image selection logic.
func (maker *Maker) handleImageSelectionLogic(imageName string) error {
	imageReference := maker.images.FindISOByFullPath(imageName)
	if imageReference == nil {
		return fmt.Errorf("image `%s` could not be found", imageName)
	}

	maker.qemuConfiguration.SetImage(imageName)
	maker.qemuConfiguration.SetImageSize(imageReference.GetSize())

	return nil
}

// handleImageResizeLogic handles the image resize logic.
func (maker *Maker) handleImageResizeLogic(storage string, newImageSize string) error {
	kilobytes, err := utils.ConvertToKilobytes(newImageSize)
	if err != nil {
		return err
	}

	if maker.qemuConfiguration.GetImageSize() > kilobytes {
		fmt.Println("The image is already larger than the specified size.")
	}

	storageReference := maker.storage.FindTargetByName(storage)

	if storageReference == nil {
		return fmt.Errorf("storage `%s` could not be found", storage)
	}

	if !storageReference.HasEnoughSpace(kilobytes) {
		return fmt.Errorf(fmt.Sprintf(
			"Not enough space on the storage. Requested: %d MB, Available: %d MB",
			kilobytes/1024,
			storageReference.GetAvailable()/1024,
		))
	}

	maker.qemuConfiguration.SetNewImageSize(kilobytes)
	maker.qemuConfiguration.SetNewImageSizeAsString(newImageSize)

	return nil
}

// handleCoreCountSelectionLogic handles the core count selection logic.
func (maker *Maker) handleCoreCountSelectionLogic(coreCount int) error {
	if coreCount > maker.availableCoreCount {
		errorMessage := fmt.Sprintf(
			"Not enough cores available. Requested: %d, Available: %d",
			coreCount,
			maker.availableCoreCount,
		)

		if maker.isPromptConfigurationFlow() {
			fmt.Println(errorMessage)
			return maker.askForCoreCount()
		}

		return fmt.Errorf(errorMessage)
	}

	if coreCount < 1 {
		errorMessage := "Core count must be greater than 0"

		if maker.isPromptConfigurationFlow() {
			fmt.Println(errorMessage)
			return maker.askForCoreCount()
		}

		return fmt.Errorf(errorMessage)
	}

	maker.qemuConfiguration.SetCores(coreCount)

	return nil
}

// handleMemorySelectionLogicFromInt handles the memory selection logic (from int).
func (maker *Maker) handleMemorySelectionLogicFromInt(memory int) error {
	if memory > int(maker.availableMemory) {
		errorMessage := fmt.Sprintf(
			"Not enough memory available. Requested: %d MB, Available: %d MB",
			memory,
			maker.availableMemory,
		)

		if maker.isPromptConfigurationFlow() {
			fmt.Println(errorMessage)
			return maker.askForMemory()
		}

		return fmt.Errorf(errorMessage)
	}

	if memory < 1 {
		errorMessage := "Memory must be greater than 0"

		if maker.isPromptConfigurationFlow() {
			fmt.Println(errorMessage)
			return maker.askForMemory()
		}

		return fmt.Errorf(errorMessage)
	}

	maker.qemuConfiguration.SetMemory(memory)

	return nil
}

// isPromptConfigurationFlow checks whether we are using the prompt configuration flow.
func (maker *Maker) isPromptConfigurationFlow() bool {
	return maker.qemuConfiguration.GetConfigurationSource() == qemu.ConfigurationSourcePrompt
}
