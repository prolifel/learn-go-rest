package main

import (
	"order-restapi/config"
	"order-restapi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	databaseConnection := &controllers.DatabaseConnection{DB: db}

	router := gin.Default()

	router.GET("/orders", databaseConnection.GetItem)
	router.Run(":3000")
}
