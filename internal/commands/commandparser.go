package commands

import (
    "errors"
    "strings"
)

func GetCommandId(command string) (CommandId, error) {
    switch command {
	case "PUSH":
	    return PUSH_ID, nil
	case "PULL":
	    return PULL_ID, nil
	case "MOVE":
	    return MOVE_ID, nil
	case "RENAME":
	    return RENAME_ID, nil
	case "DELETE":
	    return DELETE_ID, nil
    }
    return 255, errors.New("Invalid Command: " + command)
}

func GetCommandName(commandId CommandId) string {
    switch commandId {
    case PUSH_ID:
	return "PUSH"
    case PULL_ID:
	return "PULL"
    case MOVE_ID:
	return "MOVE"
    case RENAME_ID:
	return "RENAME"
    case DELETE_ID:
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
