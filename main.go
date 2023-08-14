package main

import (
	"fmt"
	user "funding/src/app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Connection is good")
	router := gin.Default()
	router.GET("/handler", user.UserHandler)
	router.Run()
}
