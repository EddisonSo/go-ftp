package main

import "eddisonso.com/go-ftp/internal/types"
import "eddisonso.com/go-ftp/internal/server"

func main() {
    host := types.Host{Hostname:"0.0.0.0"};
    config := types.ServerConfig{Host:host, Port:3000};
    server, _ := server.GetServer(config);
    server.Listen();
}
