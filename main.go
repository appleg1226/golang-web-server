package main

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"awesomeProject/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status: ", err)
	}

	err = config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Use(gin.Logger())

	err = r.Run()
}
