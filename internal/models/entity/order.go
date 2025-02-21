package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	Status    boolean   `gorm:"type:boolean" json:"order_status"`
	BuyerId   string    `gorm:"type:varchar(32)" json:"buyer_id"`
	SellerId  string    `gorm:"type:varchar(32)" json:"seller_id"`
}
