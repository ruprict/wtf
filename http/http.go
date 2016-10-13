package http

import (
	"net"
	"net/http"
)

type Server struct {
	ln net.Listener

	// Handler to serve
	Handler *Handler

	//Bind address to open
	Addr string
}

func (s *Server) Open() error {

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	// Start HTTP Server
	go func() { http.Serve(s.ln, s.Handler) }()
	return nil

}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}

	return nil

}
