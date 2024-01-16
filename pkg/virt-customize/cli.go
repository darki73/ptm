package virt_customize

import (
	"bytes"
	"fmt"
	uuc "github.com/darki73/ptm/pkg/configuration/unattended-upgrades"
	"github.com/darki73/ptm/pkg/virt-customize/command"
	uu "github.com/darki73/ptm/pkg/virt-customize/unattended-upgrades"
	"os"
	"os/exec"
	"sort"
)

const (
	// virtCustomizeCommand is the command to run virt-customize.
	virtCustomizeCommand = "virt-customize"
)

// CommandLineInterface is a structure that holds information for virt-customize CLI.
type CommandLineInterface struct {
	// configuration is the configuration of the virt-customize CLI.
	configuration *VirtCustomize
	// commands is the list of commands to run.
	commands map[int]*command.Command
	// cleanupFunctions is the list of cleanup functions.
	cleanupFunctions []func() error
}

// NewCommandLineInterface creates a new virt-customize CLI.
func NewCommandLineInterface(configuration *VirtCustomize) *CommandLineInterface {
	return &CommandLineInterface{
		configuration:    configuration,
		commands:         make(map[int]*command.Command),
		cleanupFunctions: make([]func() error, 0),
	}
}

// Execute executes the cli pipeline.
func (cli *CommandLineInterface) Execute() error {
	fmt.Println("Starting customization process for image:", cli.configuration.GetImage())
	if err := cli.buildCommandsList(); err != nil {
		return err
	}

	var keys []int
	for k := range cli.commands {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		if err := cli.executeCommand(cli.commands[k]); err != nil {
			// NOTE: We are ignoring the error for cleanup on purpose as the error from command execution is more important.
			_ = cli.cleanup()
			return err
		}
	}

	fmt.Println("Finished customization process for image:", cli.configuration.GetImage())

	return cli.cleanup()
}

// buildCommandsList builds the list of commands.
func (cli *CommandLineInterface) buildCommandsList() error {
	configuration := cli.configuration
	image := configuration.GetImage()

	cli.addCommand(command.NewUpdateCommand(image))
	cli.addCommand(command.NewInstallCommand(image, configuration.GetPackages()))

	for _, repository := range configuration.GetRepositoriesConfiguration() {
		cli.addCommand(command.NewAddGPGCommand(image, repository))
		cli.addCommand(command.NewAddRepositoryCommand(image, repository))
	}

	unattendedUpgradesConfiguration := configuration.GetUnattendedUpgradesConfiguration()
	if unattendedUpgradesConfiguration.GetEnabled() {
		if err := cli.uploadUnattendedUpgradesConfiguration(image, unattendedUpgradesConfiguration); err != nil {
			return err
		}

		if err := cli.uploadAutoUpgradesConfiguration(image); err != nil {
			return err
		}
	}

	cli.addCommand(command.NewUpdateCommand(image))

	return nil
}

// addCommand adds a command to the list of commands.
func (cli *CommandLineInterface) addCommand(command *command.Command) *CommandLineInterface {
	index := len(cli.commands) + 1
	command.SetOrder(index)

	cli.commands[index] = command
	return cli
}

// addCleanupFunction adds a cleanup function to the list of cleanup functions.
func (cli *CommandLineInterface) addCleanupFunction(cleanupFunction func() error) *CommandLineInterface {
	cli.cleanupFunctions = append(cli.cleanupFunctions, cleanupFunction)
	return cli
}

// cleanup cleans after the execution is done.
func (cli *CommandLineInterface) cleanup() error {
	for index, cleanupFunction := range cli.cleanupFunctions {
		if err := cleanupFunction(); err != nil {
			return err
		}
		delete(cli.commands, index)
	}
	return nil
}

// createUnattendedUpgradesTemporaryFile creates a temporary file for unattended upgrades.
func (cli *CommandLineInterface) createUnattendedUpgradesTemporaryFile(temporaryPath string, configuration string) (*os.File, error) {
	return cli.createTemporaryFile("unattended upgrades", temporaryPath, configuration)
}

// uploadUnattendedUpgradesConfiguration uploads the unattended upgrades configuration.
func (cli *CommandLineInterface) uploadUnattendedUpgradesConfiguration(image string, configuration *uuc.Configuration) error {
	unattendedUpgradesConfigurationFilePath := uu.GetUnattendedUpgradesConfigurationPath()
	unattendedUpgradesTemporaryFilePath := uu.GetUnattendedUpgradesTemporaryPath()
	unattendedUpgradesConfiguration, err := uu.BuildUnattendedUpgradesConfiguration(configuration)

	if err != nil {
		return err
	}

	unattendedUpgradesFileHandle, err := cli.createUnattendedUpgradesTemporaryFile(
		unattendedUpgradesTemporaryFilePath,
		unattendedUpgradesConfiguration,
	)
	if err != nil {
		return err
	}

	cli.addCleanupFunction(func() error {
		return os.Remove(unattendedUpgradesFileHandle.Name())
	})

	cli.addCommand(command.NewUploadCommand(image, unattendedUpgradesFileHandle.Name(), unattendedUpgradesConfigurationFilePath))

	return nil
}

// createAutoUpgradesTemporaryFile creates a temporary file for auto upgrades.
func (cli *CommandLineInterface) createAutoUpgradesTemporaryFile(temporaryPath string, configuration string) (*os.File, error) {
	return cli.createTemporaryFile("auto upgrades", temporaryPath, configuration)
}

// uploadAutoUpgradesConfiguration uploads the auto upgrades configuration.
func (cli *CommandLineInterface) uploadAutoUpgradesConfiguration(image string) error {
	autoUpgradesConfigurationFilePath := uu.GetAutoUpgradesConfigurationPath()
	autoUpgradesTemporaryFilePath := uu.GetAutoUpgradesTemporaryPath()
	autoUpgradesConfiguration := uu.GetAutoUpgradesTemplate()

	autoUpgradesFileHandle, err := cli.createAutoUpgradesTemporaryFile(
		autoUpgradesTemporaryFilePath,
		autoUpgradesConfiguration,
	)
	if err != nil {
		return err
	}

	cli.addCleanupFunction(func() error {
		return os.Remove(autoUpgradesFileHandle.Name())
	})

	cli.addCommand(command.NewUploadCommand(image, autoUpgradesFileHandle.Name(), autoUpgradesConfigurationFilePath))

	return nil
}

// createTemporaryFile creates a temporary file.
func (cli *CommandLineInterface) createTemporaryFile(actor string, temporaryPath string, configuration string) (*os.File, error) {
	if _, err := os.Stat(temporaryPath); err == nil {
		if err := os.Remove(temporaryPath); err != nil {
			return nil, fmt.Errorf("failed to remove temp file for %s: %v", actor, err)
		}
	}

	handle, err := os.Create(temporaryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file for %s: %v", actor, err)
	}
	defer handle.Close()

	if _, err := handle.WriteString(configuration); err != nil {
		return nil, fmt.Errorf("failed to write configuration to temp file for %s: %v", actor, err)
	}

	return handle, nil
}

// executeCommand executes a command.
func (cli *CommandLineInterface) executeCommand(command *command.Command) error {
	//fmt.Println("Executing command:", virtCustomizeCommand, strings.Join(command.BuildExecutionerCommand(), " "))
	cmd := exec.Command(virtCustomizeCommand, command.BuildExecutionerCommand()...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command execution failed: %v, stderr: %s", err, stderr.String())
	}

	return nil
}
