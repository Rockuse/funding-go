package routes

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	RouterIndex()
}

type router struct {
	engine gin.Engine
}

func NewRouter() *router {
	return &router{engine: gin.Default()}
}

// func (r *router) RouterIndex(handler handler) *gin.Engine {
// 	var Router *gin.Engine
// 	Router = gin.Default()
// 	UserRouter(Router)
// 	return Router
// }

// func UserRouter(router gin.Engine) {
// 	api := router.Group("/api/v1")
// 	api.POST("/user", userHandler.RegisterUser)
// 	api.POST("/login", userHandler.RegisterUser)
// }
