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
	user, err := h.userService.RegisterUser(input)
	fmt.Println(user)
	if err != nil {
		response := helper.ResponseHelper("Data Gagal disimpan", 400, "Fail", user)
		fmt.Println(response)
		c.JSON(http.StatusBadRequest, response)
	}
	response := helper.ResponseHelper("Data berhasil disimpan", 200, "sukses", user)
	c.JSON(http.StatusOK, response)
}

// func (h *userHandler) GetAllUsers(c *gin.Context) {
// 	var users []user.User
// 	h.userService.GetAllUser(&users)
// }

// func (h *userHandler) GetUser(c *gin.Context) {

// }
