package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Used to set a project",
	Long: `
	set is used to set the current project.
	
	Example: 
	envdost set [Project Name]
	Selects the provided project as the current project and all the other operations 
	are performed in this project.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
