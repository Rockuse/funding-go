package routes

import (
	"funding/src/app/module/user"
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
		userHandler := user.NewUserHandler(userService)

		userApi := api.Group("/user")
		userApi.POST("", userHandler.CheckEmailAvailibility, userHandler.RegisterUser)
		userApi.GET("/:id", userHandler.GetUserDataById)
		userApi.POST("/login", userHandler.Login)
		userApi.POST("/checkemail", userHandler.CheckEmailAvailibility)
		userApi.POST("/avatar", midService.AuthMiddleware(), userHandler.UploadAvatar)
	},
}
