package main

import (
	"fmt"
	"os"
)

type Command int
type TaskStatus int

// CLI commands
const (
	add Command = iota
	update
	delete
	list
	markInProgress
	markDone
)

// task statuses
const (
	taskDone TaskStatus = iota
	taskInProgress
	taskNotDone
)

type cliInfo struct {
	action   Command
	id       int
	status   TaskStatus
	taskName string
}

var commandEnum = map[string]Command{
	"add":              add,
	"update":           update,
	"delete":           delete,
	"mark-in-progress": markInProgress,
	"mark-done":        markDone,
	"list":             list,
}

var taskEnum = map[string]TaskStatus{
	"done":        taskDone,
	"todo":        taskInProgress,
	"in-progress": taskNotDone,
}

func main() {

	if info, err := parseCliArgs(os.Args[1:]); err == nil {
		fmt.Printf("%v\n", info)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
