/*
Copyright © 2023 Saroj Regmi sarojregmi.official@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Envdost",
	Short: "Envdost is a cli-application to manage your env",
	Long: `Env dost is a cli-application that lets you manage your env files,
	It helps you create a project and securely store the env files for those projects
	in the one password vault which can be easily updated, managed, and accessed 
	Whenever and where ever you are.`,
	
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


