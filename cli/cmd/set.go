package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	cmdRunner "github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Used to set a project",
	Long: `
	set is used to set the current project.
	
	Example: 
	envdost set [Project Name]
	Selects the provided project as the current project and all the other operations 
	are performed in this project.

	`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SetProject(args[0])
	},
}

func SetProject (projectName string) {
	// logins the user if not logged in
	SetupLogin()

	// starting the animation
	var wg sync.WaitGroup
	stopAnimation := make(chan struct{})
	wg.Add(1)
	go LoadingAnimation("Selecting "+ projectName +" ", stopAnimation, &wg)
	
	setProjectCmd := cmdRunner.NewCmd("op", "vault", "get", projectName, "--session", UserSession, "--format=json")
	status :=<- setProjectCmd.Start()
	data := strings.Join(status.Stdout, "")

	if status.Error != nil{ 
		fmt.Println( "Cannot find the project with the name ", projectName)
		close(stopAnimation)
		return 
	}
	if(status.Exit == 1){
		fmt.Println("More than one", projectName ,"or no",projectName,"found.\nSelecting project will improve in next version till then make sure you are typing correct project names and keep an eye at github.com/sarojregmi200/envdost\nfor a new version. :)")
		close(stopAnimation)
		return 
	}
	if status.Complete{
		close(stopAnimation)
		// setting the selected project to env
		envError:=SetEnv("SELECTED_PROJECT", data)
		if envError != nil{
			fmt.Println("Cannot set project", projectName)
			return 
		}
		fmt.Printf("\nProject %s is selected successfully\n", projectName);

		json.Unmarshal([]byte(data), &SelectedProject)
	}	
	wg.Wait()
}


func init() {
	RootCmd.AddCommand(setCmd)
}
