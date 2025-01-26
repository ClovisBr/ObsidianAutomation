/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"obsidian_automation/initData"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the name of files",
	Long:  `See it later`,
	Run: func(cmd *cobra.Command, args []string) {
		note_list := initData.InitData()
		for _, note := range *note_list {
			fmt.Println(note.Name)
		}
	},
}
