package controllers

import (
	"fmt"
	"net/http"
	"order-restapi/models"

	"github.com/gin-gonic/gin"
)

func (databaseConnection *DatabaseConnection) GetOrders(c *gin.Context) {
	var (
		order  []models.Orders
		result gin.H
	)

	databaseConnection.DB.Preload("Items").Find(&order)
	if len(order) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": order,
			"count":  len(order),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (databaseConnection *DatabaseConnection) CreateOrder(c *gin.Context) {
	var (
		order  models.Orders
		result gin.H
	)

	if err := c.BindJSON(&order); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "form data error"})
		return
	}

	// fmt.Printf("%+v\n", order)

	databaseConnection.DB.Create(&order)
	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)
}
