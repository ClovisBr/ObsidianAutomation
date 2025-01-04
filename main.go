package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	src_path, err := os.Getwd()
	checkErr(err)
	obsidian_path, _ := path.Split(src_path)
	location_dir_list := scanLocationDir(obsidian_path)
	note_list := scanDir(obsidian_path, location_dir_list)
	getYaml(note_list)
	parseYaml(note_list)

	for _, note := range *note_list {
		fmt.Println(note.path)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
