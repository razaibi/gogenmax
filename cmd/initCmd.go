package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// initCmd represents the process command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize sample files",
	Long:  `Sets up workflow, data and template files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := "data.json"
		content := `
{
  "title": "Hello, World!",
  "description": "This is a sample description."
}`
		// Call the function to create the file with content
		err := createFileWithContent(fileName, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		fileName = "workflow.yaml"
		content = `
items:
  - dataFile: data.json
    templateFile: template.liquid
    outputFile: output1.txt		
		`
		// Call the function to create the file with content
		err = createFileWithContent(fileName, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		fileName = "template.liquid"
		content = `
<h1>{{ title }}</h1>
<p>{{ description }}</p>	
		`
		// Call the function to create the file with content
		err = createFileWithContent(fileName, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	},
}

func createFileWithContent(fileName string, content string) error {
	// Create or open the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
