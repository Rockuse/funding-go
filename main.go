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
	// user := user.User{
	// 	Id:   config.Uuid(),
	// 	Name: "test",
	// }
	// userRepository.Save(user)
	users, err := userRepository.GetAll()
	fmt.Println(users)
	fmt.Println("Connection is good")
	// router := gin.Default()
	// router.GET("/handler", user.UserHandler)
	// router.Run()
}
