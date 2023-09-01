package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	engine *gin.Engine
	db     *gorm.DB
}

func NewServer(dbConnection *gorm.DB) *Server {
	return &Server{
		engine: gin.Default(),
		db:     dbConnection,
	}
}

func (s *Server) Run() error {
	// gin.SetMode(gin.ReleaseMode)
	return s.engine.Run()
}
