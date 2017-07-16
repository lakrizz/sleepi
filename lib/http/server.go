package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"../../lib/helper"
	"../../modules"
	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
	"github.com/urfave/negroni"
)

type Server struct {
	Port    string
	Running bool
	Modules []modules.Module
	mux     *mux.Router
	neg     *negroni.Negroni
}

func NewServer(port string) (*Server, error) {
	if port == "" {
		return nil, errors.New("A port would be nice")
	}
	s := &Server{Port: port, Running: false}
	s.Modules = make([]modules.Module, 0)

	// initialize gorilla
	m := mux.NewRouter()
	n := negroni.Classic()

	s.mux = m
	s.neg = n
	return s, nil
}

func (s *Server) AddModule(m modules.Module) error {
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
			s.mux.HandleFunc(strings.Join(stringhelper.Map([]string{"/" + m.GetName(), r.Path}, strings.ToLower), "/"), r.Func)
			fmt.Println("adding route", m.GetName(), r.Path)
		}
	}

	// TODO(@krizzle): To consider: Run the server on main-thread or in a goroutine?

	s.mux.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./pub/css"))))
	s.mux.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./pub/js"))))
	s.mux.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./pub/img"))))

	s.neg.UseHandler(s.mux)
	pp.Println(s.mux)
	return http.ListenAndServe(":"+s.Port, s.neg)
}

func (s *Server) stop() error {
	// internal stop
	return nil
}
