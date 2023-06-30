package cmd

import (
	"encoding/json"
	"fmt"
	"os/exec"

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
		setProject(args[0])
	},
}

func setProject (projectName string) {
	
	// logins the user if not logged in
	SetupLogin()
	
	data, errSearching := exec.Command("op", "vault", "get", projectName, "--session", UserSession, "--format=json").Output()
	if errSearching != nil{ 
		fmt.Println( "Cannot find the project with the name ", projectName)
		return 
	}

	// setting the selected project to env
	envError:= setEnv("SELECTED_PROJECT", string(data[:]))
	
	if envError != nil{
		fmt.Println("Cannot set project", projectName)
		return 
	}
	fmt.Printf("\nProject %s is selected successfully\n", projectName);

	json.Unmarshal(data, &SelectedProject)
}


func init() {
	RootCmd.AddCommand(setCmd)
}
