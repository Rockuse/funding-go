package common

import (
	"fmt"
	"funding/src/app/helper"

	"github.com/gin-gonic/gin"
)

type myContext struct {
	*gin.Context
}

func NewCommon(c *gin.Context) *myContext {
	return &myContext{c}
}

func (c *myContext) ErrorHandler(text string, status int, err gin.H) bool {
	errors := err["errors"]
	fmt.Println(errors)
	if errors != nil {
		response := helper.ResponseHelper(text, status, "fail", errors)
		c.Context.AbortWithStatusJSON(status, response)
		return true
	}
	return false
}

func (c *myContext) Error(err error) gin.H {
	if err != nil {
		return gin.H{"errors": err.Error()}
	}
	return gin.H{"errors": err}
}
