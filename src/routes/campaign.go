package routes

import (
	"funding/src/app/campaign"
	handler "funding/src/app/handlers"
	"funding/src/app/user"
	"funding/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var CampaignModule = Module{
	Name: "Campaign",
	Routes: func(api *gin.RouterGroup, db *gorm.DB) {
		midService := middleware.NewService(db)
		userRepository := user.NewRepository(db)
		userService := user.NewService(userRepository)
		campaignRepository := campaign.NewRepository(db)
		campaignService := campaign.NewService(campaignRepository, userService)
		campaignHandler := handler.NewCampaignHandler(campaignService)

		campaignApi := api.Group("/campaign")
		campaignApi.GET("/", campaignHandler.GetListCampaign)                                     // Get All Campaign
		campaignApi.GET("/:id", midService.AuthMiddleware(), campaignHandler.GetListCampaignById) // get Campaign by User ID
		campaignApi.GET("/detail/:id", midService.AuthMiddleware(), campaignHandler.GetDetail)    // Get Campaign Detail by ID
		campaignApi.POST("/", midService.AuthMiddleware(), campaignHandler.SaveCampaign)          // Add Campaign
		campaignApi.PUT("/:id", midService.AuthMiddleware(), campaignHandler.UpdateCampaign)      // Edit Campaign
		campaignApi.POST("/images", midService.AuthMiddleware(), campaignHandler.SaveCampaign)    // Upload Image
	},
}
