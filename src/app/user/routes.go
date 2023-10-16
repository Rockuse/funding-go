package user

import (
	"funding/src/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var UserModule = routes.Module{
	Name: "User",
	Routes: func(api *gin.RouterGroup, db *gorm.DB) {
		userRepository := NewRepository(db)
		userService := NewService(userRepository)
		userHandler := NewUserHandler(userService)

		userApi := api.Group("/user")
		userApi.POST("", userHandler.CheckEmailAvailibility, userHandler.RegisterUser)
		userApi.GET("/:id", userHandler.GetUserDataById)
		userApi.POST("/login", userHandler.Login)
		userApi.POST("/checkemail", userHandler.CheckEmailAvailibility)
	},
}
