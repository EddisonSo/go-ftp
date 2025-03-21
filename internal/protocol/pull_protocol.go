package protocol

import (
    "log/slog"
    "eddisonso.com/go-ftp/internal/commands"
)

type PullProtocol struct {
    BaseProtocol
    File string
}

func NewPullProtocol(content []byte, logger *slog.Logger) *PullProtocol {
    f := string(content)
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    Logger:	logger,
	    CommandId:  commands.PULL,
	},
	File: f,
    }
}

func (pp *PullProtocol) ToBytes() []byte {
    result := []byte{byte(pp.CommandId)}
    result = append(result, pp.File...)
    return result
}

func (pp *PullProtocol) PrintProtocol() {
    println("PullProtocol")
    println("File: ", pp.File)
}

func (pp *PullProtocol) Execute() {
    return
}
