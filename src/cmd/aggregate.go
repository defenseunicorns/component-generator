/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/brandtkeller/component-generator/src/internal/types"
	"github.com/brandtkeller/component-generator/src/pkg/component"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	version string
	stdout  bool
)

// aggregateCmd represents the aggregate command
var aggregateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "aggregate a collection of component definition files and produce a single artifact",
	Long: `This command aggregates local or remote component-definition OSCAL yaml files.
	The purpose of creating a single concise artifact for platforms or other systems of aggregate software components.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func init() {
	rootCmd.AddCommand(aggregateCmd)

	aggregateCmd.Flags().BoolVarP(&stdout, "stdout", "s", false, "print to stdout rather than the declaratively specified filename")
	aggregateCmd.Flags().StringVarP(&version, "file-version", "v", "", "the version of the document to be created)")

}

func run(commandArgs []string) {
	var config types.ComponentsConfig
	path := commandArgs[0]

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
		return
	}

	yamlDoc, err := component.BuildOscalDocument(config)
	if err != nil {
		log.Fatal(err)
	}

	if !stdout {
		err := os.WriteFile(config.Name, []byte(yamlDoc), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(string(yamlDoc))
	}

}
