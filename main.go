package main

import (
	"fmt"
	server "funding/src"
	"funding/src/db"
	"log"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	db := db.InitDB()
	app := server.NewServer(db) //*gorm.DB
	app.ConfigureRoutes()
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
	fmt.Println("Connection is good")
}
