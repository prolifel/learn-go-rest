package controllers

import (
	"net/http"
	"order-restapi/models"

	"github.com/gin-gonic/gin"
)

func (databaseConnection *DatabaseConnection) GetOrder(c *gin.Context) {
	var (
		order  []models.Orders
		result gin.H
	)

	// databaseConnection.DB.Find(&order)
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
