package user

import (
	"bytes"
	"funding/src/app/auth"
	"funding/src/app/common"
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
	var input RegisterInput
	commons := common.NewCommon(c)
	err := c.ShouldBindJSON(&input)
	if err != nil && commons.ErrorHandler("Data Gagal disimpan", http.StatusUnprocessableEntity, helper.FormatValidationError(err)) {
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil && commons.ErrorHandler("Data Gagal disimpan", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	formater := FormatUser(newUser)
	token, err := h.authService.GenerateToken(formater)
	if err != nil && commons.ErrorHandler("Data Gagal disimpan", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	tokenFormated := auth.TokenFormater(token)
	response := helper.ResponseHelper("Data berhasil disimpan", http.StatusOK, "success", tokenFormated)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	commons := common.NewCommon(c)
	ByteBody, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(ByteBody))
	var email EmailInput
	err := c.ShouldBindJSON(&email)
	if err != nil && commons.ErrorHandler("format email salah", http.StatusUnprocessableEntity, helper.FormatValidationError(err)) {
		return
	}
	isValid, err := h.userService.IsEmailAvailable(email)
	if err != nil && commons.ErrorHandler("Email gagal di cek", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	if !isValid && commons.ErrorHandler("Email sudah terdaftar", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(ByteBody))
	c.Set("email_isAvailable", isValid)
	c.Next()
}

func (h *userHandler) Login(c *gin.Context) {
	commons := common.NewCommon(c)
	var input LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil && commons.ErrorHandler("Email gagal di cek", http.StatusUnprocessableEntity, helper.FormatValidationError(err)) {
		return
	}
	userLogin, err := h.userService.Login(input)

	if err != nil && commons.ErrorHandler("Email/Password salah", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	formatter := FormatUser(userLogin)
	token, err := h.authService.GenerateToken(formatter)
	if err != nil && commons.ErrorHandler("Data gagal disimpan", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	tokenFormated := auth.TokenFormater(token)
	response := helper.ResponseHelper("Berhasil Login", http.StatusOK, "success", tokenFormated)

	c.JSON(http.StatusOK, response)
	//User input
	//input ditangkap handler
	//mapping dari input user ke "input struct login"
	//"input struct login" passing ke user service
	//service mencari data user menggunakan repository (user,email)
	//cocokan password
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	commons := common.NewCommon(c)
	file, err := c.FormFile("avatar")
	if err != nil && commons.ErrorHandler("Gagal upload avatar", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	//dapat dari JWT
	currentUser := c.MustGet("currentUser").(User)
	userID := currentUser.Id
	newPath, pathName := helper.PathUpload("user", strconv.Itoa(userID), file.Filename)
	err = c.SaveUploadedFile(file, newPath)
	if err != nil && commons.ErrorHandler("Gagal upload avatar", http.StatusBadRequest, commons.Error(err)) {
		return
	}

	_, err = h.userService.UpdateAvatar(userID, pathName)
	if err != nil && commons.ErrorHandler("Gagal upload avatar", http.StatusBadRequest, commons.Error(err)) {
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
	commons := common.NewCommon(c)
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil && commons.ErrorHandler("Error", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	userData, err := h.userService.GetUserById(id)
	if err != nil && commons.ErrorHandler("Error", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	formater := FormatUser(userData)
	response := helper.ResponseHelper("Get User", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, response)
}
