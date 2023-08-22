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
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ResponseHelper("Data Gagal disimpan", http.StatusUnprocessableEntity, "Fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ResponseHelper("Data Gagal disimpan", http.StatusBadRequest, "Fail", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// token, err := h.jwtService.GenerateToken()
	formater := user.FormatUser(newUser, "token")
	response := helper.ResponseHelper("Data berhasil disimpan", http.StatusOK, "sukses", formater)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//User input
	//input ditangkap handler
	//mapping dari input user ke "input struct login"
	//"input struct login" passing ke user service
	//service mencari data user menggunakan repository (user,email)
	//cocokan password
}
