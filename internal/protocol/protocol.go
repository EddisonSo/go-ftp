package protocol

import "encoding/binary"
import "log"

type BaseProtocol struct {
    Type   string;
    Size   uint32;
    Body   []byte;
}

type Protocol interface {
    ToBytes() []byte;
    GetProtocolType() string;
    GetSize() uint32;
    GetBody() []byte;
}

func (bp *BaseProtocol) ToBytes() []byte {
    size := make([]byte, 4)
    binary.LittleEndian.PutUint32(size, bp.Size)

    result := append([]byte(bp.Type), size...)
    result = append(result, bp.Body...)
    return result
}

func (bp *BaseProtocol) GetProtocolType() string {
    return bp.Type
}

func (bp *BaseProtocol) GetSize() uint32{
    return bp.Size
}

func (bp *BaseProtocol) GetBody() []byte {
    return bp.Body
}

func PrintProtocol(p Protocol) {
    println("Protocol:")
    println("Type:", p.GetProtocolType())
    println("Size:", p.GetSize())
    println("Body:", string(p.GetBody()))
}

func FromBytes(data []byte) (Protocol, error) {
    t := string(data[:4])
    if t != "PULL" && t != "PUSH" {
	log.Fatal("Invalid Protocol Type, got:", t)
    }

    size := binary.LittleEndian.Uint32(data[4:8])
    body := data[8:]
    if uint32(len(body)) != size {
	return nil, nil
    }

    switch t {
	case "PULL":
	    return NewPullProtocol(size, body), nil
	case "PUSH":
	    return NewPushProtocol(size, body), nil
    }
    return nil, nil;
}
