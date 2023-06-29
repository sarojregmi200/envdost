package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

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
	tokenCommand, err := exec.Command("op", "whoami","--session", userSession, "--format=json").Output()
	if(err != nil){
		fmt.Println("Error while getting the user info")
		log.Fatal(err)
		
	}

	// parsing the output 
	parsingErr :=	json.Unmarshal(tokenCommand, &LoggedInUser)
	
	if parsingErr != nil{
		fmt.Println(parsingErr)
	}
	loggedIn = true
}

// user types and definitions
type User struct {
	Shorthand 		string `json:"shorthand"`
	Url       		string	`json:"url"`
	Email     		string	`json:"email"`
	User_uuid   	string	`json:"user_uuid"`
	Account_uuid 	string	`json:"account_uuid"`
}

// loggedin user details
var LoggedInUser User

// contains the user session token
var userSession  string = ""
var loggedIn 	 bool = false // status of userlogin

// a boolean that controls the animation loop
var animate 	 bool = false


func loadingAnimation (txt string, pid int) {
	
	// handling the process and animation state accordingly
	process, err := os.FindProcess(pid)
	fmt.Println("process id", pid)
	if err != nil{
		// it means process is not found or invalid pid 
		animate = false
		panic(err)
	} 
	// turning the animation on if the process exists
	animate = true
	sequence  := [8] string {"⣾", "⣽", "⣻", "⢿" ,"⡿", "⣟", "⣯", "⣷"}  
	var counter int = 0 // moves till the array index
	for {
		// for looking the process completion
		err := process.Signal(syscall.Signal(0)) 
		// sending 0 as a signal will not disturb the process if it exists if it doesnot then we can look the error msg
		
		if animate == false || err != nil  {
			animate =  false
			break
		}

		fmt.Print( txt + " " + sequence[counter] + "\r")
		counter++
		if counter >= 7{
			counter = 0
		}
		time.Sleep(500 * time.Millisecond) // pauses the loop for 500 ms
		// to create a smooth animation
	}
} 

