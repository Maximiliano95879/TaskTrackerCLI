package main

import (
	"fmt"
)

// CLI arguments
const (
	add            string = "add"
	update         string = "update"
	delete         string = "delete"
	list           string = "list"
	markInProgress string = "mark-in-progress"
	markDone       string = "mark-done"
)

// task statuses
const (
	taskDone       string = "done"
	taskInProgress string = "todo"
	taskNotDone    string = "in-progress"
)

func main() {
	fmt.Println("Hello World")

	//Infinite loop so that I can see output
	for {

	}
}
