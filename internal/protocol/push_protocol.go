package protocol

type PushProtocol struct {
    BaseProtocol
}

func NewPushProtocol(size uint32, body []byte) *PullProtocol {
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    Type:   "PULL",
	    Size:   size,
	    Body:   body,
	},
    }
}
