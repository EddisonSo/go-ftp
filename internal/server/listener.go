package server

import (
	"net"
	"strconv"

	"eddisonso.com/go-ftp/internal/commands"
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
	s.Logger.Info("Waiting for connection...")
	conn, err := listener.Accept()
	s.Logger.Info("Connection from: " + conn.RemoteAddr().String())

	if err != nil {
	    panic(err);
	}

	go func(conn net.Conn) {
	    for {
		c := make([]byte, 1)
		conn.Read(c)
		if err != nil {
		    s.Logger.Error(err.Error())
		    return
		}

		if c[0] == byte(commands.EXIT_ID) {
		    s.Logger.Info("Connection closed by client")
		    conn.Close()
		    break
		}

		p := protocol.NewPushProtocol(0, "", "", s.Logger)
		p.ExecuteServer(conn)
	    }
	}(conn)
    }
}
