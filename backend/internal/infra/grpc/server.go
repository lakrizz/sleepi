package grpc

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	srv *http.Server
	mux *http.ServeMux

	routes []string
}

func New() (*Server, error) {
	mux := http.NewServeMux()

	p := new(http.Protocols)

	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	s := &http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: p,
	}

	server := &Server{
		srv:    s,
		mux:    mux,
		routes: make([]string, 0),
	}

	return server, nil
}

func (s *Server) Start() error {
	slog.Info("starting grpc server", "address", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) RegisterHandler(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
	s.routes = append(s.routes, fmt.Sprintf("%v", path))
}

func (s *Server) DebugRoutes() {
	for _, r := range s.routes {
		slog.Debug(r)
	}
}
