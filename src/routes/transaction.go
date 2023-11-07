package routes

import (
	"funding/src/app/module/transaction"
	"funding/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var TransactionModule = Module{
	Name: "Transaction",
	Routes: func(api *gin.RouterGroup, db *gorm.DB) {
		midService := middleware.NewService(db)
		transactionRepository := transaction.NewRepository(db)
		transactionService := transaction.NewService(transactionRepository)
		transactionHandler := transaction.NewTransactionHandler(transactionService)

		transactionApi := api.Group("/transaction")
		transactionApi.POST("", midService.AuthMiddleware(), transactionHandler.AddTransaction)
		transactionApi.GET("", midService.AuthMiddleware(), transactionHandler.GetListTransaction)
		transactionApi.GET("/:id", midService.AuthMiddleware(), transactionHandler.GetTransactionById)
		transactionApi.GET("/:id/transaction", midService.AuthMiddleware(), transactionHandler.GetTransactionByCampaignId)
	},
}
