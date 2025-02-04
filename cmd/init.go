/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path"

	"github.com/gulegulzartechnologies/gext/pkg/filesystem"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("initializing gext")

		githubFolderPath := path.Join(os.Getenv("HOME"), "GITHUB")

		if err := filesystem.CheckFolder(githubFolderPath); err != nil {
			log.Println("folder not found, initializing folder...")
			log.Println(githubFolderPath)
			filesystem.CreateFolder(githubFolderPath)
			log.Println("folder created")
		} else {
			log.Println("folder exists")
		}

		log.Println("gext initialized")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
