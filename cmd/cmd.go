package cmd

import (
	"gotemplate/cmd/http"

	"github.com/spf13/cobra"
)

func RegistryCommand() {
	cmds := []*cobra.Command{
		{
			Use:   "serve",
			Short: "Starting http server",
			Args:  cobra.MaximumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				http.RegistryHttpServer()
			},
		},
	}

	cmdEngine := cobra.Command{}
	cmdEngine.AddCommand(cmds...)

	cmdEngine.Execute()
}
