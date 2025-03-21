package protocol

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"eddisonso.com/go-ftp/internal/commands"
)

type BaseProtocol struct {
    CommandId commands.CommandId;
    Other     []byte;
}

type Protocol interface {
    ToBytes() []byte;
    PrintProtocol();
}

func PrintProtocol(p Protocol) {
    p.PrintProtocol()
}

func FromBytes(data []byte) (Protocol, error) {
    if len(data) <= 5 {
	return nil, errors.New("Invalid protocol size, got: " + fmt.Sprint(len(data)) + " <= 5")
    }
    command := commands.CommandId(data[0])
    if !commands.ValidCommandId(command){
	return nil, errors.New("Invalid protocol command id, got: " + string(data[0]))
    }

    size := binary.LittleEndian.Uint32(data[1:5])
    body := data[5:]
    if uint32(len(body)) != size {
	return nil, errors.New("Invalid body size, got: " + strconv.FormatUint(uint64(len(body)), 10) + " expected: " + strconv.FormatUint(uint64(size), 10))
    }

    switch command {
	case commands.PUSH:
	    return NewPullProtocol(size, body), nil
	case commands.PULL:
	    return NewPushProtocol(size, body), nil
    }
    return nil, nil;
}
