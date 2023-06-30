package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/registry"
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

	// contains data from previous login
	var previousLogin string
	// checking the users os first
	if runtime.GOOS == "windows"{
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Failed to open registry key:", err)
	}
	defer key.Close()

	// getting the env variable
	loginToken, _,err := key.GetStringValue("LOGIN_TOKEN")
	sessionToken, _, err:= key.GetStringValue("USER_SESSION")
	if err != nil {
		fmt.Println("Failed to set environment variable:", err)
	}else{
		// setting the env variable to the previous login 
		previousLogin = loginToken
		UserSession = sessionToken

	}
	}else{
		previousLogin = os.Getenv("LOGIN_TOKEN")
		UserSession = os.Getenv("USER_SESSION")
	}

	if previousLogin != ""{ 
		json.Unmarshal([]byte(previousLogin), &LoggedInUser)
		LoggedIn = true
		return LoggedInUser // if user is already logged in no need to login again
	}  

	

	// if user is logged in return that
	if LoggedIn {
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
	UserSession = out.String()

	setLoggedInUser() // sets the LoggedInuser global variable

	return LoggedInUser
}


// sets the LoggedIn user details from the given session id 
func setLoggedInUser () {
	// gets the user info from the session token
	tokenCommand, err := exec.Command("op", "whoami","--session", UserSession, "--format=json").Output()
	if(err != nil){
		fmt.Println("Error while getting the user info")
		log.Fatal(err)	
	}

	// parsing the output 
	parsingErr :=	json.Unmarshal(tokenCommand, &LoggedInUser)
	if parsingErr != nil{
		fmt.Println(parsingErr)
	}


	// setting the loggedin user in the env
	if runtime.GOOS == "windows"{
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Failed to open registry key:", err)
	}
	defer key.Close()

	// getting the env variable
	tokenError := key.SetStringValue("LOGIN_TOKEN", string(tokenCommand[:]))
	sessionError := key.SetStringValue("USER_SESSION", UserSession)

	if tokenError != nil || sessionError != nil {
		fmt.Println("Failed to set environment variable:", tokenError, sessionError)
	}
	}else{

	// for unix
	tokenError := os.Setenv("LOGIN_TOKEN", string(tokenCommand[:]))
	sessionError := os.Setenv("USER_SESSION", UserSession)
	if tokenError != nil || sessionError != nil{
		fmt.Println("Session storage for terminal is turned off", tokenError)
	}
	}

	LoggedIn = true
}

// user types and definitions
type User struct {
	Shorthand 		string `json:"shorthand"`
	Url       		string	`json:"url"`
	Email     		string	`json:"email"`
	User_uuid   	string	`json:"user_uuid"`
	Account_uuid 	string	`json:"account_uuid"`
}

// LoggedIn user details
var LoggedInUser User

// contains the user session token
var UserSession  string = ""
var LoggedIn 	 bool = false // status of userlogin
 

// a boolean that controls the animation loop
var Animate 	 bool = false


func LoadingAnimation (txt string, pid int) {
	
	// handling the process and animation state accordingly
	process, err := os.FindProcess(pid)
	fmt.Println("process id", pid)
	if err != nil{
		// it means process is not found or invalid pid 
		Animate = false
		panic(err)
	} 
	// turning the animation on if the process exists
	Animate = true
	sequence  := [8] string {"⣾", "⣽", "⣻", "⢿" ,"⡿", "⣟", "⣯", "⣷"}  
	var counter int = 0 // moves till the array index
	for {
		// for looking the process completion
		err := process.Signal(syscall.Signal(0)) 
		// sending 0 as a signal will not disturb the process if it exists if it doesnot then we can look the error msg
		
		if Animate == false || err != nil  {
			Animate =  false
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

// for the selected projects
type Project struct {
	Name 		string `json:"name"`
	Id 			string `json:"id"`
}	

var SelectedProject Project