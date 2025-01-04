package main

import "os"

type note struct {
	name               string
	path               string
	statusMap          map[status]bool
	unrecognizedStatus []string
	rowYaml            string
	yamlData           yamlData
}

type status int

const (
	DELETE status = iota
	TEMP
)

var statusName = map[status]string{
	DELETE: "DELETE",
	TEMP:   "TEMP",
}

func (s status) String() string {
	return statusName[s]
}

// func (n *note) getStatus() {
// 	for _, status := range n.yamlData.StatusList {
// 		switch status {
// 		case "DELETE":
// 			note.statusMap = true
// 		case "TEMP":
// 			note.status[TEMP] = true
// 		}
// 	}
// }

func (n *note) delete() {
	err := os.Remove(n.path)
	checkErr(err)
	print(n.path)
}
