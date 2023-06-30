/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pulls the env file from the server",
	Long: `
	pull is used to pull the env file from the server.

	Example: 
	envdost pull [File Name]
	pullesh the provided file from the server.
	
	`,
	
	Run: func(cmd *cobra.Command, args []string) {
		// user login
		SetupLogin()
		
		// setting the selected project
		SetSelectedProject()

		// if no args it will be true
		var pullAll bool = false

		for i:=0; i< len(args); i++{
			pullFile(args[i], pullAll)
		}
		if len(args) < 1{
			pullFile("", true)
		}
	},
}


// fetches the data and creates a env file or files
func pullFile (fileName string, pullAll bool){
	fmt.Println("Pulling file ", fileName, "pull all = ", pullAll)
}

func init() {
	RootCmd.AddCommand(pullCmd)
}
