package main

import (
	"io/fs"
	"os"
	"regexp"
)

func scanLocationDir(obsidian_path string) []string {
	obsdian_dir_list, err := os.ReadDir(obsidian_path)
	checkErr(err)
	var location_dir_list []string
	for _, file := range obsdian_dir_list {
		r, err := regexp.Compile("^[0-9] - [A-Z]")
		checkErr(err)
		if r.MatchString(file.Name()) == true && file.IsDir() {
			location_dir_list = append(location_dir_list, file.Name())
		}
	}
	return location_dir_list
}

func scanDir(obsidian_path string, location_dir_list []string) *[]*note {
	var note_list []*note
	os.Chdir(obsidian_path)
	for _, dir := range location_dir_list {
		fileSystem := os.DirFS(obsidian_path + dir)

		fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
			checkErr(err)
			// d_info, err:= d.Info()
			// if !d.IsDir() && d_info.ModTime() > previous_runtime {
			if !d.IsDir() {
				checkErr(err)
				var file note
				file.name = d.Name()
				file.path = obsidian_path + dir + "/" + path
				note_list = append(note_list, &file)
			}
			return nil
		})
	}
	return &note_list
}
