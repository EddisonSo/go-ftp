package commands

type Commandhandler struct {}

func (c *Commandhandler) HandleCommand(input string) {
    command := parseCommand(input)

    switch command.Id {
	case PUSH:
		// Handle push command
	case PULL:
		// Handle pull command
	case MOVE:
		// Handle move command
	case RENAME:
		// Handle rename command
	case DELETE:
		// Handle delete command
	}
}
