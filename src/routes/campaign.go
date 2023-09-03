package routes

import (
	"funding/src/app/campaign"
	handler "funding/src/app/handlers"
	"funding/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var CampaignModule = Module{
	Name: "Campaign",
	Routes: func(api *gin.RouterGroup, db *gorm.DB) {
		midService := middleware.NewService(db)
		campaignRepository := campaign.NewRepository(db)
		campaignService := campaign.NewService(campaignRepository)
		campaignHandler := handler.NewCampaignHandler(campaignService)

		campaignApi := api.Group("/campaign")
		campaignApi.GET("/", midService.AuthMiddleware(), campaignHandler.GetListCampaign)
		campaignApi.GET("/detail", midService.AuthMiddleware(), campaignHandler.GetListCampaignById)
		campaignApi.POST("/", midService.AuthMiddleware(), campaignHandler.SaveCampaign)
	},
}
