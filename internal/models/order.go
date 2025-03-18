package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	Status    bool      `gorm:"type:boolean" json:"order_status"`
	BuyerId   string    `gorm:"type:varchar(32)" json:"buyer_id"`
	SellerId  string    `gorm:"type:varchar(32)" json:"seller_id"`
	GNFTId    string    `gorm:"type:varchar(32)" json:"gnft_id"`
	Amount    int       `gorm:"type:int" json:"amount"`
}
