package main

import "fmt"

func main() {
	note_list := initData()
	for _, note := range *note_list {
		fmt.Println(note.path)
	}
}
