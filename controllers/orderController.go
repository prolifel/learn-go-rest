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

func (databaseConnection *DatabaseConnection) UpdateOrder(c *gin.Context) {
	id := c.Query("id")

	var (
		order  models.Orders
		result gin.H
	)

	if err := databaseConnection.DB.First(&order, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if err := c.BindJSON(&order); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "form data error"})
		return
	}

	if err := databaseConnection.DB.Model(&order).Updates(&order).Error; err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": order,
			"status": "update success",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (databaseConnection *DatabaseConnection) DeleteOrder(c *gin.Context) {
	id := c.Query("id")

	var (
		order  models.Orders
		result gin.H
	)

	if err := databaseConnection.DB.First(&order, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if err := databaseConnection.DB.Model(&order).Delete(&order).Error; err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"status": "delete success",
		}
	}

	c.JSON(http.StatusOK, result)
}
