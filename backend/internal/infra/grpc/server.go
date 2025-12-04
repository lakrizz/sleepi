package grpc

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-michi/michi"
)

type Server struct {
	srv *http.Server
	mux *michi.Router

	routes []string
}

func New() (*Server, error) {
	michi := michi.NewRouter()

	c := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
	})

	michi.Use(middleware.Logger)

	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(false)

	s := &http.Server{
		Addr:      ":8080",
		Handler:   c(michi),
		Protocols: p,
	}

	server := &Server{
		srv:    s,
		mux:    michi,
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
