package handler

import (
	user "funding/src/app/user"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	// dsn := "host=satao.db.elephantsql.com user=sxelyaew password=3r9FiVZVfvCtli4eR4ZiL9bb0KYTItW_ dbname=sxelyaew port=5432"
	var users []user.User
	user, err := user.Repository.Get(&users)
	if err != nil {
		c.JSON(200, nil)
	}
	c.JSON(200, user)
}
