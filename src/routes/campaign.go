package routes

import (
	"funding/src/app/module/campaign"
	"funding/src/app/module/user"
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
		campaignHandler := campaign.NewCampaignHandler(campaignService)

		campaignApi := api.Group("/campaign")
		campaignApi.GET("/", campaignHandler.GetListCampaign)                                     // Get All Campaign
		campaignApi.GET("/:id", midService.AuthMiddleware(), campaignHandler.GetListCampaignById) // get Campaign by User ID
		campaignApi.GET("/detail/:id", midService.AuthMiddleware(), campaignHandler.GetDetail)    // Get Campaign Detail by ID
		campaignApi.POST("/", midService.AuthMiddleware(), campaignHandler.SaveCampaign)          // Add Campaign
		campaignApi.PUT("/:id", midService.AuthMiddleware(), campaignHandler.UpdateCampaign)      // Edit Campaign
		campaignApi.POST("/image", midService.AuthMiddleware(), campaignHandler.SaveImages)       // Upload Image
	},
}
