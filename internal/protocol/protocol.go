package protocol

import (
    "log/slog"
    "eddisonso.com/go-ftp/internal/commands"
    "errors"
    "fmt"
    "net"
)

type BaseProtocol struct {
    Logger    *slog.Logger
    CommandId commands.CommandId;
    Other     []byte;
}

type Protocol interface {
    ToBytes() []byte;
    PrintProtocol();
    ExecuteServer(conn net.Conn);
    ExecuteClient(conn net.Conn);
}

func PrintProtocol(p Protocol) {
    p.PrintProtocol()
}

func FromBytes(data []byte, logger *slog.Logger) (Protocol, error) {
    if len(data) <= 1 {
	return nil, errors.New("Invalid protocol size, got: " + fmt.Sprint(len(data)) + " <= 1")
    }
    command := commands.CommandId(data[0])
    if !commands.ValidCommandId(command){
	return nil, errors.New("Invalid protocol command id, got: " + string(data[0]))
    }
    
    body := data[1:]

    switch command {
	case commands.PUSH:
	    return NewPushFromBytes(body, logger), nil
	case commands.PULL:
	    return NewPullFromBytes(body, logger), nil
    }
    return nil, nil;
}
