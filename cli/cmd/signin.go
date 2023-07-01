/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

// signinCmd represents the signin command
var signinCmd = &cobra.Command{
	Use:   "signin",
	Short: "Used to signin to you account",
	Long: `
	Signin lets you login to you account via password or any other method such as fingerprint or faceid
	Note: for fingerprint or faceid login make sure to turn on onepassword cli from the application.
	If you are unable to use the faceid or fingerprint you can always login with your password

	Example:
	envdost signin
	prompts you to verify yourself via password, fingerprint or faceid
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var wgNew sync.WaitGroup
		stopAnimation := make(chan struct{})
		wgNew.Add(1)
		
		go LoadingAnimation("Checking previous logins", stopAnimation, &wgNew)
		getPreviousLoginData()
		if LoggedIn{
			fmt.Println("\nYou are already loggedin as", LoggedInUser.Email)
			close(stopAnimation)
			return
		}
		close(stopAnimation)
		wgNew.Wait()
		
		// main signin function
		SetupLogin()
		
		fmt.Println("\nSuccessfully signin as:", LoggedInUser.Email)
	},
}

func init() {
	RootCmd.AddCommand(signinCmd)

	
}

