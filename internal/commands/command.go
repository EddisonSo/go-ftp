package commands

type CommandId uint8 

const (
    PUSH 	CommandId = iota
    PULL
    MOVE
    RENAME
    DELETE
)

type Command struct {
    Id	  CommandId
    Args  []string
}

func ValidCommandId(id CommandId) bool {
    return id >= PUSH && id <= DELETE
}
