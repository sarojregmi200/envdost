package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

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
	Animate = true
	// go LoadingAnimation("Selecting "+ projectName +" ")
	
	setProjectCmd := cmdRunner.NewCmd("op", "vault", "get", projectName, "--session", UserSession, "--format=json")
	status :=<- setProjectCmd.Start()
	data := strings.Join(status.Stdout, "")

	if status.Error != nil{ 
		fmt.Println( "Cannot find the project with the name ", projectName)
		return 
	}
	if(status.Exit == 1){
		fmt.Println("More than one", projectName + " found.\nSelecting project with multiple same name is not available in this version.\nIt will be available in next version, till then try making project name different and keep an eye at github.com/sarojregmi200/envdost\nfor a new version. :)")
		return 
	}
	if status.Complete{
		Animate = false
		// setting the selected project to env
		envError:=SetEnv("SELECTED_PROJECT", data)
		if envError != nil{
			fmt.Println("Cannot set project", projectName)
			return 
		}
		fmt.Printf("\nProject %s is selected successfully\n", projectName);

		json.Unmarshal([]byte(data), &SelectedProject)
	}	
}


func init() {
	RootCmd.AddCommand(setCmd)
}
