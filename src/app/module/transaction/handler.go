package transaction

import (
	"funding/src/app/common"
	"funding/src/app/helper"
	"funding/src/app/module/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func (s *handler) NewTransactionHandler(service Service) *handler {
	return &handler{service}
}

func (s *handler) AddTransaction(c *gin.Context) {
	var input InputTransaction
	commons := common.NewCommon(c)
	userData := c.MustGet("currentUser").(user.User)
	err := c.ShouldBindJSON(&input)
	if err != nil && commons.ErrorHandler("Validation Error", http.StatusBadRequest, helper.FormatValidationError(err)) {
		return
	}
	input.UserId = userData.Id
	data, err := s.service.Add(input)
	if err != nil && commons.ErrorHandler("Error DB", http.StatusBadRequest, commons.Error(err)) {
		return
	}
	c.JSON(http.StatusOK, data)
}
