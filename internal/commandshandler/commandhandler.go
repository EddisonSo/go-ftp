package commandshandler

import (
    "log/slog"
    "net"
    "eddisonso.com/go-ftp/internal/commands"
    "eddisonso.com/go-ftp/internal/filehandler"
    "eddisonso.com/go-ftp/internal/protocol"
)

func HandleCommand(input string, logger *slog.Logger, conn net.Conn) commands.CommandId {
    command, err := parseCommand(input)
    if err != nil {
	return 255
    }

    switch command.Id {
	case commands.EXIT_ID:
	    conn.Write([]byte{byte(commands.EXIT_ID)})
	    return commands.EXIT_ID
	case commands.PUSH_ID:
	    if len(command.Args) < 2 {
		logger.Error("Invalid number of arguments for push command")
		return 255
	    }
	    reader, err := filehandler.NewFilereader(command.Args[0], logger)
	    if err != nil {
		logger.Error(err.Error())
		return 255
	    }

	    n, err := reader.Getsize()
	    if err != nil {
		logger.Error(err.Error())
		return 255
	    }

	    p := protocol.NewPushProtocol(n, command.Args[0], command.Args[1], logger)
	    p.ExecuteClient(conn)
	    return commands.PUSH_ID
	case commands.PULL_ID:
		// Handle pull command
	case commands.MOVE_ID:
		// Handle move command
	case commands.RENAME_ID:
		// Handle rename command
	case commands.DELETE_ID:
		// Handle delete command
	}
	return 255
}
