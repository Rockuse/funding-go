package middleware

import (
	"funding/src/app/auth"
	"funding/src/app/helper"
	"funding/src/app/module/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authMiddleware struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *authMiddleware {
	return &authMiddleware{db}
}

func (a *authMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authService := auth.NewService()
		userRepository := user.NewRepository(a.db)
		userService := user.NewService(userRepository)
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
