package command

import (
	"fmt"
)

// Command is a structure that holds information for VirtCustomize command.
type Command struct {
	// order is the order of the command.
	order int
	// image is the path to the image to customize.
	image string
	// command is the command to run.
	command string
	// arguments is the list of arguments to pass to the command.
	arguments []string
}

// NewCommand creates a new VirtCustomize command.
func NewCommand(image string, command string, arguments ...interface{}) *Command {
	argumentsList := make([]string, 0)

	for _, argument := range arguments {
		if argument == nil {
			continue
		}
		argumentsList = append(argumentsList, fmt.Sprintf("%v", argument))
	}

	return &Command{
		order:     0,
		image:     image,
		command:   command,
		arguments: argumentsList,
	}
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

// GetImage returns the path to the image to customize.
func (command *Command) GetImage() string {
	return command.image
}

// GetCommand returns the command to run.
func (command *Command) GetCommand() string {
	return command.command
}

// GetArguments returns the list of arguments to pass to the command.
func (command *Command) GetArguments() []string {
	return command.arguments
}

// BuildExecutionerCommand builds the command to run.
func (command *Command) BuildExecutionerCommand() []string {
	commandParts := []string{
		"-a",
		command.GetImage(),
		command.GetCommand(),
	}

	commandParts = append(commandParts)

	for _, argument := range command.GetArguments() {
		commandParts = append(commandParts, argument)
	}

	return commandParts
}
