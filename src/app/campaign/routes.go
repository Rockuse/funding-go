package campaign

import (
	"funding/src/app/user"
	"funding/src/middleware"
	"funding/src/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var CampaignModule = routes.Module{
	Name: "Campaign",
	Routes: func(api *gin.RouterGroup, db *gorm.DB) {
		midService := middleware.NewService(db)
		userRepository := user.NewRepository(db)
		userService := user.NewService(userRepository)
		campaignRepository := NewRepository(db)
		campaignService := NewService(campaignRepository, userService)
		campaignHandler := NewCampaignHandler(campaignService)

		campaignApi := api.Group("/campaign")
		campaignApi.GET("/", campaignHandler.GetListCampaign)                                     // Get All Campaign
		campaignApi.GET("/:id", midService.AuthMiddleware(), campaignHandler.GetListCampaignById) // get Campaign by User ID
		campaignApi.GET("/detail/:id", midService.AuthMiddleware(), campaignHandler.GetDetail)    // Get Campaign Detail by ID
		campaignApi.POST("/", midService.AuthMiddleware(), campaignHandler.SaveCampaign)          // Add Campaign
		campaignApi.PUT("/:id", midService.AuthMiddleware(), campaignHandler.UpdateCampaign)      // Edit Campaign
		campaignApi.POST("/images", midService.AuthMiddleware(), campaignHandler.SaveCampaign)    // Upload Image
	},
}

// Module{

// }
