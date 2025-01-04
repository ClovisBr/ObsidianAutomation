package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type yamlData struct {
	Uuid       string      `yaml:"uuid"`
	TagList    StringArray `yaml:"tags"`
	StatusList StringArray `yaml:"status"`
}

type StringArray []string

func (a *StringArray) UnmarshalYAML(value *yaml.Node) error {
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

type Data struct {
	Field StringArray
}

func parseYaml(ptr_note_list *[]*note) {
	for _, file := range *ptr_note_list {
		err := yaml.Unmarshal([]byte(file.rowYaml), &file.yamlData)
		if err != nil {
			fmt.Printf("ERROR at file %s\n", file.name)
			panic(err)
		}
	}
}
func getYaml(ptr_note_list *[]*note) {
	for _, file := range *ptr_note_list {
		f, err := os.Open(file.path)
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
				file.rowYaml = strings.TrimRight(file.rowYaml, "\n")
				break
			}

			file.rowYaml += string(line) + "\n"

			if err != nil {
				fmt.Println("ERROR : No token at the end, cannot parse YAML")
				// end of file
				break
			}
		}
	}
}
