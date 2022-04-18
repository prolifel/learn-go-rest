package controllers

import "github.com/jinzhu/gorm"

type DatabaseConnection struct {
	DB *gorm.DB
}
