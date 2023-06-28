package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Envdost",
	Short: "Envdost is a cli-application to manage your env",
	Long: `
	Env dost is a cli-application that lets you manage your env files,
	It helps you create a project and securely store the env files for those projects
	in the one password vault which can be easily updated, managed, and accessed 
	Whenever and where ever you are.
	`,
	
}


func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}




// used to signin the user if it is not signin
func signinUser () User{
	
	// if user is logged in return that
	if loggedIn {
		return LoggedInUser
	}	

	// one password singin command
	signinCmd := exec.Command("op", "signin", "-f", "--raw")
	signinCmd.Stdin = os.Stdin
	var out strings.Builder // creating a var out to store output
	signinCmd.Stdout = &out 
	signinCmd.Stderr = os.Stderr

	// handeling error while running command
	err := signinCmd.Run()
	if err != nil{
		fmt.Println("Error while signing up the user")
		log.Fatal(err)
	}
	userSession = out.String()

	setLoggedInUser() // sets the loggedinuser global variable

	return LoggedInUser
}


// sets the loggedin user details from the given session id 
func setLoggedInUser () {
	// gets the user info from the session token
	tokenCommand, err := exec.Command("op", "whoami","--session", userSession).Output()
	if(err != nil){
		fmt.Println("Error while getting the user info")
		log.Fatal(err)
		
	}

	// parsing the output
	var rawUserData string = string(tokenCommand[:])

	loggedIn = true
	lines := strings.Split(rawUserData, "\n") // extracting each line
	// formatting the string and extracting loggedin user details
	for _,v := range lines{
		data := strings.Split(v, ": ") 
		// checking if the data is valid or not

		// filtering the unwanted things
		if len(data) < 2{
			continue
		}
		// extracting the key value and formatting
		key := strings.ToLower(strings.ReplaceAll(data[0]," ", "")) 
		value := strings.TrimSpace(data[1])
		
		// setting the logged in user
		switch key{
		case "shorthand":
			LoggedInUser.shorthand = value
		case "url":
			LoggedInUser.url = value
		case "email":
			LoggedInUser.email = value
		case "userid":
			LoggedInUser.userid = value
		}
	}
}

// user types and definitions
type User struct {
	shorthand 	string
	url       	string
	email     	string
	userid    	string
}

// loggedin user details
var LoggedInUser User = User{
	shorthand:	 "",
	url:      	 "",
	email:   	 "",
	userid:      "",
}

// contains the user session token
var userSession  string = ""
var loggedIn 	 bool = false // status of userlogin