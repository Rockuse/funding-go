package handler

import (
	"funding/src/app/campaign"
	"funding/src/app/common"
	"funding/src/app/helper"
	"funding/src/app/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *handler {
	return &handler{campaignService}
}

func (h *handler) GetListCampaign(c *gin.Context) {
	commons := &common.MyContext{Context: c}
	host := c.Request.URL.Host
	campaignArr, err := h.campaignService.FindAll()

	if commons.ErrorHandler("Error DB", http.StatusBadRequest, helper.Error(err)) {
		return
	}

	formated := campaign.FormatAllCampaigns(campaignArr, host)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetListCampaignById(c *gin.Context) {
	commons := &common.MyContext{Context: c}
	host := c.Request.URL.Host
	// var list []campaign.CampaignFormat
	// currentUser := c.MustGet("currentUser").(user.User)
	// userId := currentUser.Id
	id := c.Param("id")
	userId, err := strconv.Atoi(id)

	if commons.ErrorHandler("Error Param", http.StatusBadRequest, helper.Error(err)) {
		return
	}
	campaignArr, err := h.campaignService.FindByUserId(userId)

	if commons.ErrorHandler("Error DB", http.StatusBadRequest, helper.Error(err)) {
		return
	}
	formated := campaign.FormatAllCampaigns(campaignArr, host)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) SaveCampaign(c *gin.Context) {
	commons := &common.MyContext{Context: c}
	host := c.Request.URL.Host
	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.CampaignInput
	err := c.ShouldBindJSON(&input)
	input.UserId = currentUser.Id
	input.CreatedBy = strconv.Itoa(currentUser.Id)

	if err != nil && commons.ErrorHandler("Error Input", http.StatusBadRequest, helper.FormatValidationError(err)) {
		return
	}
	data, err := h.campaignService.SaveCampaign(input)

	if commons.ErrorHandler("Error DB", http.StatusBadRequest, helper.Error(err)) {
		return
	}

	formated := campaign.FormatCampaign(data, host)
	response := helper.ResponseHelper("Campaign Saved", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}
func (h *handler) UpdateCampaign(c *gin.Context) {
	commons := &common.MyContext{Context: c}
	host := c.Request.URL.Host
	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.CampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil && commons.ErrorHandler("Error Input", http.StatusBadRequest, helper.FormatValidationError(err)) {
		return
	}

	input.UserId = currentUser.Id
	input.CreatedBy = strconv.Itoa(currentUser.Id)
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	input.Id = intId
	if err != nil && commons.ErrorHandler("Convert Error", http.StatusBadRequest, helper.Error(err)) {
		return
	}

	data, err := h.campaignService.UpdateCampaign(input)
	if commons.ErrorHandler("Error DB", http.StatusBadRequest, helper.Error(err)) {
		return
	}
	

	formated := campaign.FormatCampaign(data, host)
	response := helper.ResponseHelper("Campaign Saved", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetDetail(c *gin.Context) {
	var input campaign.CampaignUri
	commons := &common.MyContext{Context: c}

	campaignData, err := h.campaignService.GetCampaignById(input)
	if commons.ErrorHandler(err.Error(), http.StatusBadRequest, helper.Error(err)) {
		return
	}
	host := c.Request.URL.Host
	formated := campaign.FormatDetail(campaignData, host)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}
