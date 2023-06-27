package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "deletes the project from the server",
	Long: `
	project sub-commandis used to delete the project from the server.

	Example: 
	envdost delete project [Project Name]
	deletes the provided project from the server.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete project called")
	},
}

func init() {

	// adding the project sub-command to the delete main command
	deleteCmd.AddCommand(projectCmd)
}
