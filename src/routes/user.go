package routes

import (
	handler "funding/src/app/handlers"
	"funding/src/app/user"
	"funding/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var UserModule = Module{
	Name: "User",
	Routes: func(api *gin.RouterGroup, db *gorm.DB) {
		midService := middleware.NewService(db)
		userRepository := user.NewRepository(db)
		userService := user.NewService(userRepository)
		userHandler := handler.NewUserHandler(userService)

		userApi := api.Group("/user")
		userApi.POST("/user", userHandler.RegisterUser)
		userApi.GET("/user/:id", userHandler.GetUserDataById)
		userApi.POST("/login", userHandler.Login)
		userApi.POST("/checkemail", userHandler.CheckEmailAvailibility)
		userApi.POST("/avatar", midService.AuthMiddleware(), userHandler.UploadAvatar)
	},
}
