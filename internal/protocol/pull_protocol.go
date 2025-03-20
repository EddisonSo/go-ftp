package protocol

type PullProtocol struct {
    BaseProtocol
}

func NewPullProtocol(size uint32, body []byte) *PullProtocol {
    return &PullProtocol{
	BaseProtocol: BaseProtocol{
	    Type:   "PULL",
	    Size:   size,
	    Body:   body,
	},
    }
}
