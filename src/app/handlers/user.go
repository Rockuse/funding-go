package handler

import (
	"fmt"
	"funding/src/app/helper"
	"funding/src/app/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
	}
	newUser, err := h.userService.RegisterUser(input)
	fmt.Println(newUser)
	if err != nil {
		response := helper.ResponseHelper("Data Gagal disimpan", 400, "Fail", newUser)
		fmt.Println(response)
		c.JSON(http.StatusBadRequest, response)
	}
	// token, err := h.jwtService.GenerateToken()
	formater := user.FormatUser(newUser, "token")
	response := helper.ResponseHelper("Data berhasil disimpan", 200, "sukses", formater)
	c.JSON(http.StatusOK, response)
}

// func (h *userHandler) GetAllUsers(c *gin.Context) {
// 	var users []user.User
// 	h.userService.GetAllUser(&users)
// }

// func (h *userHandler) GetUser(c *gin.Context) {

// }
