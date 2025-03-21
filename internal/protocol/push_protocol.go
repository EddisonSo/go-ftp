package protocol

import (
    "encoding/binary"
    "fmt"
    "log/slog"
    "eddisonso.com/go-ftp/internal/commands"
    "eddisonso.com/go-ftp/internal/filehandler"
)

type PushProtocol struct {
    BaseProtocol
    Size uint32
    Filename string
    Content []byte
}

func NewPushProtocol(body []byte, logger *slog.Logger) *PushProtocol{
    s := binary.LittleEndian.Uint32(body[0:4])
    content := body[4:]

    return &PushProtocol{
	BaseProtocol: BaseProtocol{
	    Logger: 	logger,
	    CommandId:  commands.PUSH,
	},
	Size: s,
	Content: content,
    }
}

func (pp *PushProtocol) ToBytes() []byte {
    size := make([]byte, 4)
    binary.LittleEndian.PutUint32(size, pp.Size)

    filename := make([]byte, 4096)
    copy(filename, pp.Filename)

    result := []byte{byte(pp.CommandId)}
    result = append(result, size...)
    result = append(result, filename...)
    result = append(result, pp.Content...)
    return result
}

func (pp *PushProtocol) PrintProtocol() {
    fmt.Println("PushProtocol")
    fmt.Println("Size: ", pp.Size)
    fmt.Println("Filename: ", pp.Filename)
    fmt.Println("Content: ", string(pp.Content))
}

func (pp *PushProtocol) GetContent() []byte {
    return pp.Content
}

func (pp *PushProtocol) Execute() {
    writer, err := filehandler.NewFilewriter(pp.Filename, pp.Logger)
    if err != nil {
	pp.Logger.Error(err.Error())
    }
    
    writer.Write(pp.Content)
}
