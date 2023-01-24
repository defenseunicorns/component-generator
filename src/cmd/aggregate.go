/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	outputFileName string
)

// aggregateCmd represents the aggregate command
var aggregateCmd = &cobra.Command{
	Use:   "aggregate",
	Short: "aggregate a collection of component definition files and produce a single artifact",
	Long: `This command aggregates local or remote component-definition OSCAL yaml files.
	The purpose of creating a single concise artifact for platforms or other systems of aggregate software components.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(aggregateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aggregateCmd.PersistentFlags().String("foo", "", "A help for foo")
	aggregateCmd.Flags().StringVarP(&outputFileName, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aggregateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	fmt.Println("executing the aggregate run function")
}
