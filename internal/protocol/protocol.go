package protocol

import (
    "log/slog"
    "errors"
    "net"
    "eddisonso.com/go-ftp/internal/commands"
)

type BaseProtocol struct {
    Logger    *slog.Logger
    CommandId commands.CommandId;
    Other     []byte;
}

type Protocol interface {
    PrintProtocol();
    ExecuteServer(conn net.Conn);
    ExecuteClient(conn net.Conn);
}

func PrintProtocol(p Protocol) {
    p.PrintProtocol()
}

func NewProtocol(cbyte byte, args []byte, logger *slog.Logger) (Protocol, error) {
    command := commands.CommandId(cbyte)
    if !commands.ValidCommandId(command){
	return nil, errors.New("Invalid protocol command id, got: " + string(cbyte))
    }
    
    switch command {
	case commands.PUSH_ID:
	    return NewPushFromBytes(args, logger), nil
	case commands.PULL_ID:
	    return NewPullFromBytes(args, logger), nil
    }
    return nil, nil;
}
