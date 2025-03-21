package main

import "eddisonso.com/go-ftp/internal/protocol"
import "eddisonso.com/go-ftp/internal/filehandler"
import "fmt"
import "net"
import "log/slog"
import "os"

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil));
    
    conn, err := net.Dial("tcp", "127.0.0.1:3000")
    if err != nil {
	panic(err)
    }

    reader, err := filehandler.NewFilereader("input.txt", logger)
    if err != nil {
	panic(err)
    }

    n, err := reader.Getsize()
    if err != nil {
	panic(err)
    }

    body := make([]byte, n)
    reader.Read(body)

    protocol := protocol.NewPushProtocol(n, body)
    fmt.Println(protocol.ToBytes())

    conn.Write(protocol.ToBytes())
}
