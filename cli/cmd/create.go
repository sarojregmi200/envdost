package cmd

import (
	"fmt"

	"github.com/google/uuid"
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
		fmt.Print("Creating ")

		if len(args) > 1 {
			fmt.Println("Projects :")
		}else{
			fmt.Println("Project :")
		}
		for i := 0 ; i< len(args); i++{
			projectName := args[i]
			fmt.Println(projectName)
			createProject(args[i])
		}
	},
}

// creates a new op vault with the name of the project
// creates a new project in server with the name
func createProject(projectName string) {
	 projectId, error := uuid.NewUUID() 
	 handleError(error)

	 fmt.Println( projectName + " : " + projectId.String())
}


// accepts the error as param and panic if error occurs
func handleError(e error){
	if(e != nil){
		panic(e)
	 }
}


func init() {
	RootCmd.AddCommand(createCmd)
}
