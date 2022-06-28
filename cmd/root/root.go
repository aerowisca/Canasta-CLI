package cmd

import (
	"os"

	createCmd "github.com/CanastaWiki/Canasta-CLI-Go/cmd/create"
	deleteCmd "github.com/CanastaWiki/Canasta-CLI-Go/cmd/delete"
	importCmd "github.com/CanastaWiki/Canasta-CLI-Go/cmd/import"
	listCmd "github.com/CanastaWiki/Canasta-CLI-Go/cmd/list"
	startCmd "github.com/CanastaWiki/Canasta-CLI-Go/cmd/start"
	stopCmd "github.com/CanastaWiki/Canasta-CLI-Go/cmd/stop"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "canasta",
	Short: "A CLI tool for Canasta installations.",
	Long:  `A CLI tool to create, import, start, stop and backup multiple Canasta installations`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(createCmd.NewCmdCreate())
	rootCmd.AddCommand(importCmd.NewCmdCreate())
	rootCmd.AddCommand(startCmd.NewCmdCreate())
	rootCmd.AddCommand(stopCmd.NewCmdCreate())
	rootCmd.AddCommand(listCmd.NewCmdCreate())
	rootCmd.AddCommand(deleteCmd.NewCmdCreate())

}
