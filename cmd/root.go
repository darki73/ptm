package cmd

import (
	"fmt"
	config "github.com/darki73/ptm/pkg/configuration"
	"github.com/darki73/ptm/pkg/log"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var (
	// configurationPath is the path to the configuration file.
	configurationPath string
	// configurationName is the name of the configuration file.
	configurationName string
	// configurationExtension is the extension of the configuration file.
	configurationExtension string
)

// rootCmd is the root command.
var rootCmd = &cobra.Command{
	Use:   "ptm",
	Short: "Proxmox Templates Maker",
	Long:  "This application allows you to customize images and create templates for Proxmox VE.",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// init initializes the root command.
func init() {
	rootCmd.PersistentFlags().StringVar(&configurationPath, "configuration-path", "/etc/ptm", "Path to the configuration file")
	rootCmd.PersistentFlags().StringVar(&configurationName, "configuration-name", "config", "Name of the configuration file")
	rootCmd.PersistentFlags().StringVar(&configurationExtension, "configuration-extension", "yaml", "Extension of the configuration file")
}

// getConfigurationOptions returns the configuration options.
func getConfigurationOptions() *config.Options {
	return &config.Options{
		Path:      configurationPath,
		Name:      configurationName,
		Extension: configurationExtension,
	}
}

// initializeConfiguration initializes the configuration.
func initializeConfiguration() {
	if err := config.LoadConfiguration(getConfigurationOptions()); err != nil {
		log.Fatalf("failed to initialize configuration: %s", err)
	}
}

// getConfiguration returns the configuration.
func getConfiguration() *config.Configuration {
	if !config.IsConfigurationLoaded() {
		initializeConfiguration()
	}

	return config.GetConfiguration()
}

// printAndExit prints the error and exits with the given code.
func printAndExit(err string, code int) {
	fmt.Println(err)
	os.Exit(code)
}

// printAndErrorOut prints the error and exits with code 1.
func printAndErrorOut(err string) {
	printAndExit(err, 1)
}

// ensureRoot ensures that the user is root.
func ensureRoot() {
	currentUser, err := user.Current()
	if err != nil {
		printAndErrorOut(fmt.Sprintf("failed to get current user: %s", err))
	}

	uid := currentUser.Uid
	if uid != "0" {
		printAndErrorOut("you need to be root to run this command")
	}
}

// ensurePackageAvailable ensures that the given package is available.
func ensurePackageAvailable(packageName string) bool {
	cmd := exec.Command("dpkg", "-l")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running dpkg command:", err)
		return false
	}

	return strings.Contains(string(output), packageName)
}
