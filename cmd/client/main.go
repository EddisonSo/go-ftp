package main

import "eddisonso.com/go-ftp/internal/protocol"
import "eddisonso.com/go-ftp/internal/filehandler"
import "fmt"
import "net"
import "log"

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:3000")
    if err != nil {
	log.Fatal(err)
    }

    reader, err := filehandler.NewFilereader("input.txt")
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
