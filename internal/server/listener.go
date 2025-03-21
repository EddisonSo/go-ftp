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
    s.Logger.Info("Listening on: " + s.Config.Host.Hostname + ":" + strconv.Itoa(s.Config.Port))

    if err != nil {
	panic(err);
    }

    defer listener.Close();

    for {
	conn, err := listener.Accept()
	s.Logger.Info("Connection from: " + conn.RemoteAddr().String())

	if err != nil {
	    panic(err);
	}

	go func(conn net.Conn) {
	    defer conn.Close()

	    data, err := io.ReadAll(conn)
	    if err != nil {
		panic(err)
	    }

	    s.Logger.Info("Got: " + string(data))

	    p, err := protocol.FromBytes(data)
	    if err != nil {
		s.Logger.Error(err.Error())
		return
	    }

	    writer, err := filehandler.NewFilewriter("output.txt", s.Logger)
	    if err != nil {
		s.Logger.Error(err.Error())
		return
	    }

	    writer.Write(p.GetBody())
	}(conn)
    }
}
