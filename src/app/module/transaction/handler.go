package transaction

import (
	"funding/src/app/common"
	"funding/src/app/helper"
	"funding/src/app/module/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewTransactionHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) AddTransaction(c *gin.Context) {
	var input InputTransaction
	commons := common.NewCommon(c)
	userData := c.MustGet("currentUser").(user.User)
	err := c.ShouldBindJSON(&input)
	if err != nil && commons.ErrorHandler("Validation Error", http.StatusBadRequest, helper.FormatValidationError(err)) {
		return
	}
	input.UserId = userData.Id
	data, err := h.service.Add(input)
	if err != nil && commons.ErrorHandler("Error DB", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	formated := FormatTransaction(data)
	response := helper.ResponseHelper("Transaction Saved", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetTransactionById(c *gin.Context) {
	commons := common.NewCommon(c)
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil && commons.ErrorHandler("Numbers only", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	data, err := h.service.GetById(userId)
	if err != nil && commons.ErrorHandler("Error DB", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *handler) GetListTransaction(c *gin.Context) {
	commons := common.NewCommon(c)
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id
	data, err := h.service.GetByUser(userId)
	if err != nil && commons.ErrorHandler("Error DB", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	formated := FormatListTransaction(data)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetTransactionByCampaignId(c *gin.Context) {
	commons := common.NewCommon(c)
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil && commons.ErrorHandler("Parse Error", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	data, err := h.service.GetByCampaignId(id, userId)
	if err != nil && commons.ErrorHandler("Error DB", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	formated := FormatListTransaction(data)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}
