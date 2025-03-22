package commands

type CommandId uint8 

const (
    PUSH_ID 	CommandId = iota
    PULL_ID
    MOVE_ID
    RENAME_ID
    DELETE_ID
)

const (
    PUSH = "push"
    PULL = "pull"
    MOVE = "move"
    RENAME = "rename"
    DELETE = "delete"
)

var Commands = []string{
    PUSH,
    PULL,
    MOVE,
    RENAME,
    DELETE,
}

type Command struct {
    Id	  CommandId
    Args  []string
}

func ValidCommandId(id CommandId) bool {
    return id >= PUSH_ID && id <= DELETE_ID
}
