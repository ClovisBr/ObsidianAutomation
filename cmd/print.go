/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"obsidian_automation/initData"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"obsidian_automation/utils"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the name of files",
	Long:  `See it later`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		obsidian_path := strings.Join(args, " ")
		fileInfo, err := os.Stat(obsidian_path)
		utils.CheckErr(err)

		if fileInfo.IsDir() {
			note_list := initData.InitData(obsidian_path)
			for _, note := range *note_list {
				fmt.Println(note.Name)
			}
			fmt.Println("#####\nPrint: " + strings.Join(args, " "))
		} else {
			fmt.Println("ERR not a dir")
		}

	},
}
