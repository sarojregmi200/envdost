/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the env file",
	Long: `
	update is used to update the remote env file with the provided file.
	
	Example: 
	envdost update
	This will take the current projects all env file.
	
	Example: 
	envdost update [File Name]
	This will update only the specified env file.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
