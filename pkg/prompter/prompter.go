package prompter

import (
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
	"github.com/cqroot/prompt/multichoose"
	"strconv"
	"strings"
)

// PromptString prompts user for string response.
func PromptString(message string, defaultValue string, opts ...input.Option) (string, error) {
	return prompt.New().Ask(appendColon(message)).Input(defaultValue, opts...)
}

// PromptInteger prompts user for integer response.
func PromptInteger(message string, defaultValue int, opts ...input.Option) (int, error) {
	opts = append(opts, input.WithInputMode(input.InputInteger))

	resultString, err := prompt.New().Ask(appendColon(message)).Input(
		strconv.Itoa(defaultValue),
		opts...,
	)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(resultString)
}

// PromptChoiceString prompts user for choice response (while returning string).
func PromptChoiceString(message string, choices []choose.Choice) (string, error) {
	return prompt.New().Ask(appendColon(message)).AdvancedChoose(choices)
}

// PromptChoiceYesNo prompts user for choice response (while returning boolean).
func PromptChoiceYesNo(message string) (bool, error) {
	value, err := prompt.New().Ask(appendColon(message)).Choose([]string{"Yes", "No"})
	if err != nil {
		return false, err
	}

	return value == "Yes", nil
}

// PromptMultiChoice prompts user for multi-choice response.
func PromptMultiChoice(message string, choices []string) ([]string, error) {
	return prompt.New().Ask(appendColon(message)).MultiChoose(
		choices,
		multichoose.WithHelp(true),
	)
}

// appendColon appends colon to the message if it does not end with punctuation.
func appendColon(message string) string {
	if !strings.HasSuffix(message, ".") || !strings.HasSuffix(message, "?") || !strings.HasSuffix(message, "!") || !strings.HasSuffix(message, ":") {
		message += ":"
	}
	return message
}
