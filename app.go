package main

import (
	"order-restapi/config"
	"order-restapi/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Seeding the database
func init() {
	config.Seeder()
}

func main() {
	db := config.DBInit()
	databaseConnection := &controllers.DatabaseConnection{DB: db}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "orders api served 🚀",
		})
	})
	router.GET("/orders", databaseConnection.GetOrders)
	router.POST("/orders", databaseConnection.CreateOrder)
	router.Run(":3000")
}
