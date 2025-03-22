package main

import (
    "log/slog"
    "eddisonso.com/go-ftp/internal/client"
    "eddisonso.com/go-ftp/internal/config"
)

func main() {
    logger := slog.Logger{}
    cConfig := config.ClientConfig{Host: "localhost", Port: 3000, Homedir: "~/"}
    c := client.NewClient(cConfig, &logger)
    term := client.NewTerm(c)
    term.Prompt()
}

