package protocol

import (
	"encoding/binary"
	"fmt"
	"eddisonso.com/go-ftp/internal/commands"
)

type PushProtocol struct {
    BaseProtocol
    Size uint32
    Content []byte
}

func NewPushProtocol(size uint32, body []byte) *PushProtocol{
    s := binary.LittleEndian.Uint32(body[0:4])
    content := body[4:]

    return &PushProtocol{
	BaseProtocol: BaseProtocol{
	    CommandId:  commands.PUSH,
	},
	Size: s,
	Content: content,
    }
}

func (pp *PushProtocol) ToBytes() []byte {
    size := make([]byte, 4)
    binary.LittleEndian.PutUint32(size, pp.Size)

    result := []byte{byte(pp.CommandId)}
    result = append(result, size...)
    result = append(result, pp.Content...)
    return result
}

func (pp *PushProtocol) PrintProtocol() {
    fmt.Println("PushProtocol")
    fmt.Println("Size: ", pp.Size)
    fmt.Println("Content: ", string(pp.Content))
}

func (p *PushProtocol) GetContent() []byte {
    return p.Content
}
