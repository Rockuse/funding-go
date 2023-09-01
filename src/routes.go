package server

import (
	"funding/src/routes"
)

func ConfigureRoutes(server *Server) {
	var module = []routes.Module{
		routes.UserModule,
		routes.CampaignModule,
	}
	db := server.db
	api := server.engine.Group("/api/v1")
	for _, m := range module {
		m.Routes(api, db)
	}
}
