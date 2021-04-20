package global

import "fmt"

type Server struct {
	Host    string
	Port    int
	Version string
	Env     string
}

var GlbServer Server

func init() {
	GlbServer = Server{
		Host:    "0.0.0.0",
		Port:    8080,
		Version: "v0.0.1",
		Env:     "debug",
	}
}

func (s *Server) AddressString() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
