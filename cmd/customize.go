package cmd

import (
	"github.com/darki73/ptm/pkg/customizer"
	"github.com/darki73/ptm/pkg/downloader"
	"github.com/spf13/cobra"
)

// customizeCommand represents the customize command.
var customizeCommand = &cobra.Command{
	Use:   "customize",
	Short: "Customizes the base image",
	Long:  "Performs the customization of the base image based on the provided configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		ensureRoot()
		if !ensurePackageAvailable("libguestfs-tools") {
			printAndErrorOut("libguestfs-tools is not installed. Please install it and try again.")
		}

		initializeConfiguration()
		configuration := getConfiguration()

		downloadClient, err := downloader.NewDownloader(configuration)
		if err != nil {
			printAndErrorOut(err.Error())
		}

		if err := downloadClient.Download(); err != nil {
			printAndErrorOut(err.Error())
		}

		handler := customizer.NewCustomizer(getConfiguration())
		if err := handler.Run(); err != nil {
			printAndErrorOut(err.Error())
		}
	},
}

// init initializes the customize command.
func init() {
	rootCmd.AddCommand(customizeCommand)
}
