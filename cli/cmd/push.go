package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	cmdRunner "github.com/go-cmd/cmd"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes the env file to the server",
	Long: `
	push is used to push the env file to the server.

	Example: 
	envdost push [File Name]
	pushesh the provided file to the server.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// sets the requirements before calling the push 
		setup()

		for i:=0 ; i< len(args); i++{
			var filePath string = args[i]

			// actual name of the file
			var fileName string
			
			// due to unix and windows system.
			arr := strings.Split(filePath, "\\")
			if  len(arr) < 2{
				arr = strings.Split(fileName, "/")
			}
			fileName = arr[len(arr) - 1]

			processedData := processFile(filePath)
			uploadFile(processedData, filePath, fileName)
		}
	},
}


// processes each file and returns a string that
// in a way it can be stored in the op vault as a item
// item string with lines as feilds
func processFile (filePath string) []string{
	
	var itemString []string 
	
	// opening the file
	file, err := os.Open(filePath)
	if err != nil{
		fmt.Println("Cannot find the file ", filePath)
	} 
	defer file.Close()

	// creating a new scanner to read the file
	reader := bufio.NewScanner(file)

	// line number 
	i:= 0
	// reading file line by line
	for reader.Scan(){
		i = i+1
		line := reader.Text()
		// adding each line together for a item string
		itemString =append(itemString, lineParser(i, line)...)
	}
	return itemString
}

// parses individual line and returns a string which 
// represents each field of a item in op vault
func lineParser (lineNumber int , line string) []string {
	
	var parsedLine []string 
	if strings.TrimSpace(line) == ""{return parsedLine}

	keyValueFilter := strings.Split(strings.TrimSpace(line), "=") 
	if len(keyValueFilter) < 2{
		// it is a comment
		parsedLine = append(parsedLine , fmt.Sprintf(" %d. comment[text]=%s , ",lineNumber, line))
		return parsedLine
	}

	// left side of the = is key
	key := keyValueFilter[0]

	// unfiltered value i.e value which may contain comment
	// [1:] slices the array removing the first element
	rawValue := strings.Join(keyValueFilter[1:], "")
	
	valueCommentFilter := strings.Split(rawValue, "#")
	if len(valueCommentFilter) < 2 {
		parsedLine = append(parsedLine, fmt.Sprintf(" %d. %s[text]=%s ", lineNumber, key, rawValue ))
		return parsedLine
	}
	value := valueCommentFilter[0]
	
	comment := strings.Join(valueCommentFilter[1:], "")

	parsedLine =append(parsedLine, fmt.Sprintf(" %d. %s[text]=%s  ", lineNumber, key, value), fmt.Sprintf(" %d. comment[text]=%s",  lineNumber, comment) )

	return parsedLine
}

// calls the op and sets the vault 
// with the given file details
func uploadFile(data []string, filePath string, fileName string){
	// file location in disk
	diskLoc := fmt.Sprintf(` location[text]=%s `, filePath)


	var args [] string
	args = append(args, "item", "create", "--title", fileName, "--vault", SelectedProject.Id, "--session", UserSession, "--category", "Secure Note" , diskLoc)

	args = append(args, data...)

	addItemCmd := cmdRunner.NewCmd( "op",args...  );
 
	
 	status := <- addItemCmd.Start()
	fmt.Println(status)


	fmt.Println(data)
}

// checks and setsup the selected project if not panics
func setup (){
	if !LoggedIn {
		signinUser()
	}
	data, err := getEnv("SELECTED_PROJECT")
	if err != nil{
		log.Panic("Project not selected")
	}

	// parse the string data to struct
 	json.Unmarshal([]byte(data), &SelectedProject)
}
func init() {
	RootCmd.AddCommand(pushCmd)
}