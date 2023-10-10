package common

import (
	"funding/src/app/helper"
	"reflect"

	"github.com/gin-gonic/gin"
)

type MyContext struct {
	*gin.Context
}

func (c *MyContext) ErrorHandler(text string, status int, err interface{}) bool {
	if err != nil && reflect.TypeOf(err).Name() == "string" {
		errors := gin.H{"errors": err}
		response := helper.ResponseHelper(text, status, "fail", errors)
		c.Context.JSON(status, response)
		return true
	} else if err != nil && reflect.TypeOf(err).Name() == "gin.H" {
		errors := err
		response := helper.ResponseHelper(text, status, "fail", errors)
		c.Context.JSON(status, response)
		return true
	}
	return false
}
