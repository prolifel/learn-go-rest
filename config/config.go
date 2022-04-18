package config

import (
	"fmt"
	"order-restapi/models"

	"github.com/jinzhu/gorm"
)

const (
	username = "root"
	psw      = "clement"
	db       = "orders_by"
)

var conn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, psw, db)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", conn)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(models.Items{}, models.Orders{})
	return db
}
