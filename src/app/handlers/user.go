package handler

import (
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
	err := c.ShouldBindJSON(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ResponseHelper("Data berhasil disimpan", 200, "sukses", user)
		c.JSON(http.StatusBadRequest, response)
	}
	response := helper.ResponseHelper("Data berhasil disimpan", 200, "sukses", user)
	c.JSON(http.StatusOK, response)
}

// func (h *userHandler) GetUser(c *gin.Context) {

// }
