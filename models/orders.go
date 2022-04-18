package models

import "time"

type Orders struct {
	Order_id      uint      `gorm:"primary_key"`
	Customer_name string    `gorm:"type:varchar(100);not null"`
	Ordered_at    time.Time `gorm:"type:timestamp"`
	Items         []Items   `gorm:"foreignkey:Order_id"`
}
