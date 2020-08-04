package config

import (
	"net/http"

	config "github.com/chutified/market-info/config"
)

// Server controls the web server behaviour.
type Server struct {
	srv *http.Server
}

// New is a constructor for the Server.
func New() *Server {
	return &Server{}
}

// Set applies the configuration, set up the services and the server.
func (s *Server) Set(cfg *config.Config) error {
}

func (s *Server) Start() error {
}

func (s *Server) Stop() error {
}
