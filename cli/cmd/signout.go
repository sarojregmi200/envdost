/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

// signoutCmd represents the signout command
var signoutCmd = &cobra.Command{
	Use:   "signout",
	Short: "Used to signout.",
	Long: `
	signout will log you out of the current account, and reset the environment variable
	
	Example:
	envdost signout
	logs you out of the currently loggedin account
	`,
	Run: func(cmd *cobra.Command, args []string) {
		getPreviousLoginData()
		var wg sync.WaitGroup
		wg.Add(1)
		stopAnimation := make(chan struct{})

		if !LoggedIn || LoggedInUser.User_uuid == ""{
			fmt.Println("You are not logged in. Please login first to signout")
			return
		}
		go LoadingAnimation("Loggin out "+ LoggedInUser.Email, stopAnimation, &wg)

		loginToken := SetEnv("LOGIN_TOKEN", "")
		userSession:= SetEnv("USER_SESSION", "")

		if loginToken != nil || userSession != nil{
			fmt.Println("An error occured while signing out", LoggedInUser.Email)
		}
		fmt.Println("User", LoggedInUser.Email , "is logged out successfully.")
		close(stopAnimation)
		wg.Wait()
	},
}

func getPreviousLoginData (){
	// getting the previous user session
	previousUserSession, sessionError := getEnv("USER_SESSION")
	// if there are no error while getting the previous user session then setting it
	if sessionError == nil{
		UserSession = previousUserSession
	}
	// contains data from previous login
	previousLogin, loginTokenError := getEnv("LOGIN_TOKEN")
	if loginTokenError == nil {
		json.Unmarshal([]byte(previousLogin), &LoggedInUser)
		LoggedIn = true
	}
}

func init() {
	RootCmd.AddCommand(signoutCmd)
}
