package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendFile(c *gin.Context) {
	folder := c.Param("folder")
	path := fmt.Sprintf("public/images%s", folder)
	c.File(path)
}
