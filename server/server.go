package server

import (
	"conceitoExato/env"

	"github.com/gin-gonic/gin"
)

type server struct {
	host, port   string
	serverEngine *gin.Engine
}

func CreateServer() *server {
	server := NewServer()

	server.SetPort(env.Server.PORT)
	server.SetHost(env.Server.HOST)
	return server
}

func NewServer() *server {
	newServer := &server{
		host:         "",
		port:         "",
		serverEngine: gin.Default(),
	}
	return newServer
}

func (s server) GetServerEngine() *gin.Engine {
	return s.serverEngine
}

func (s *server) SetPort(port string) {
	s.port = port
}

func (s *server) SetHost(host string) {
	s.host = host
}
