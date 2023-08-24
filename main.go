package main

import (
	"fmt"
	handler "funding/src/app/handlers"
	"funding/src/app/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dns := "host=satao.db.elephantsql.com user=sxelyaew password=3r9FiVZVfvCtli4eR4ZiL9bb0KYTItW_ dbname=sxelyaew port=5432"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()

	// input := user.LoginInput{Email: "fahmi.muza@gmail.com", Password: "Passivea"}
	// userData, err := user.Service.Login(userService, input)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(userData, "tes")

	api := router.Group("/api/v1")
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/checkemail", userHandler.CheckEmailAvailibility)
	// api.POST("/avatar", userHandler.CheckEmailAvailibility)
	router.Run()
	fmt.Println("Connection is good")
	// router := gin.Default()
	// router.GET("/handler", user.UserHandler)
	// router.Run()
}
