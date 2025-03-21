package commands

import (
    "errors"
    "strings"
)

func GetCommandId(command string) (CommandId, error) {
    switch command {
	case "PUSH":
	    return PUSH, nil
	case "PULL":
	    return PULL, nil
	case "MOVE":
	    return MOVE, nil
	case "RENAME":
	    return RENAME, nil
	case "DELETE":
	    return DELETE, nil
    }
    return 255, errors.New("Invalid Command: " + command)
}

func GetCommandName(commandId CommandId) string {
    switch commandId {
    case PUSH:
	return "PUSH"
    case PULL:
	return "PULL"
    case MOVE:
	return "MOVE"
    case RENAME:
	return "RENAME"
    case DELETE:
	return "DELETE"
    }
    return "INVALID"
}

func parseCommand(command string) Command {
    commandId, err := GetCommandId(command)
    if err != nil {
	panic(err)
    }

    args := strings.Split(command, " ")[1:]
    return Command{Id:commandId, Args: args}
}
