package commands

type CommandId uint8 

const (
    EXIT_ID CommandId = iota
    PUSH_ID
    PULL_ID
    MOVE_ID
    RENAME_ID
    DELETE_ID
)

const (
    EXIT = "exit"
    PUSH = "push"
    PULL = "pull"
    MOVE = "move"
    RENAME = "rename"
    DELETE = "delete"
)

var Commands = []string{
    EXIT,
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
    return id >= EXIT_ID && id <= DELETE_ID
}
