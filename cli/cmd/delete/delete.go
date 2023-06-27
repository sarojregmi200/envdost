package delete

import (
	"envdost/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes the env file or the project from the server",
	Long: `
	delete is used to delete the env file or the project from the server.

	Example: 
	envdost delete [project || env] [Project Name || File Name]
	deletes the provided file or project from the server.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
	},
}

func init() {

	// adding the delete command to the main root command
	cmd.RootCmd.AddCommand(deleteCmd)
}
