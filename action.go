package main

import "os"

func (n *note) delete() {
	err := os.Remove(n.path)
	checkErr(err)
	print(n.path)
}

func (n *note) pin() {
	println("not implemented yet")
}

func (n *note) archive() {
	println("not implemented yet")
}
