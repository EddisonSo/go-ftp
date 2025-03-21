package client

import (
    "eddisonso.com/go-ftp/internal/config"
    "log/slog"
)

type Client struct {
    Config config.ClientConfig
    Logger *slog.Logger
}


func NewClient(config config.ClientConfig, logger *slog.Logger) *Client {
    return &Client{
	Config: config,
	Logger: logger,
    }
}
