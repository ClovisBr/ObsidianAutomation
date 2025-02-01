/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"obsidian_automation/appConfig"
	"obsidian_automation/initData"
	// "obsidian_automation/utils"
)

func main() {
	// note_list := initData()
	// for _, note := range *note_list {
	// 	fmt.Println(note.name)
	// }
	appConfig.InitConfig()

	obsidian_path := viper.GetString("ObsidianPath")
	fileInfo, err := os.Stat(obsidian_path)
	if err != nil {
		panic(err)
	}

	if fileInfo.IsDir() {
		note_list := initData.InitData(obsidian_path)
		for _, note := range *note_list {
			fmt.Println(note.Name)
		}
		fmt.Println("#####\nPrint: " + obsidian_path)
	} else {
		fmt.Println("ERR not a dir")
	}
}
