package main

import (
    "log/slog"
    "eddisonso.com/go-ftp/internal/client"
    "eddisonso.com/go-ftp/internal/config"
    "net"
    "strconv"
    "os"
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil));
    cConfig := config.ClientConfig{Host: "localhost", Port: 3000, Homedir: "~/"}
    c := client.NewClient(cConfig, logger)

    conn, err := net.Dial("tcp", c.Config.Host + ":" + strconv.Itoa(c.Config.Port))
    if err != nil {
	panic(err)
    }

    term := client.NewTerm(&conn, logger)
    term.Prompt()
}

