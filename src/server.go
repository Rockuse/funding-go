package server

import (
	"funding/src/routes"

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

func (s *Server) ConfigureRoutes() {
	var module = []routes.Module{
		routes.UserModule,
		routes.CampaignModule,
	}
	db := s.db
	api := s.engine.Group("/api/v1")
	for _, m := range module {
		m.Routes(api, db)
	}
}

func (s *Server) Run() error {
	// gin.SetMode(gin.ReleaseMode)
	return s.engine.Run()
}
