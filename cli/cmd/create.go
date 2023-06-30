package cmd

import (
	"fmt"

	cmdRunner "github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a project",
	Long: `
	create is used to create a project.

	Example:
	envdost create [Project Name]
	creates a project with the provided project name.

	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0 ; i< len(args); i++{
			projectName := args[i] 
			createProject(projectName)
		}
	},
}
 
// creates a new op vault with the name of the project
// creates a new project in server with the name
func createProject(projectName string) {
	if LoggedIn == false{
		signinUser() // will set the session id as well 
	} 
	// creating one password vault
	cmd:= cmdRunner.NewCmd("op", "vault", "create", projectName, "--session", UserSession)
	statusChannel := cmd.Start()

	status := <-statusChannel
	 
	if status.Error != nil{	
		fmt.Println("error")
	}
	
	// LoadingAnimation("Creating " + projectName + " project :", cmd.Process.Pid)
	 	 
}

func init() {
	RootCmd.AddCommand(createCmd)
}
