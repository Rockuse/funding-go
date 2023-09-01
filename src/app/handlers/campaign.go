package handler

import (
	"funding/src/app/campaign"
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
	var list []campaign.CampaignFormat
	campaignList, err := h.campaignService.FindAll()
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error DB", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, data := range campaignList {
		list = append(list, campaign.FormatCampaign(data))
	}
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", list)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetListCampaignById(c *gin.Context) {
	var list []campaign.CampaignFormat
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id
	campaignList, err := h.campaignService.FindByUserId(userId)
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper("Error DB", http.StatusBadRequest, "fail", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	for _, data := range campaignList {
		list = append(list, campaign.FormatCampaign(data))
	}
	response := helper.ResponseHelper("Data berhasil ditampilkan", http.StatusOK, "success", list)
	c.JSON(http.StatusOK, response)
}

func (h *handler) SaveCampaign(c *gin.Context) {
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
	formated := campaign.FormatCampaign(data)
	response := helper.ResponseHelper("Campaign Saved", http.StatusOK, "success", formated)
	c.JSON(http.StatusOK, response)
}
