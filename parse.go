package main

import (
	"errors"
)

func parseCliArgs(args []string) (cliInfo, error) {

	if len(args) == 0 {
		return cliInfo{}, errors.New("no commands provided")
	}

	cmd, ok := commandEnum[args[0]]
	if !ok {
		return cliInfo{}, errors.New("invalid command")
	}

	numArgs := len(args)

	var id string
	var taskName string
	var status TaskStatus
	var cliInput cliInfo
	switch cmd {
	case add:
		if numArgs != 2 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}
		taskName = args[1]
		cliInput = cliInfo{TaskName: taskName}
	case update:
		if numArgs != 3 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		id = args[1]
		taskName = args[2]

		cliInput = cliInfo{Action: cmd, Id: id, TaskName: taskName}
	case delete:
		if numArgs != 2 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		id = args[1]

		cliInput = cliInfo{Action: cmd, Id: id}
	case list:
		if numArgs > 2 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		if numArgs == 2 {
			status, ok = taskEnum[args[1]]
			if !ok {
				return cliInfo{}, errors.New("not a valid task status")
			}
		}

		cliInput = cliInfo{Action: cmd, Status: status}
	case markInProgress:
		if numArgs != 1 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		id = args[1]

		cliInput = cliInfo{Action: cmd, Id: id}
	case markDone:
		if numArgs != 1 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		id = args[1]
		cliInput = cliInfo{Action: cmd, Id: id}
	default:

	}

	return cliInput, nil
}
