package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() {
	err := godotenv.Load()

	if err != nil{
		fmt.Println("Error loading .env file:", err)
	}

	user := os.Getenv("MYSQL_USER")

	pass:= os.Getenv("MYSQL_PASSWORD")

	database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf(

		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,pass,database,
	)


	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		fmt.Println("err")
	}
	DB.AutoMigrate(&User{})


}