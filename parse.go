package main

import (
	"errors"
	"strconv"
)

func parseCliArgs(args []string) (cliInfo, error) {

	if len(args) == 0 {
		return cliInfo{}, errors.New("no commands provided")
	}

	cmd, ok := commandEnum[args[0]]
	if !ok {
		return cliInfo{}, errors.New("invalid command")
	}

	numArgs := len(args[1:])

	var id int
	var taskName string
	var status TaskStatus
	var cliInput cliInfo
	switch cmd {
	case add:
		if numArgs != 2 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}
		taskName = args[1]
		cliInput = cliInfo{taskName: taskName}
	case update:
		if numArgs != 3 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		var convError error
		id, convError = strconv.Atoi(args[1])
		if convError != nil {
			return cliInfo{}, errors.New("invalid task id")
		}

		taskName = args[2]
		cliInput = cliInfo{action: cmd, id: id, taskName: taskName}
	case delete:
		if numArgs != 2 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		var convError error
		id, convError = strconv.Atoi(args[1])
		if convError != nil {
			return cliInfo{}, errors.New("invalid task id")
		}

		cliInput = cliInfo{action: cmd, id: id}
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

		cliInput = cliInfo{action: cmd, status: status}
	case markInProgress:
		if numArgs != 1 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		var convError error
		id, convError = strconv.Atoi(args[1])
		if convError != nil {
			return cliInfo{}, errors.New("invalid task id")
		}

		cliInput = cliInfo{action: cmd, id: id}
	case markDone:
		if numArgs != 1 {
			return cliInfo{}, errors.New("invalid number of arguments")
		}

		var convError error
		id, convError = strconv.Atoi(args[1])
		if convError != nil {
			return cliInfo{}, errors.New("invalid task id")
		}

		cliInput = cliInfo{action: cmd, id: id}
	default:

	}

	return cliInput, nil
}
