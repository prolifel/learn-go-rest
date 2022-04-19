package models

import "time"

type Orders struct {
	Order_id      uint      `gorm:"primary_key"`
	Customer_name string    `gorm:"type:varchar(100);not null" json:"customerName"`
	Ordered_at    time.Time `json:"orderedAt"`
	Items         []Items   `gorm:"foreignkey:Order_id;constraint:OnDelete:CASCADE"`
}
