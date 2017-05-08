package http

import (
	"errors"
	"net/http"
	"strings"

	"../../lib/helper"
	"../../modules"
	"github.com/gorilla/mux"
)

type Server struct {
	Port    string
	Running bool
	Modules []*modules.Module
	mux     *mux.Router
}

func NewServer(port string) (*Server, error) {
	if port == "" {
		return nil, errors.New("A port would be nice")
	}
	s := &Server{Port: port, Running: false}
	s.Modules = make([]*modules.Module, 0)

	// initialize gorilla
	m := mux.NewRouter()

	s.mux = m
	return s, nil
}

func (s *Server) AddModule(m *modules.Module) error {
	// check if module is already in
	for _, ms := range s.Modules {
		if m.GetName() == ms.GetName() {
			return errors.New("module is already part of this pack.")
		}
	}
	s.Modules = append(s.Modules, m)
	return nil
}

func (s *Server) Start() error {
	if s.Port == "" {
		return errors.New("No port given, how did you do that?")
	}
	if s.Running {
		return errors.New("server is already running, how did you do that?")
	}

	if e := s.start(); e != nil {
		return e
	}
	s.Running = true
	return nil
}

func (s *Server) Stop() error {
	if !s.Running {
		return errors.New("server is not running.")
	}
	if e := s.stop(); e != nil {
		return e
	}
	s.Running = false
	return nil
}

func (s *Server) start() error {
	// internal start
	// TODO(@krizzle): What happens when the routes are already added? look at gorilla docs/src
	for _, m := range s.Modules {
		for _, r := range m.GetRoutes() {
			s.mux.HandleFunc(strings.Join(stringhelper.Map([]string{m.GetName(), r.Path}, strings.ToLower), "/"), r.Func)
		}
	}

	// TODO(@krizzle): To consider: Run the server on main-thread or in a goroutine?
	return http.ListenAndServe(":"+s.Port, s.mux)

}

func (s *Server) stop() error {
	// internal stop
	return nil
}
