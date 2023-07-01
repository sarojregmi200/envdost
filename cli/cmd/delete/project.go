package delete

import (
	root "envdost/cmd"
	"fmt"

	cmdRunner "github.com/go-cmd/cmd"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "deletes the project",
	Long: `
	project sub-commandis used to delete the project.

	Example: 
	envdost delete project [Project Name]
	deletes the provided project from the server.

	envdost delete project 
	deletes the selected project.
	
	`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// user login
		root.SetupLogin()

		// setting up the selected project
		root.SetSelectedProject()
		
		if len(args) > 0{
			deleteProject(findProject(args[0]))
		}
		deleteProject(root.SelectedProject)
	},
}


func deleteProject (project root.Project){
	root.Animate = true
	go root.LoadingAnimation("Deleting project"+ project.Name)
	
	deleteProCmd := cmdRunner.NewCmd("op", "vault", "delete", project.Id ,"--session", root.UserSession)
	status :=<- deleteProCmd.Start()

	if status.Error != nil{
		fmt.Println("\nSry, cannot delete project", project.Name)
	}

	if status.Complete {
		root.Animate = false
		fmt.Println("\nProject "+project.Name+" deleted successfully!!")
		// removing the project from env 
		envError := root.SetEnv("SELECTED_PROJECT", "")
		if envError != nil{
			fmt.Println("Error while removing selected project from the session, please select another project with set command before using anything else")
		}
		return
	}
}
func findProject(projectName string) root.Project{
	var project root.Project 
	return project
}


func init() {
	// adding the project sub-command to the delete main command
	deleteCmd.AddCommand(projectCmd)
}
