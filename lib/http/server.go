package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	modules "../../controllers"

	stringhelper "../../lib/helper"
	"github.com/gorilla/mux"
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
	for _, m := range s.Modules {
		for _, r := range m.GetRoutes() {
			s.mux.HandleFunc(strings.Join(stringhelper.Map([]string{"/" + m.GetName(), r.Path}, strings.ToLower), ""), r.Func)
			fmt.Println("adding route", m.GetName(), r.Path)
		}
	}

	// TODO(@kk): refactor in a method
	s.handleStatic([]string{"css", "js", "img"})

	s.neg.UseHandler(s.mux)
	return http.ListenAndServe(":"+s.Port, s.neg)
}

func (s *Server) handleStatic(paths []string) {
	for _, v := range paths {
		s.mux.PathPrefix(v).Handler(http.StripPrefix(v, http.FileServer(http.Dir(fmt.Sprintf("./pub/%s", v)))))
	}
}

func (s *Server) stop() error {
	// internal stop
	return nil
}
