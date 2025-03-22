package protocol

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"eddisonso.com/go-ftp/internal/commands"
	"eddisonso.com/go-ftp/internal/filehandler"
)

type PushProtocol struct {
    BaseProtocol
    Size uint32
    InFilename string
    OutFilename string
}

func NewPushProtocol(s uint32, inf string, outf string, logger *slog.Logger) *PushProtocol{
    if len(inf) > 4096 {
	logger.Error("input filename too long")
    }

    if len(outf) > 4096 {
	logger.Error("output filename too long")
    }
    
    return &PushProtocol{
	BaseProtocol: BaseProtocol{
	    Logger: 	logger,
	    CommandId:  commands.PUSH_ID,
	},
	InFilename: inf,
	OutFilename: outf,
	Size: s,
    }
}

func NewPushFromBytes(body []byte, logger *slog.Logger) *PushProtocol{
    s := binary.LittleEndian.Uint32(body[0:4])
    filename := string(body[4:4100])
    filename = strings.TrimRight(filename, "\000")

    return &PushProtocol{
	BaseProtocol: BaseProtocol{
	    Logger: 	logger,
	    CommandId:  commands.PUSH_ID,
	},
	Size: s,
	OutFilename: filename,
    }
}

func (pp *PushProtocol) PrintProtocol() {
    fmt.Println("PushProtocol")
    fmt.Println("Size: ", pp.Size)
    fmt.Println("Input Filename: ", pp.InFilename)
    fmt.Println("Output Filename: ", pp.OutFilename)
}

func (pp *PushProtocol) ExecuteServer(conn net.Conn) {
    args := make([]byte, 4100)
    conn.Read(args)
    pp.Size = binary.LittleEndian.Uint32(args[0:4])
    pp.OutFilename = strings.ReplaceAll(string(args[4:4100]), "\x00", "")
    pp.Logger.Info("Got args: " + fmt.Sprint(pp.Size) + " " + pp.OutFilename)

    writer, err := filehandler.NewFilewriter(string(pp.OutFilename), pp.Logger)
    if err != nil {
	pp.Logger.Error(err.Error())
    }
    
    content := make([]byte, pp.Size)
    conn.Read(content)
    writer.Write(content)
}

func (pp *PushProtocol) ExecuteClient(conn net.Conn) {
    conn.Write([]byte{byte(pp.CommandId)})
    
    var args []byte
    size := make([]byte, 4)
    binary.LittleEndian.PutUint32(size, pp.Size)

    outf := make([]byte, 4096)
    copy(outf, pp.OutFilename)
    args = append(args, size...)
    args = append(args, outf...)
    conn.Write(args)

    reader, err := filehandler.NewFilereader(string(pp.InFilename), pp.Logger)
    if err != nil {
	pp.Logger.Error(err.Error())
    }
    
    content := make([]byte, pp.Size)
    reader.Read(content)
    conn.Write(content)
}
