package server

import (
	"io"
	"net"
	"strconv"
	"eddisonso.com/go-ftp/internal/filehandler"
	"eddisonso.com/go-ftp/internal/protocol"
)

type Listener interface {
    Listen() error;
}

func (s *Server) Listen() {
    listener, err := net.Listen("tcp", s.Config.Host.Hostname + ":" + strconv.Itoa(s.Config.Port))
    s.logger.Info("Listening on: " + s.Config.Host.Hostname + ":" + strconv.Itoa(s.Config.Port))

    if err != nil {
	panic(err);
    }

    defer listener.Close();

    for {
	conn, err := listener.Accept()
	s.logger.Info("Connection from: " + conn.RemoteAddr().String())

	if err != nil {
	    panic(err);
	}

	go func(conn net.Conn) {
	    defer conn.Close()

	    data, err := io.ReadAll(conn)
	    if err != nil {
		panic(err)
	    }

	    s.logger.Info("Got: " + string(data))

	    p, err := protocol.FromBytes(data)
	    if err != nil {
		panic(err)
	    }

	    protocol.PrintProtocol(p)

	    writer, err := filehandler.NewFilewriter("output.txt")
	    if err != nil {
		panic(err)
	    }

	    writer.Write(p.GetBody())
	}(conn)
    }
}
