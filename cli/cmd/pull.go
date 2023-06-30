/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
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
		SetSelectedProject()

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
	// getting all the files
	getAllFilesCommand := cmdRunner.NewCmd("op",  "items" ,"list", "--vault", SelectedProject.Id, "--session", UserSession, "--format=json")

	response :=<- getAllFilesCommand.Start()
	data := response.Stdout

	json.Unmarshal([]byte(strings.Join(data, "")), &files)
	
	// checking if the asked files exists or not
	if !pullAll {
		var fileExists bool = false
		var existingFileNames [] string 
		for i := 0; i< len(files); i++{
			currentFileName := files[i].Name
			// adding to the existing file name
			existingFileNames = append(existingFileNames, currentFileName)
			if currentFileName == fileName{
				fileExists = true
				break
			}
		} 
		if !fileExists{
			// contains all the existing file names separated by ,
			fileNamesStr := strings.Join(existingFileNames, ",")
			fmt.Printf("\nThere is no file with the name %s\n", fileName)
			fmt.Printf("Is you file name listed?  \n%s",fileNamesStr)
		}
	}
}

func init() {
	RootCmd.AddCommand(pullCmd)
}

type File struct{
	Name		string `json:"title"`
	Id			string `json:"id"`
	Category	string `json:"category"`
}

var files[] File



// references

// to get the files in a project
	// op items list --vault vaultId i.e projectId --session sessionid --format=json
// response
// [
//   {
//     "id": "4zhjg2cuhp5psuzqypvtxfxbq4",
//     "title": ".env",
//     "version": 1,
//     "vault": {
//       "id": "toyvtochukuoekmwfd6rp65j3i",
//       "name": "validTest"
//     },
//     "category": "SECURE_NOTE",
//     "last_edited_by": "B5ZLBJDUZJD6ZNGBUNCDMESOTM",
//     "created_at": "2023-06-30T12:16:27Z",
//     "updated_at": "2023-06-30T12:16:27Z"
//   }
// ]

// get the individual file content
// op items get fileId --vault toyvtochukuoekmwfd6rp65j3i --session 3FRT02DIG5kol2imZBlFyZ9RbZ8vzHZF_jMdN8trB8E --format=json
// {
// 	"id": "4zhjg2cuhp5psuzqypvtxfxbq4",
// 	"title": ".env",
// 	"version": 1,
// 	"vault": {
// 	  "id": "toyvtochukuoekmwfd6rp65j3i",
// 	  "name": "validTest"
// 	},
// 	"category": "SECURE_NOTE",
// 	"last_edited_by": "B5ZLBJDUZJD6ZNGBUNCDMESOTM",
// 	"created_at": "2023-06-30T12:16:27Z",
// 	"updated_at": "2023-06-30T12:16:27Z",
// 	"sections": [
// 	  {
// 		"id": "Section_7m5uzc3d2vjlzh2ririqkokp24"
// 	  },
// 	  {
// 		"id": "Section_bvopeayqfsiykauankvrdoyvxe",
// 		"label": " 1"
// 	  },
// 	  {
// 		"id": "Section_xn4b56au3uuk2y6ilwovgkdqhm",
// 		"label": " 2"
// 	  },
// 	  {
// 		"id": "Section_i2hryaih6rtdubkyvgy3acl3jq",
// 		"label": " 3"
// 	  }
// 	],
// 	"fields": [
// 	  {
// 		"id": "notesPlain",
// 		"type": "STRING",
// 		"purpose": "NOTES",
// 		"label": "notesPlain",
// 		"reference": "op://validTest/.env/notesPlain"
// 	  },
// 	  {
// 		"id": "k7nh4z4gdgbyhkcgj5lalhbe4i",
// 		"section": {
// 		  "id": "Section_7m5uzc3d2vjlzh2ririqkokp24"
// 		},
// 		"type": "STRING",
// 		"label": " location",
// 		"value": "..\\api\\config\\.env ",
// 		"reference": "op://validTest/.env/Section_7m5uzc3d2vjlzh2ririqkokp24/ location"
// 	  },
// 	  {
// 		"id": "c5qxepwmqs7kq2g5nw6nblpof4",
// 		"section": {
// 		  "id": "Section_bvopeayqfsiykauankvrdoyvxe",
// 		  "label": " 1"
// 		},
// 		"type": "STRING",
// 		"label": " comment",
// 		"value": "# connection string , ",
// 		"reference": "op://validTest/.env/ 1/ comment"
// 	  },
// 	  {
// 		"id": "4e2sf5w7vb7mskfv2nxooru2py",
// 		"section": {
// 		  "id": "Section_xn4b56au3uuk2y6ilwovgkdqhm",
// 		  "label": " 2"
// 		},
// 		"type": "STRING",
// 		"label": " DBCONNECTIONSTRING",
// 		"value": "dburl  ",
// 		"reference": "op://validTest/.env/ 2/ DBCONNECTIONSTRING"
// 	  },
// 	  {
// 		"id": "u6ujrqymporzi7tcd36aac5tdi",
// 		"section": {
// 		  "id": "Section_xn4b56au3uuk2y6ilwovgkdqhm",
// 		  "label": " 2"
// 		},
// 		"type": "STRING",
// 		"label": " comment",
// 		"value": "url is the test one",
// 		"reference": "op://validTest/.env/ 2/ comment"
// 	  },
// 	  {
// 		"id": "lnem4x54fouj6lyrcooxkruqxm",
// 		"section": {
// 		  "id": "Section_i2hryaih6rtdubkyvgy3acl3jq",
// 		  "label": " 3"
// 		},
// 		"type": "STRING",
// 		"label": " PORT",
// 		"value": "3000  ",
// 		"reference": "op://validTest/.env/ 3/ PORT"
// 	  },
// 	  {
// 		"id": "aezhkkxo7ucoqsc3hthafp6mwi",
// 		"section": {
// 		  "id": "Section_i2hryaih6rtdubkyvgy3acl3jq",
// 		  "label": " 3"
// 		},
// 		"type": "STRING",
// 		"label": " comment",
// 		"value": " this is dope though",
// 		"reference": "op://validTest/.env/ 3/ comment"
// 	  }
// 	]
//   }
