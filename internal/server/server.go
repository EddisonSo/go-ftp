package server 

import (
    "eddisonso.com/go-ftp/internal/config"
    "log/slog"
)

type Server struct {
    Config config.ServerConfig;
    Logger *slog.Logger;
}

func GetServer(config config.ServerConfig, logger *slog.Logger) (*Server, error) {
    return &Server{
	Config: config,
	Logger: logger,
    }, nil;
}


