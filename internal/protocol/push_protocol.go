package protocol

import "eddisonso.com/go-ftp/internal/commands"

type PushProtocol struct {
    BaseProtocol
}

func NewPushProtocol(size uint32, body []byte) *PullProtocol {
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    CommandId:  commands.PUSH,
	    Size:   size,
	    Body:   body,
	},
    }
}
