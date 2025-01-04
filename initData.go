package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func initData() *noteList {
	src_path, err := os.Getwd()
	checkErr(err)
	obsidian_path, _ := path.Split(src_path)
	location_dir_list := initLocationDirList(obsidian_path)
	note_list := initNoteList(obsidian_path, location_dir_list)

	return note_list
}

func initLocationDirList(obsidian_path string) []string {
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

func initNoteList(obsidian_path string, location_dir_list []string) *noteList {
	var note_list noteList
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
	for _, n := range note_list {
		n.initRowYaml()
		n.initYamlData()
		n.initStatus()
	}
	return &note_list
}

func (a *yamlList) UnmarshalYAML(value *yaml.Node) error {
	var multi []string
	err := value.Decode(&multi)
	if err != nil {
		var single string
		err := value.Decode(&single)
		if err != nil {
			return err
		}
		if strings.Contains(single, ", ") {
			*a = strings.Split(single, ", ")
		} else {
			*a = []string{single}
		}
	} else {
		*a = multi
	}
	return nil
}

func (n *note) initRowYaml() {
	err := yaml.Unmarshal([]byte(n.rowYaml), &n.yamlData)
	if err != nil {
		fmt.Printf("ERROR at file %s\n", n.name)
		panic(err)
	}
}
func (n *note) initYamlData() {
	f, err := os.Open(n.path)
	checkErr(err)
	checkIfMarker := make([]byte, 4)
	n1, err := f.Read(checkIfMarker)
	checkErr(err)
	if string(checkIfMarker[:n1]) != "---\n" {
		fmt.Println("ERROR : No token at the start, cannot parse YAML")
	}

	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if string(line) == "---" {
			n.rowYaml = strings.TrimRight(n.rowYaml, "\n")
			break
		}

		n.rowYaml += string(line) + "\n"

		if err != nil {
			fmt.Println("ERROR : No token at the end, cannot parse YAML")
			// end of file
			break
		}
	}
}

func (n *note) initStatus() {
	for _, status := range n.yamlData.StatusList {
		switch status {
		case "DELETE":
			n.statusMap[DELETE] = true
		case "TEMP":
			n.statusMap[TEMP] = true
		default:
			n.unrecognizedStatus = append(n.unrecognizedStatus, status)
		}
	}
}
