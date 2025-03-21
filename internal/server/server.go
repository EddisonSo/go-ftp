package server 

import "eddisonso.com/go-ftp/internal/types"
import "log/slog"
import "os"

type Server struct {
    Config types.ServerConfig;
    logger *slog.Logger;
}

func GetServer(config types.ServerConfig) (*Server, error) {
    return &Server{
	Config: config,
	logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
    }, nil;
}


