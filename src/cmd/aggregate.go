package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/defenseunicorns/component-generator/src/internal/types"
	"github.com/defenseunicorns/component-generator/src/pkg/component"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const oscalVer = "1.0.4"

var (
	input   string
	name    string
	version string
	title   string
	stdout  bool
	remotes []string
	locals  []string
)

// aggregateCmd represents the aggregate command
var aggregateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "aggregate a collection of component definition files and produce a single artifact",
	Long: `This command aggregates local or remote component-definition OSCAL yaml files.
	The purpose of creating a single concise artifact for platforms or other systems of aggregate software components.
	`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(aggregateCmd)

	aggregateCmd.Flags().BoolVarP(&stdout, "stdout", "s", false, "print to stdout rather than the declaratively specified filename")
	aggregateCmd.Flags().StringVarP(&input, "input", "i", "", "Path to the file to be processed")
	aggregateCmd.Flags().StringVarP(&name, "name", "n", "", "Path/Name of the file to be created")
	aggregateCmd.Flags().StringVarP(&version, "file-version", "v", "", "the version of the document to be created")
	aggregateCmd.Flags().StringVarP(&title, "title", "t", "", "the title of the document to be created")
	aggregateCmd.Flags().StringArrayVarP(&locals, "local", "l", []string{}, "path to a local component file - component.yaml")
	aggregateCmd.Flags().StringArrayVarP(&remotes, "remote", "r", []string{}, "path to a remote component file - REPO_URI[.git]/PKG_PATH[@VERSION]")

}

func run() {
	var config types.ComponentsConfig
	path := input

	// If there is no input path specified for the declarative document
	// Then this must be an imperative run
	if path == "" {
		if version == "" {
			log.Fatal("Version is Required")
		}
		if title == "" {
			log.Fatal("Title is Required")
		}
		if name == "" {
			log.Fatal("Name is Required")
		}
		if len(remotes) == 0 && len(locals) == 0 {
			log.Fatal("Minimum 1 remote or local is Required")
		}

		config.Name = name
		config.Metadata.Version = version
		config.Metadata.Title = title
		config.Metadata.OscalVersion = oscalVer

		for _, v := range remotes {
			var remote types.Remote

			repoSplit := strings.Split(v, ".git")
			verSplit := strings.Split(repoSplit[1], "@")

			remote.Git = repoSplit[0] + ".git@" + verSplit[1]
			remote.Path = "." + verSplit[0]

			config.Components.Remotes = append(config.Components.Remotes, remote)
		}

		for _, v := range locals {
			var local types.Local
			local.Name = v

			config.Components.Locals = append(config.Components.Locals, local)
		}

	} else {
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			fmt.Printf("Path: %v does not exist - unable to digest document\n", path)
		}

		rawDoc, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(rawDoc, &config)
		if err != nil {
			log.Fatal(err)
		}
	}

	config.BaseDirectory, _ = filepath.Split(path)

	yamlDoc, oscalObj, err := component.BuildOscalDocument(config)
	if err != nil {
		log.Fatal(err)
	}
	_, error := os.Stat(config.Name)
	if error == nil {
		// if the file exists - read/unmarshall and compare
		fmt.Println("File exists - running comparison")
		var existingObj types.OscalComponentDocument
		rawExist, err := os.ReadFile(config.Name)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(rawExist, &existingObj)
		if err != nil {
			log.Fatal(err)
		}

		// Document now exists - compare

		unmodified := component.DiffComponentObjects(existingObj, oscalObj)

		if unmodified {
			// If not modified, no need to write new file
			fmt.Println("No fields have been updated - not updating document")
			if stdout {
				fmt.Print(string(yamlDoc))
			}
		} else {
			if !stdout {
				err := os.WriteFile(config.Name, []byte(yamlDoc), 0644)
				if err != nil {
					log.Fatalf("writing output: %s", err)
				}
			} else {
				fmt.Print(string(yamlDoc))
			}
		}

	} else {
		fmt.Println("File does not exist - running output")
		if !stdout {
			err := os.WriteFile(config.Name, []byte(yamlDoc), 0644)
			if err != nil {
				log.Fatalf("writing output: %s", err)
			}
		} else {
			fmt.Print(string(yamlDoc))
		}
	}

}
