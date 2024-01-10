/*
Copyright © 2024 Squareofn anderzanovoleg@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"protocomp/cmd/config"
	"protocomp/cmd/generate"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "protocomp",
	Short: "Protocomp - easy to compile proto",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Logo)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generate.GenerateCmd)
}
