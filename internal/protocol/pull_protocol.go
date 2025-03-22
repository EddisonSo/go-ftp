package protocol

import (
    "eddisonso.com/go-ftp/internal/commands"
    "log/slog"
    "net"
)

type PullProtocol struct {
    BaseProtocol
    File string
}

func NewPullProtocol(f string, logger *slog.Logger) *PullProtocol {
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    Logger:	logger,
	    CommandId:  commands.CommandId(commands.PUSH_ID),
	},
	File: f,
    }
}

func NewPullFromBytes(content []byte, logger *slog.Logger) *PullProtocol {
    f := string(content)
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    Logger:	logger,
	    CommandId:  commands.PULL_ID,
	},
	File: f,
    }
}

func (pp *PullProtocol) PrintProtocol() {
    println("PullProtocol")
    println("File: ", pp.File)
}

func (pp *PullProtocol) ExecuteClient(conn net.Conn) {
    return
}

func (pp *PullProtocol) ExecuteServer(conn net.Conn) {
    return
}
