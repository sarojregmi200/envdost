package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
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

var ReferenceMode bool = false
func init() {
}

// used to signin the user if it is not signin
func signinUser () User{
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
	LoggedIn = true

	// setting the loggedin user in the env
	tokenError :=SetEnv("LOGIN_TOKEN", string(tokenCommand[:]))
	sessionError :=SetEnv("USER_SESSION", UserSession)
	if tokenError != nil || sessionError != nil{
		fmt.Println("Session storage for terminal is turned off")
	}
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

func LoadingAnimation (txt string, stopAnimation <- chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // closing the function

	fmt.Println("") // to avoid writing in previous line of other content
	sequence  := [8] string {"⣾", "⣽", "⣻", "⢿" ,"⡿", "⣟", "⣯", "⣷"}  
	var counter int = 0 // moves till the array index
	for {
		select{
		case <-stopAnimation:
			return
		default:
			fmt.Print( txt + " " + sequence[counter] + "\r")
			// fmt.Print( txt + " " + sequence[counter] +"\n" )
			counter++
			if counter >= 7{
				counter = 0
			}		
			// to create a smooth animation
			time.Sleep(100 * time.Millisecond) // pauses the loop for 500 ms
		}

		} 
} 

// for the selected projects
type Project struct {
	Name 		string `json:"name"`
	Id 			string `json:"id"`
}	

var SelectedProject Project

// fetches the value from the environment variable and returns the value and error
func getEnv(envVariable string) (string,error){
	// contains value from environment variable
	var envVariableValue string
	if runtime.GOOS == "windows"{
	// windows registery key
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	if err != nil {
		return "" , err
	}
	defer key.Close()

	// getting the env variable
	value, _ , fetchingError := key.GetStringValue(envVariable)
	if fetchingError != nil{
		return "" , fetchingError
	}
	// if the value exists but is empty
	if value == ""{
		return "", errors.New("No value found")
	}

	envVariableValue = value 
	return envVariableValue , nil  
}

// for linux and unix systems
	envVariableValue = os.Getenv(envVariable)

	if envVariableValue == ""{
		return "", errors.New("No value found")
	}
 
	return envVariableValue , nil  
}

// sets the provided value to the provided env variable
func SetEnv(envVariable string, envVariableValue string) error {
	// if windows
	if runtime.GOOS == "windows"{
	// windows registery key
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, registry.ALL_ACCESS)
	if err != nil {
		return errors.New("Error while setting the registery key")
	}
	defer key.Close()
	// setting the env variable
	key.SetStringValue(envVariable, envVariableValue)
	return nil
}

// for linux and unix systems
	os.Setenv(envVariable, envVariableValue)
	return nil
}


// prompts the user to login if not logged in
func SetupLogin(){
	if LoggedIn == false{
		signinUser() // will set the session id as well 
	} 
}

// looks for selected project and updates the selected project if not found panics
func SetSelectedProject() error{
	data, err := getEnv("SELECTED_PROJECT")
	if err != nil{
	return errors.New("Please set a project first!!")
	}
	
	// parse the string data to struct
 	json.Unmarshal([]byte(data), &SelectedProject)
	return nil
}