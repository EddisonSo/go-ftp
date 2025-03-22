package client

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"
	"strings"
	"eddisonso.com/go-ftp/internal/commands"
	"eddisonso.com/go-ftp/internal/commandshandler"
	"github.com/peterh/liner"
)

type term struct {
    conn *net.Conn
    tabPressed bool
    initialPressed bool
    Logger	 *slog.Logger
}

const (
    blue  = "\033[34m" // Blue for directories
    green = "\033[32m" // Green for files
    reset = "\033[0m"  // Reset color
)

func (t *term) completePath(word string) []string {
    dir := word
    if !strings.HasSuffix(word, string(os.PathSeparator)) {
        dir = filepath.Dir(word)
    }

    pattern := filepath.Join(dir, "*")
    matches, err := filepath.Glob(pattern)
    if err != nil {
        return nil
    }

    var filteredMatches []string
    for _, match := range matches {
        if strings.HasPrefix(match, word) {
            filteredMatches = append(filteredMatches, match)
        }
    }

    return filteredMatches
}

func (t *term) completeCommand(word string) []string {
    var matches []string
    for _, command := range commands.Commands{
	if strings.HasPrefix(command, word) {
	    matches = append(matches, command)
	}
    }
    return matches
}


func (t *term) Prompt() {
    l := liner.NewLiner()
    defer l.Close()

    l.SetWordCompleter(func(line string, pos int) (string, []string, string) {
        wordStart := pos
        for wordStart > 0 && line[wordStart-1] != ' ' {
            wordStart--
        }

        word := line[wordStart:pos]

	var matches []string
	if wordStart == 0 {
	    matches = t.completeCommand(word)
	} else {
	    matches = t.completePath(word)
	}

        return line[:wordStart], matches, line[pos:]
    })

    for {
        line, err := l.Prompt("> ")
        if err != nil {
            fmt.Println("Error reading input:", err)
            break
        }

        t.tabPressed = false
	ret := commandshandler.HandleCommand(line, t.Logger, *t.conn)

	if ret == commands.EXIT_ID {
	    os.Exit(0)
	}

	if ret == 255 {
	    fmt.Println("Invalid command")
	}
    }
}


func NewTerm(c *net.Conn, logger *slog.Logger) *term {
    logger.Info("Creating new term")
    return &term{conn: c, tabPressed: false, initialPressed: false, Logger: logger}
}

