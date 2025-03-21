package protocol

import (
    "errors"
    "fmt"
    "eddisonso.com/go-ftp/internal/commands"
    "log/slog"
)

type BaseProtocol struct {
    Logger    *slog.Logger
    CommandId commands.CommandId;
    Other     []byte;
}

type Protocol interface {
    ToBytes() []byte;
    PrintProtocol();
    Execute();
}

func PrintProtocol(p Protocol) {
    p.PrintProtocol()
}

func FromBytes(data []byte, logger *slog.Logger) (Protocol, error) {
    if len(data) <= 5 {
	return nil, errors.New("Invalid protocol size, got: " + fmt.Sprint(len(data)) + " <= 5")
    }
    command := commands.CommandId(data[0])
    if !commands.ValidCommandId(command){
	return nil, errors.New("Invalid protocol command id, got: " + string(data[0]))
    }
    
    body := data[1:]

    switch command {
	case commands.PUSH:
	    return NewPullProtocol(body, logger), nil
	case commands.PULL:
	    return NewPushProtocol(body, logger), nil
    }
    return nil, nil;
}
