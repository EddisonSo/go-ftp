package main

import (
	"log/slog"
	"os"
	"eddisonso.com/go-ftp/internal/server"
	"eddisonso.com/go-ftp/internal/config"
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil));
    host := config.Host{Hostname:"0.0.0.0"};
    config := config.ServerConfig{Host:host, Port:3000};
    server, _ := server.GetServer(config, logger);
    server.Listen();
}
