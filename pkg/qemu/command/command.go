package command

import (
	"fmt"
	"strconv"
)

const (
	// qemuCommandCreate is the command to create a new QEMU VM.
	qemuCommandCreate = "create"
	// qemuCommandSet is the command to set a QEMU VM property.
	qemuCommandSet = "set"
	// qemuCommandDisk is the command to manage QEMU VM disks.
	qemuCommandDisk = "disk"
	// qemuCommandTemplate is the command to manage QEMU VM templates.
	qemuCommandTemplate = "template"
)

// Command is a structure that holds information for QEMU command.
type Command struct {
	// order is the order of the command.
	order int
	// command is the command to run.
	command string
	// identifier is the QEMU VM identifier.
	identifier int
	// arguments is the list of arguments to pass to the command.
	arguments []string
}

// NewCommand creates a new QEMU command.
func NewCommand(command string, identifier int, arguments ...interface{}) *Command {
	argumentsList := make([]string, 0)

	for _, argument := range arguments {
		if argument == nil {
			continue
		}
		argumentsList = append(argumentsList, fmt.Sprintf("%v", argument))
	}

	return &Command{
		order:      0,
		command:    command,
		identifier: identifier,
		arguments:  argumentsList,
	}
}

// NewCreateCommand creates a new QEMU create command.
func NewCreateCommand(identifier int, arguments ...interface{}) *Command {
	return NewCommand(qemuCommandCreate, identifier, arguments...)
}

// NewSetCommand creates a new QEMU set command.
func NewSetCommand(identifier int, arguments ...interface{}) *Command {
	return NewCommand(qemuCommandSet, identifier, arguments...)
}

// NewDiskCommand creates a new QEMU disk command.
func NewDiskCommand(identifier int, arguments ...interface{}) *Command {
	return NewCommand(qemuCommandDisk, identifier, arguments...)
}

// NewTemplateCommand creates a new QEMU template command.
func NewTemplateCommand(identifier int, arguments ...interface{}) *Command {
	return NewCommand(qemuCommandTemplate, identifier, arguments...)
}

// SetOrder sets the order of the command.
func (command *Command) SetOrder(order int) *Command {
	command.order = order
	return command
}

// GetOrder returns the order of the command.
func (command *Command) GetOrder() int {
	return command.order
}

// GetCommand returns the command to run.
func (command *Command) GetCommand() string {
	return command.command
}

// GetIdentifier returns the QEMU VM identifier.
func (command *Command) GetIdentifier() int {
	return command.identifier
}

// GetArguments returns the list of arguments to pass to the command.
func (command *Command) GetArguments() []string {
	return command.arguments
}

// BuildExecutionerCommand builds the command to run.
func (command *Command) BuildExecutionerCommand() []string {
	commandParts := []string{
		command.GetCommand(),
	}

	if command.GetCommand() != qemuCommandDisk {
		commandParts = append(commandParts, strconv.Itoa(command.GetIdentifier()))

		for _, argument := range command.GetArguments() {
			commandParts = append(commandParts, argument)
		}
	} else {
		for _, argument := range command.GetArguments() {
			if argument == "%VM_ID%" {
				argument = strconv.Itoa(command.GetIdentifier())
			}
			commandParts = append(commandParts, argument)
		}
	}

	return commandParts
}
