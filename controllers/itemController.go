package controllers

import (
	"net/http"
	"order-restapi/models"

	"github.com/gin-gonic/gin"
)

// Get all items
func (databaseConnection *DatabaseConnection) GetItem(c *gin.Context) {
	var (
		item   []models.Items
		result gin.H
	)

	databaseConnection.DB.Find(&item)
	if len(item) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": item,
			"count":  len(item),
		}
	}

	c.JSON(http.StatusOK, result)
}
