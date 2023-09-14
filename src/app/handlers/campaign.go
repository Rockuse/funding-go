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
	host := c.Request.URL.Host
	campaignArr, err := h.campaignService.FindAll()
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error DB", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formated := campaign.FormatAllCampaigns(campaignArr, host)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetListCampaignById(c *gin.Context) {
	host := c.Request.URL.Host
	// var list []campaign.CampaignFormat
	// currentUser := c.MustGet("currentUser").(user.User)
	// userId := currentUser.Id
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error Param", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignArr, err := h.campaignService.FindByUserId(userId)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error DB", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formated := campaign.FormatAllCampaigns(campaignArr, host)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) SaveCampaign(c *gin.Context) {
	host := c.Request.URL.Host
	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.CampaignInput
	err := c.ShouldBindJSON(&input)
	input.UserId = currentUser.Id
	input.CreatedBy = strconv.Itoa(currentUser.Id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("Binding Failed", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.campaignService.SaveCampaign(input)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error DB", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formated := campaign.FormatCampaign(data, host)
	response := helper.ResponseHelper("Campaign Saved", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}
func (h *handler) UpdateCampaign(c *gin.Context) {
	host := c.Request.URL.Host
	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.CampaignInput
	err := c.ShouldBindJSON(&input)
	input.UserId = currentUser.Id
	input.CreatedBy = strconv.Itoa(currentUser.Id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ResponseHelper("Binding Failed", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.campaignService.SaveCampaign(input)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error DB", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formated := campaign.FormatCampaign(data, host)
	response := helper.ResponseHelper("Campaign Saved", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetDetail(c *gin.Context) {
	var input campaign.CampaignUri
	commons := &common.MyContext{Context: c}
	err := c.ShouldBindUri(&input)
	if !commons.ErrorHandler("Error Input", http.StatusBadRequest, err) {
		return
	}
	campaignData, err := h.campaignService.GetCampaignById(input)
	if !commons.ErrorHandler(err.Error(), http.StatusBadRequest, err) {
		return
	}
	host := c.Request.URL.Host
	formated := campaign.FormatDetail(campaignData, host)
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}
