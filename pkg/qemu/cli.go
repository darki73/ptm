package qemu

import (
	"bytes"
	"fmt"
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"github.com/darki73/ptm/pkg/qemu/command"
	"os"
	"os/exec"
	"sort"
	"strings"
)

const (
	// qemuCommand is the command to run QEMU.
	qemuCommand = "qm"
)

// CommandLineInterface is a structure that holds information for QEMU CLI.
type CommandLineInterface struct {
	// configuration is the configuration of the QEMU CLI.
	configuration *Qemu
	// commands is the list of commands to run.
	commands map[int]*command.Command
	// cleanupFunctions is the list of cleanup functions.
	cleanupFunctions []func() error
}

// NewCommandLineInterface creates a new QEMU CLI.
func NewCommandLineInterface(configuration *Qemu) *CommandLineInterface {
	return &CommandLineInterface{
		configuration:    configuration,
		commands:         make(map[int]*command.Command),
		cleanupFunctions: make([]func() error, 0),
	}
}

// Execute executes the cli pipeline.
func (cli *CommandLineInterface) Execute() error {
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

	return cli.cleanup()
}

// buildCommandsList builds the list of commands.
func (cli *CommandLineInterface) buildCommandsList() error {
	configuration := cli.configuration

	identifier := configuration.GetIdentifier()

	cli.addCommand(command.NewNameCommand(identifier, configuration.GetName()))
	cli.addCommand(command.NewResourcesCommand(identifier, configuration.GetCores(), configuration.GetMemory(), configuration.GetCpuType()))
	cli.addCommand(command.NewGraphicsCommand(identifier))
	cli.addCommand(command.NewNetworkCommand(identifier, configuration.GetNetworkDriver(), configuration.GetNetworkBridge()))
	cli.addCommand(command.NewMainStorageCommand(identifier, configuration.GetStorage(), configuration.GetImage()))
	cli.addCommand(command.NewBootOrderCommand(identifier, "scsi0", "virtio-scsi-single"))
	cli.addCommand(command.NewGuestAgentCommand(identifier, true, true))

	cloudInit := configuration.GetCloudInit()

	if cloudInit != nil {
		cli.addCommand(command.NewCloudStorageCommand(identifier, configuration.GetStorage()))

		username := cloudInit.GetUsername()
		if username != "" {
			cli.addCommand(command.NewCloudUsernameCommand(identifier, cloudInit))
		}

		password := cloudInit.GetPassword()
		if password != "" {
			cli.addCommand(command.NewCloudPasswordCommand(identifier, cloudInit))
		}

		if len(cloudInit.GetKeys()) > 0 {
			fileHandle, err := cli.createTemporarySshKeysFile(cloudInit)
			if err != nil {
				return err
			}

			cli.addCleanupFunction(func() error {
				return os.Remove(fileHandle.Name())
			})

			cli.addCommand(command.NewCloudKeysCommand(identifier, cloudInit))
		}

		cli.addCommand(command.NewNetworkCloudCommand(identifier, cloudInit))
	}

	if configuration.IsResizingRequired() {
		cli.addCommand(command.NewResizeCommand(identifier, "scsi0", configuration.GetNewImageSizeAsString()))
	}

	cli.addCommand(command.NewTemplateCommand(identifier))

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

// createTemporarySshKeysFile creates a temporary file with SSH keys.
func (cli *CommandLineInterface) createTemporarySshKeysFile(cloudInit *ci.CloudInit) (*os.File, error) {
	if _, err := os.Stat(cloudInit.GetSSHKeysTemporaryFilePath()); err == nil {
		if err := os.Remove(cloudInit.GetSSHKeysTemporaryFilePath()); err != nil {
			return nil, fmt.Errorf("failed to remove temp file for SSH keys: %v", err)
		}
	}

	shellKeys := strings.Join(cloudInit.GetKeys(), "\n")

	handle, err := os.Create(cloudInit.GetSSHKeysTemporaryFilePath())
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file for SSH keys: %v", err)
	}
	defer handle.Close()

	if _, err := handle.WriteString(shellKeys); err != nil {
		return nil, fmt.Errorf("failed to write SSH keys to temp file: %v", err)
	}

	return handle, nil
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

// executeCommand executes a command.
func (cli *CommandLineInterface) executeCommand(command *command.Command) error {
	//fmt.Println("Executing command:", qemuCommand, strings.Join(command.BuildExecutionerCommand(), " "))
	cmd := exec.Command(qemuCommand, command.BuildExecutionerCommand()...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command execution failed: %v, stderr: %s", err, stderr.String())
	}

	return nil
}
