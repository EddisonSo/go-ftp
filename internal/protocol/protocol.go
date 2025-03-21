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
    Size   uint32;
    Body   []byte;
}

type Protocol interface {
    ToBytes() []byte;
    GetBody() []byte;
}

func (bp *BaseProtocol) ToBytes() []byte {
    size := make([]byte, 4)
    binary.LittleEndian.PutUint32(size, bp.Size)

    result := []byte{byte(bp.CommandId)}
    result = append(result, size...)
    result = append(result, bp.Body...)
    return result
}

func (bp *BaseProtocol) GetBody() []byte {
    return bp.Body
}

func PrintProtocol(p BaseProtocol) {
    println("Protocol:")
    println("Command: ", p.CommandId, " (", commands.GetCommandName(p.CommandId), ")")
    println("Size:", p.Size)
    println("Body:", string(p.Body))
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
