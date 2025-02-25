package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hse-telesope/template-service/internal/config"
	"github.com/hse-telesope/template-service/internal/providers/people"
)

type Provider interface {
	GetPeople(ctx context.Context) ([]people.Person, error)
}

type Server struct {
	server   http.Server
	provider Provider
}

func New(conf config.Config, provider Provider) *Server {
	s := new(Server)
	s.server.Addr = fmt.Sprintf(":%d", conf.Port)
	s.server.Handler = s.setRouter()
	s.provider = provider
	return s
}

func (s *Server) setRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", s.pingHandler)
	mux.HandleFunc("GET /people", s.getPeopleHandler)
	return mux
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
