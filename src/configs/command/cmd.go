package command

import (
	"github.com/depri11/vehicle/src/configs/serve"
	"github.com/depri11/vehicle/src/database"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(serve.ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
