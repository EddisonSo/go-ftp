package main

import "eddisonso.com/go-ftp/internal/protocol"
import "fmt"
import "net"
import "log"

func main() {
    protocol := protocol.NewPushProtocol(8, []byte("TESTTEST"))
    fmt.Println(protocol.ToBytes())
    conn, err := net.Dial("tcp", "127.0.0.1:3000")
    if err != nil {
	log.Fatal(err)
    }
    conn.Write(protocol.ToBytes())
}
