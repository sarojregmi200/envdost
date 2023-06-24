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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
