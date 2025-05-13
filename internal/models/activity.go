package models

import "time"

type Item struct {
	CollectionId string `gorm:"type:varchar(32)" json:"collection_id"`
	Identifier   string `gorm:"type:varchar(32)" json:"identifier"`
}

type Activity struct {
	Id           string    `gorm:"type:varchar(36);not null;primaryKey" json:"id"`
	UserAddress  string    `gorm:"type:varchar(42);not null;index:idx_activity" json:"user_address"`
	Time         time.Time `gorm:"type:datetime;not null;index:idx_activity" json:"time"`
	ActivityType string    `gorm:"type:varchar(255);not null;" json:"activity_type"`
	Item
	Price    float64 `gorm:"type:decimal(20,8);not null;" json:"price"`
	Quantity int     `gorm:"type:int;not null;" json:"quantity"`
	From     string  `gorm:"type:varchar(42);not null;" json:"from"`
	To       string  `gorm:"type:varchar(42);not null;" json:"to"`
	Link     string  `gorm:"type:varchar(255);" json:"link"`
}
