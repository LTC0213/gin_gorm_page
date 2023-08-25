package main

import (
	"fmt"
	"gin_gorm_page/Config"
	models "gin_gorm_page/Model"
	routes "gin_gorm_page/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&models.User{})
	r := routes.SetUpRouter()

	r.Run(":8000")
}
