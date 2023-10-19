package user

import (
	"bytes"
	"fmt"
	"funding/src/app/auth"
	"funding/src/app/helper"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Service
	authService auth.Service
}

func NewUserHandler(userService Service) *userHandler {
	return &userHandler{userService, auth.NewService()}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	check := c.MustGet("email_isAvailable")
	var input RegisterInput
	err := c.ShouldBindJSON(&input)
	fmt.Println(check)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("Data Gagal disimpan", http.StatusUnprocessableEntity, "Fail", errors)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("Data Gagal disimpan", http.StatusBadRequest, "fail", errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.Id)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Data gagal disimpan", http.StatusOK, "success", errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	formater := FormatUser(newUser, token)
	response := helper.ResponseHelper("Data berhasil disimpan", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	ByteBody, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(ByteBody))
	var email EmailInput
	err := c.ShouldBindJSON(&email)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("format Email salah", http.StatusUnprocessableEntity, "fail", errors)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	isValid, err := h.userService.IsEmailAvailable(email)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Email gagal di cek", http.StatusBadRequest, "fail", errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !isValid {
		errors := gin.H{"errors": "Email sudah terdaftar"}
		response := helper.ResponseHelper("Email sudah terdaftar", http.StatusBadRequest, "fail", errors)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(ByteBody))
	c.Set("email_isAvailable", isValid)
	c.Next()
}

func (h *userHandler) Login(c *gin.Context) {
	var input LoginInput
	err := c.ShouldBindJSON(&input)
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
	token, err := h.authService.GenerateToken(userLogin.Id)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Data gagal disimpan", http.StatusOK, "success", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := FormatUser(userLogin, token)
	response := helper.ResponseHelper("Berhasil Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	//User input
	//input ditangkap handler
	//mapping dari input user ke "input struct login"
	//"input struct login" passing ke user service
	//service mencari data user menggunakan repository (user,email)
	//cocokan password
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		errors := gin.H{"is_uploaded": false}
		response := helper.ResponseHelper("Gagal upload avatar", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//dapat dari JWT
	currentUser := c.MustGet("currentUser").(User)
	userID := currentUser.Id
	newPath, pathName := helper.PathUpload("user", strconv.Itoa(userID), file.Filename)
	err = c.SaveUploadedFile(file, newPath)
	if err != nil {
		errors := gin.H{"is_uploaded": false}
		response := helper.ResponseHelper("Gagal upload avatar", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UpdateAvatar(userID, pathName)
	if err != nil {
		errors := gin.H{"is_uploaded": false}
		response := helper.ResponseHelper("Gagal upload avatar", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	c.JSON(http.StatusOK, data)
	// upload user foto
	// simpan gambar di folder /images
	// di service panggil repo
	// tambahkan JWT
	// repo ambil data user  id=1
	// repo update lokasi avatar user id=1
}

func (h *userHandler) GetUserDataById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
	}
	userData, err := h.userService.GetUserById(id)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formater := FormatUser(userData, "tes")
	response := helper.ResponseHelper("Get User", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, response)
}
