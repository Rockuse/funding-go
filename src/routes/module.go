package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Name   string
	Routes func(*gin.RouterGroup, *gorm.DB)
}

func RegisteredModules() []Module {
	return []Module{TransactionModule, CampaignModule, UserModule}
}
