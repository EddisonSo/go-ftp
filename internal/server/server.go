package server 

import "eddisonso.com/go-ftp/internal/types"

type Server struct {
    Config types.ServerConfig;
}

func GetServer(config types.ServerConfig) (*Server, error) {
    return &Server{Config: config}, nil;
}


