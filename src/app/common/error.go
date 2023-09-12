package common

import (
	"funding/src/app/helper"

	"github.com/gin-gonic/gin"
)

type MyContext struct {
	*gin.Context
}

func (c *MyContext) ErrorHandler(text string, status int, err error) bool {
	if err != nil {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper(text, status, "fail", errors)
		c.JSON(status, response)
		return false
	}
	return true
}
