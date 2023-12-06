package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "deletes the env file  from the server",
	Long: `
	env sub-command is used to delete the env file from the server.

	Example: 
	envdost delete env [File Name]
	deletes the provided file or project from the server.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete env called")
	},
}

func init() {
	// adding the env subcommand to the delete main command
	deleteCmd.AddCommand(envCmd)
}
