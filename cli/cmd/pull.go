/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	cmdRunner "github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pulls the env file from the server",
	Long: `
	pull is used to pull the env file from the server.

	Example: 
	envdost pull [File Name]
	pullesh the provided file from the server.
	
	`,
	
	Run: func(cmd *cobra.Command, args []string) {
		// user login
		SetupLogin()
		
		// setting the selected project
		err := SetSelectedProject()
		if err!=nil{
			fmt.Println(err)
			return 
		}
		// if no args it will be true
		var pullAll bool = false

		for i:=0; i< len(args); i++{
			pullFile(args[i], pullAll)
		}
		if len(args) < 1{
			pullFile("", true)
		}
	},
}


// fetches the data and creates a env file or files
func pullFile (fileName string, pullAll bool){

	// animation
	Animate = true
	if pullAll {
		go LoadingAnimation("Pulling all config files from project "+ SelectedProject.Name)
	}else{
		go LoadingAnimation("Pulling "+ fileName +" from project "+ SelectedProject.Name)
	}

	// getting all the files
	getAllFilesCommand := cmdRunner.NewCmd("op",  "items" ,"list", "--vault", SelectedProject.Id, "--session", UserSession, "--format=json")

	response :=<- getAllFilesCommand.Start()
	data := response.Stdout

	
	json.Unmarshal([]byte(strings.Join(data, "")), &files) 
	
	if len(files) < 1 && pullAll{
		Animate = false
		fmt.Println("\nNo files found")
		return  
	}

	// contains the matched file
	var currentFile File 
	// checking if the asked files exists or not
	if !pullAll {
		var fileExists bool = false
		var existingFileNames [] string 
		for i := 0; i< len(files); i++{
			currentFileName := files[i].Name
			// adding to the existing file name
			existingFileNames = append(existingFileNames, currentFileName)
			if currentFileName == fileName{
				// setting the current file to extract it's content
				currentFile = files[i]
				fileExists = true
				break
			}
		} 
		if !fileExists{
			// contains all the existing file names separated by ,
			fileNamesStr := strings.Join(existingFileNames, ",")
			fmt.Printf("\nThere is no file with the name %s\n", fileName)
			fmt.Printf("Is you file name listed?  \n%s",fileNamesStr)
			return 
		}
	}

	// accounting for the pull all mode
	if pullAll{
		for i:=0; i< len(files); i++{

			createFile(files[i])
		}
		return
	}

	createFile(currentFile)
	return
}

func init() {
	pullCmd.Flags().BoolVarP(&ReferenceMode,"refmode" ,"r", false, "pulls env file in reference mode.")
	RootCmd.AddCommand(pullCmd)
}

type File struct{
	Name		string `json:"title"`
	Id			string `json:"id"`
	Category	string `json:"category"`
}

var files[] File

// fetches the file content and creates the file
func createFile (currentFile File){

	fetchDataCommand := cmdRunner.NewCmd("op","item", "get", currentFile.Id, "--vault", SelectedProject.Id, "--session", UserSession , "--format=json" )

	response :=<- fetchDataCommand.Start()
	data := strings.Join(response.Stdout, "")

	var doc Lines
	json.Unmarshal([]byte(data), &doc)
	

	var fileData string
	var previousLineNumber int = 0
	var fileLocation string
	var newFile os.File 
	for i:=0; i< len(doc.Fields); i++{
		currentLine := doc.Fields[i]

		// for creating the file in the stored location
		if strings.TrimSpace(currentLine.Label) == "location"{
			// creating the file in the stored location
			fileLocation = currentLine.Value
			file, err :=os.Create(fileLocation)
			if err != nil {
				fmt.Println("Error creating file:", currentFile.Name)
				return
			}
			defer file.Close()
			newFile = *file
			fmt.Printf("File %s created successfully.", currentFile.Name)
		}

		// removing the items without the line number
		if currentLine.Sec.LineNumber == ""{
			continue
		}

		// valid lines reach here 
		// handeling the comments
		if strings.TrimSpace(currentLine.Label) == "comment"{
			fileData += "#"+ strings.TrimSpace(currentLine.Value)
		}else{ 
			// handling the non comment lines
			// checking for the ref mode
			if ReferenceMode{
				fileData += strings.TrimSpace(currentLine.Label) +"='"+ strings.TrimSpace(currentLine.Reference)+"'"
				}else{
					fileData += strings.TrimSpace(currentLine.Label) +"="+ strings.TrimSpace(currentLine.Value)
				}
		}
		// looking for place where line shouldnot break
		if string(previousLineNumber) == strings.TrimSpace(currentLine.Sec.LineNumber){
			continue 
		}
		// breaking the line
		fileData += "\n"
	}

	//  writing to the generated file
	_, writingError := newFile.Write([]byte(fileData))
	if writingError !=nil{
		fmt.Println("Error while writing to the file", currentFile.Name, writingError)
	}
}


type Lines struct {
	Fields []Line `json:"fields"`
}
type Line struct {
	Id 			string 		`json:"id"`
	Sec 		Section 	`json:"section"`  
	Type 		string 		`json:"type"`
	Label 		string 		`json:"label"`
	Value 		string 		`json:"value"`
	Reference 	string 		`json:"reference"`
}

type Section struct{
	Id 				string 	`json:"id"`
	LineNumber  	string 	`json:"label"`
}
