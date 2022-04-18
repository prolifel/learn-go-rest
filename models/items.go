package models

type Items struct {
	Item_id     uint   `gorm:"primary_key"`
	Item_code   string `gorm:"type:varchar(20);not null"`
	Description string `gorm:"type:varchar(200)"`
	Quantity    uint
	Order_id    uint
	Order       Orders `gorm:"foreignkey:Order_id"`
}
