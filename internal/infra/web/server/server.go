package server

import (
	"net/http"
)

type Handlers map[string]http.Handler

type Server struct {
	addr     string
	handlers Handlers
}

func NewServer(addr string, handlers Handlers) *Server {
	return &Server{addr, handlers}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	for pattern, handler := range s.handlers {
		mux.Handle(pattern, handler)
	}
	return http.ListenAndServe(s.addr, mux)
}
