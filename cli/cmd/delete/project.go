package delete

import (
	"encoding/json"
	root "envdost/cmd"
	"fmt"
	"strings"

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
	deletes all the projects that match with the provided project name.

	envdost delete project 
	deletes the selected project.
	
	`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// user login
		root.SetupLogin()

		// setting up the selected project
		selectingError := root.SetSelectedProject()

		if selectingError !=nil{
		fmt.Println(selectingError)
		return
		}
		// creating a project array and then passing to delete project
		var projects [] root.Project
		if len(args) > 0{
			projects = findProject(args[0]) 
		}else{
			projects = append(projects, root.SelectedProject) 
		}
		
		deleteProject(projects)
	},
}


func deleteProject (projects []root.Project){
	for _, project := range projects{
		if project.Name == "" || project.Id == ""{
			continue
		}
		// root.Animate = true
		// go root.LoadingAnimation("Deleting project"+ project.Name)
		
		deleteProCmd := cmdRunner.NewCmd("op", "vault", "delete", project.Id ,"--session", root.UserSession)
		status :=<- deleteProCmd.Start()
		
		if status.Error != nil{
			fmt.Println("\nSry, cannot delete project", project.Name)
		}
		
		if status.Complete {
			// root.Animate = false
			fmt.Println("\nProject "+project.Name+" deleted successfully!!")
			// removing the project from env 
			envError := root.SetEnv("SELECTED_PROJECT", "")
			if envError != nil{
				fmt.Println("Error while removing selected project from the session, please select another project with set command before using anything else")
			}
			continue
		}
	}
}
func findProject(projectName string) []root.Project{
	var allProjects []root.Project 
	var projects [] root.Project
	// animation
	// root.Animate = true
	// go root.LoadingAnimation("Searching project "+ projectName)
	
	// finding the project
	findProjectCmd := cmdRunner.NewCmd("op", "vault", "list", "--session", root.UserSession, "--format=json")
	status :=<- findProjectCmd.Start()
	data := strings.Join(status.Stdout, "")

	if status.Error != nil{ 
		fmt.Println( "Cannot find the project with the name ", projectName)
		return allProjects
	}
	if(status.Exit == 1){
		fmt.Println("No project named", projectName + " found ")
		return allProjects
	}
	if status.Complete{
		root.Animate = false
		json.Unmarshal([]byte(data), &allProjects)
		
		for _,project := range allProjects{
			if(projectName == strings.TrimSpace(project.Name)){
				projects = append(projects, project) // adding to the project list
			}
		}

		if len(projects) > 0{
			fmt.Printf("\nProject %s is found %d times \n", projectName, len(projects));
		}else{
			fmt.Printf("\nProject %s is not found \n", projectName);
		}
	}

	return projects
}


func init() {
	// adding the project sub-command to the delete main command
	deleteCmd.AddCommand(projectCmd)
}
