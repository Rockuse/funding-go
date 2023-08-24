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
		response := helper.ResponseHelper("Data Gagal disimpan", http.StatusUnprocessableEntity, "Fail", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("Data Gagal disimpan", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// token, err := h.jwtService.GenerateToken()
	formater := user.FormatUser(newUser, "token")
	response := helper.ResponseHelper("Data berhasil disimpan", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("Login gagal", http.StatusUnprocessableEntity, "fail", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userLogin, err := h.userService.Login(input)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Email/Password salah", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(userLogin, "testing")
	response := helper.ResponseHelper("Berhasil Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	//User input
	//input ditangkap handler
	//mapping dari input user ke "input struct login"
	//"input struct login" passing ke user service
	//service mencari data user menggunakan repository (user,email)
	//cocokan password
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	var email user.EmailInput
	err := c.ShouldBindJSON(&email)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("format Email salah", http.StatusUnprocessableEntity, "fail", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isValid, err := h.userService.IsEmailAvailable(email)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Email gagal di cek", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if !isValid {
		errors := gin.H{"errors": "Email sudah terdaftar"}
		response := helper.ResponseHelper("Email sudah terdaftar", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseHelper("Email tersedia", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	//user input email
	//input email ditangkap handler
	//input email mapping ke struct
	//validasi data binding email
	//panggil service check email
	//service panggil repo find email
	//apabila ada return error, apabila tidak ada sukses
}
