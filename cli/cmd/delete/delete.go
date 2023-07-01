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
		fmt.Println("It is used to delete projects or env files, use -h to learn about it along with  examples")
	},
}

func init() {

	// adding the delete command to the main Root command
	cmd.RootCmd.AddCommand(deleteCmd)
}
