package cmd

import (
	"gotemplate/cfg"
	"gotemplate/cmd/http"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(cfg.RegistryConfig)
}

func RegistryCommand() {
	cmds := []*cobra.Command{
		//put all command to here
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
