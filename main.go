package main

import (
	"fmt"
	"funding/src/app/auth"
	"funding/src/app/campaign"
	handler "funding/src/app/handlers"
	"funding/src/app/helper"
	"funding/src/app/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	dns := "host=satao.db.elephantsql.com user=sxelyaew password=3r9FiVZVfvCtli4eR4ZiL9bb0KYTItW_ dbname=sxelyaew port=5432"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	authService := auth.NewService()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService, authService)

	router := gin.Default()
	if err != nil {
		fmt.Println(err)
	}
	// input := user.LoginInput{Email: "fahmi.muza@gmail.com", Password: "Passivea"}
	// userData, err := user.Service.Login(userService, input)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(userData, "tes")

	api := router.Group("/api/v1")
	api.POST("/user", userHandler.RegisterUser)
	api.GET("/user/:id", userHandler.GetUserDataById)
	api.POST("/login", userHandler.Login)
	api.POST("/checkemail", userHandler.CheckEmailAvailibility)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaign", authMiddleware(authService, userService), campaignHandler.GetListCampaign)
	api.POST("/campaign", authMiddleware(authService, userService), campaignHandler.SaveCampaign)
	router.Run()
	fmt.Println("Connection is good")
	// router := gin.Default()
	// router.GET("/handler", user.UserHandler)
	// router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponseHelper("Bearer Not Found", http.StatusUnauthorized, "fail", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ResponseHelper("Unauthorized 1", http.StatusUnauthorized, "fail", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ResponseHelper("Unauthorized 2", http.StatusUnauthorized, "fail", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId := int(claim["user_id"].(float64))
		userData, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.ResponseHelper("There is error", http.StatusNotFound, "fail", err)
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}
		if userData.Id == 0 {
			response := helper.ResponseHelper("user not found", http.StatusNotFound, "fail", err)
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}
		c.Set("currentUser", userData)
	}
}
