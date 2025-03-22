package commands

type Commandhandler struct {}

func (c *Commandhandler) HandleCommand(input string) {
    command := parseCommand(input)

    switch command.Id {
	case PUSH_ID:
		// Handle push command
	case PULL_ID:
		// Handle pull command
	case MOVE_ID:
		// Handle move command
	case RENAME_ID:
		// Handle rename command
	case DELETE_ID:
		// Handle delete command
	}
}
