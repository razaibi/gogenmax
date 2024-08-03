package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/osteele/liquid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Config holds the filenames for data, template, and output
type Config struct {
	Items []Item `yaml:"items"`
}

// Item holds individual data, template, and output file paths
type Item struct {
	DataFile     string `yaml:"dataFile"`
	TemplateFile string `yaml:"templateFile"`
	OutputFile   string `yaml:"outputFile"`
}

func init() {
	rootCmd.AddCommand(runCmd)
}

// processCmd represents the process command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Process data and template files",
	Long:  `Reads a configuration YAML file, processes the specified data and template files, and generates output files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage: myapp process <config.yaml>")
		}
		configFile := args[0]

		config := readConfig(configFile)

		for _, item := range config.Items {
			data := readJSON(item.DataFile)
			templateContent := readFile(item.TemplateFile)

			// Parse the Liquid template
			output, err := liquid.NewEngine().ParseAndRenderString(templateContent, data)
			if err != nil {
				log.Fatalf("Failed to render template: %v", err)
			}

			// Write the output to the specified file
			os.WriteFile(item.OutputFile, []byte(output), 0644)
			if err != nil {
				log.Fatalf("Failed to write output file: %v", err)
			}

			fmt.Printf("Output generated successfully for %s!\n", item.OutputFile)
		}
	},
}

// readConfig reads and parses the YAML configuration file
func readConfig(filename string) Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}

// readJSON reads and parses the JSON data file
func readJSON(filename string) map[string]interface{} {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read data file: %v", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalf("Failed to parse data file: %v", err)
	}

	return data
}

// readFile reads the content of a file
func readFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	return string(content)
}
