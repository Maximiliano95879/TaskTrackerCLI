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
	noStatus TaskStatus = iota
	taskDone
	taskInProgress
	taskNotDone
)

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

	var info cliInfo
	var err error
	if info, err = parseCliArgs(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch info.Action {
	case add:
		if id, err := addTask(info); err == nil {
			fmt.Printf("Task added successfully (ID: %d)", id)
		} else {
			fmt.Printf("Task failed to be added: %v", err)
		}
	case update:
		updateTask(info)
	case delete:
		deleteTask(info)
	case list:
		listTasks(info)
	case markInProgress:
		markTaskInProgress(info)
	case markDone:
		markTaskDone(info)
	}
}
