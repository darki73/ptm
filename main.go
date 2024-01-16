package main

import (
	"github.com/darki73/ptm/cmd"
	"github.com/darki73/ptm/pkg/log"
)

// main is the entrypoint of the application.
func main() {
	if err := cmd.Execute(); err != nil {
		log.FatalfWithFields(
			"failed to execute command: %s",
			log.FieldsMap{
				"source": "main",
			},
			err.Error(),
		)
	}
}
