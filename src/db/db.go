package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=satao.db.elephantsql.com user=sxelyaew password=3r9FiVZVfvCtli4eR4ZiL9bb0KYTItW_ dbname=sxelyaew port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
