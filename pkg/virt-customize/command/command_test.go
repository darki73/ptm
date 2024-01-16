package command

import (
	"github.com/darki73/ptm/pkg/configuration/repositories"
	"reflect"
	"testing"
)

var (
	// imagePathForTesting is the path to the image used for testing.
	imagePathForTesting = "/etc/ptm/images/ubuntu-22.04-cloudimage-amd64.img"
	// repositoryConfigurationForTesting is the configuration for the repository used for testing.
	repositoryConfigurationForTesting = &repositories.Configuration{
		Name:      "docker",
		GPG:       "https://download.docker.com/linux/ubuntu/gpg",
		URL:       "https://download.docker.com/linux/ubuntu",
		Release:   "jammy",
		Component: "stable",
		KeyName:   "docker-archive-keyring",
	}
)

// TestNewCommand tests the NewCommand function.
func TestNewCommand(t *testing.T) {
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(imagePathForTesting, "--example", arguments...)

	if cmd.GetCommand() != "--example" {
		t.Errorf("NewCommand did not set command correctly")
	}

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("NewCommand did not set arguments correctly")
	}
}

// TestGetOrder tests the GetOrder function.
func TestGetOrder(t *testing.T) {
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(imagePathForTesting, "--example", arguments...)

	if cmd.GetOrder() != 0 {
		t.Errorf("GetOrder did not return 0")
	}
}

// TestSetOrder tests the SetOrder function.
func TestSetOrder(t *testing.T) {
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(imagePathForTesting, "--example", arguments...)

	cmd.SetOrder(1)
	if cmd.GetOrder() != 1 {
		t.Errorf("SetOrder did not set order correctly")
	}
}

// TestGetCommand tests the GetCommand function.
func TestGetCommand(t *testing.T) {
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(imagePathForTesting, "--example", arguments...)

	if cmd.GetCommand() != "--example" {
		t.Errorf("GetCommand did not return --example")
	}
}

// TestGetArguments tests the GetArguments function.
func TestGetArguments(t *testing.T) {
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(imagePathForTesting, "--example", arguments...)

	if !reflect.DeepEqual(cmd.GetArguments(), []string{"arg1", "arg2"}) {
		t.Errorf("GetArguments did not return arg1 and arg2")
	}
}

// TestGetImage tests the GetImage function.
func TestGetImage(t *testing.T) {
	arguments := []interface{}{"arg1", "arg2"}
	cmd := NewCommand(imagePathForTesting, "--example", arguments...)

	if cmd.GetImage() != imagePathForTesting {
		t.Errorf("GetImage did not return %s", imagePathForTesting)
	}
}
