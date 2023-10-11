package common

import (
	"fmt"
	"funding/src/app/helper"

	"github.com/gin-gonic/gin"
)

type MyContext struct {
	*gin.Context
}

func (c *MyContext) ErrorHandler(text string, status int, err gin.H) bool {
	errors := err["errors"]
	fmt.Println(errors)
	if errors != nil {
		response := helper.ResponseHelper(text, status, "fail", errors)
		c.Context.JSON(status, response)
		return true
	}
	return false
}
