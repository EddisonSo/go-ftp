package protocol

import (
    "eddisonso.com/go-ftp/internal/filehandler"
    "eddisonso.com/go-ftp/internal/commands"
    "encoding/binary"
    "fmt"
    "log/slog"
    "net"
    "strings"
)

type PushProtocol struct {
    BaseProtocol
    Size uint32
    Filename string
    Content []byte
}

func NewPushProtocol(s uint32, f string, content []byte, logger *slog.Logger) *PushProtocol{
    if len(f) > 4096 {
	logger.Error("Filename too long")
    }
    
    return &PushProtocol{
	BaseProtocol: BaseProtocol{
	    Logger: 	logger,
	    CommandId:  commands.PUSH_ID,
	},
	Filename: f,
	Size: s,
	Content: content,
    }
}

func NewPushFromBytes(body []byte, logger *slog.Logger) *PushProtocol{
    s := binary.LittleEndian.Uint32(body[0:4])
    filename := string(body[4:4100])
    filename = strings.TrimRight(filename, "\000")
    content := body[4100:]

    return &PushProtocol{
	BaseProtocol: BaseProtocol{
	    Logger: 	logger,
	    CommandId:  commands.PUSH_ID,
	},
	Size: s,
	Filename: filename,
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

func (pp *PushProtocol) ExecuteServer(conn net.Conn) {
    writer, err := filehandler.NewFilewriter(string(pp.Filename), pp.Logger)
    if err != nil {
	pp.Logger.Error(err.Error())
    }
    
    writer.Write(pp.Content)
}

func (pp *PushProtocol) ExecuteClient(conn net.Conn) {
    conn.Write(pp.ToBytes())
}
