package server

import "fmt"
import "net"
import "io"
import "strconv"
import "eddisonso.com/go-ftp/internal/protocol"

type Listener interface {
    Listen() error;
}

func (s *Server) Listen() {
    fmt.Println("Listening on port: " + strconv.Itoa(s.Config.Port));
    listener, err := net.Listen("tcp", ":" + strconv.Itoa(s.Config.Port))

    if err != nil {
	panic(err);
    }

    defer listener.Close();

    for {
	conn, err := listener.Accept()
	fmt.Println("Accepted connection from: " + conn.RemoteAddr().String());

	if err != nil {
	    panic(err);
	}

	go func(conn net.Conn) {
	    defer conn.Close()

	    data, err := io.ReadAll(conn)
	    if err != nil {
		panic(err)
	    }

	    fmt.Println("Data: " + string(data))

	    p, err := protocol.FromBytes(data)
	    if err != nil {
		panic(err)
	    }

	    protocol.PrintProtocol(p)
	}(conn)
    }
}
