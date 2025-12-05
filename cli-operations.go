package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type cliInfo struct {
	Action   Command
	Id       string
	Status   TaskStatus
	TaskName string
}

type task struct {
	Id          string
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func addTask(info cliInfo) (string, error) {
	id := uuid.New().String()
	timeStamp := time.Now().Format(time.Stamp)
	encoding, jsonErr := json.Marshal(task{Id: id, Description: info.TaskName, CreatedAt: timeStamp})
	if jsonErr == nil {
		fileName := fmt.Sprintf("%s.json", info.TaskName)
		if file, fileErr := os.Create(fileName); fileErr == nil {
			if _, writeErr := file.Write(encoding); writeErr != nil {
				return "", errors.New("failed to write to .json file")
			}
		} else {
			return "", errors.New("*.json file could not be created")
		}
	}
	return id, nil
}

func updateTask(info cliInfo) {

}

func deleteTask(info cliInfo) {

}

func markTaskInProgress(info cliInfo) {

}

func markTaskDone(info cliInfo) {

}

func listTasks(info cliInfo) {

}
