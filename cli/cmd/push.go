package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes the env file to the server",
	Long: `
	push is used to push the env file to the server.

	Example: 
	envdost push [File Name]
	pushesh the provided file to the server.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
