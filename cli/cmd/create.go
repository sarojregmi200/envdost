package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a project",
	Long: `
	create is used to create a project.

	Example:
	envdost create [Project Name]
	creates a project with the provided project name.

	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Creating ")

		if len(args) > 1 {
			fmt.Println("Projects :")
		}else{
			fmt.Println("Project :")
		}
		for i := 0 ; i< len(args); i++{
			fmt.Println(args[i])
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
