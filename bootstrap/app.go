package bootstrap

import (
	"core-gin/commands"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "core-gin",
	Short: "Commander for core gin",
	Long: `
		This is a command runner or cli for api architecture in golang.
		Using this we can use underlying dependency injection container for running scripts.
		Main advantage is that, we can use same services, repositories, infrastructure present in the application itself`,
	TraverseChildren: true,
}

// App root of the application
type App struct {
	*cobra.Command
}

// NewApp creates new root command
func NewApp() App {
	cmd := App{
		Command: RootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
