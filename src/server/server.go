package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := ConfigRoutes(s.server)
	log.Print("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
