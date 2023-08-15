package main

import (
	"fmt"
	"funding/src/app/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dns := "host=satao.db.elephantsql.com user=sxelyaew password=3r9FiVZVfvCtli4eR4ZiL9bb0KYTItW_ dbname=sxelyaew port=5432"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterInput{}
	userInput.Name = "Nurul"
	userInput.Email = "nurularifah@gmail.com"
	userInput.Occupation = "Ibu Rumah Tangga"
	userInput.Password = "Jadi kie"
	userService.RegisterUser(userInput)
	fmt.Println("Connection is good")
	// router := gin.Default()
	// router.GET("/handler", user.UserHandler)
	// router.Run()
}
