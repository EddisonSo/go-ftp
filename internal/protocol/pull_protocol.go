package protocol

import "eddisonso.com/go-ftp/internal/commands"

type PullProtocol struct {
    BaseProtocol
}

func NewPullProtocol(size uint32, body []byte) *PullProtocol {
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    CommandId:  commands.PULL,
	    Size:   size,
	    Body:   body,
	},
    }
}
