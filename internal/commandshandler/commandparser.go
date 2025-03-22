package commandshandler

import (
    "errors"
    "strings"
    "eddisonso.com/go-ftp/internal/commands"
)

func GetCommandId(command string) (commands.CommandId) {
    switch command {
	case "EXIT":
	    return commands.EXIT_ID
	case "PUSH":
	    return commands.PUSH_ID
	case "PULL":
	    return commands.PULL_ID
	case "MOVE":
	    return commands.MOVE_ID
	case "RENAME":
	    return commands.RENAME_ID
	case "DELETE":
	    return commands.DELETE_ID
    }
    return 255
}

func parseCommand(input string) (*commands.Command, error) {
    inputsplit := strings.SplitN(input, " ", 2)
    command := inputsplit[0]
    command = strings.Trim(command, " ")
    command = strings.ToUpper(command)
    commandId := GetCommandId(command)
    if commandId == 255 {
	return nil, errors.New("Invalid command")
    }

    args := []string{}
    if len(inputsplit) > 1 {
	args = strings.Split(inputsplit[1], " ")
    }

    return &commands.Command{Id:commandId, Args: args}, nil
}
