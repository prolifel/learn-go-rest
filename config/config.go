package config

import (
	"fmt"
	"order-restapi/controllers"
	"order-restapi/models"
	"time"

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

func Seeder() {
	db := DBInit()
	databaseConnection := &controllers.DatabaseConnection{DB: db}

	databaseConnection.DB.Create(&models.Orders{
		Order_id:      1,
		Customer_name: "Clement",
		Ordered_at:    time.Now(),
		Items: []models.Items{
			{
				Item_id:     1,
				Item_code:   "A",
				Description: "Apple",
				Quantity:    1,
			},
			{
				Item_id:     2,
				Item_code:   "B",
				Description: "Banana",
				Quantity:    2,
			},
		},
	})
	databaseConnection.DB.Create(&models.Orders{
		Order_id:      2,
		Customer_name: "Denta",
		Ordered_at:    time.Now(),
		Items: []models.Items{
			{
				Item_id:     3,
				Item_code:   "C",
				Description: "Cherry",
				Quantity:    3,
			},
			{
				Item_id:     4,
				Item_code:   "D",
				Description: "Durian",
				Quantity:    4,
			},
		},
	})
	databaseConnection.DB.Create(&models.Orders{
		Order_id:      3,
		Customer_name: "Angga",
		Ordered_at:    time.Now(),
		Items: []models.Items{
			{
				Item_id:     5,
				Item_code:   "E",
				Description: "Eggplant",
				Quantity:    5,
			},
			{
				Item_id:     6,
				Item_code:   "F",
				Description: "Fig",
				Quantity:    6,
			},
		},
	})
}
