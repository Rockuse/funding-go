package server

import (
	handler "funding/src/app/handlers"
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
		routes.TransactionModule,
	}
	router := s.engine
	router.Static("/images", "./public/images")
	db := s.db
	api := router.Group("/api/v1")
	api.GET("/images/:folder", handler.SendFile)
	for _, m := range module {
		m.Routes(api, db)
	} 
}

func (s *Server) Run() error {
	// gin.SetMode(gin.ReleaseMode)
	return s.engine.Run()
}
