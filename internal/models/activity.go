package models

import "time"

type Activity struct {
	Id              string    `gorm:"type:varchar(36);not null;primaryKey" json:"id"`
	UserAddress     string    `gorm:"type:varchar(42);not null;index:idx_activity" json:"user_address"`
	TransactionHash string    `gorm:"type:varchar(66);not null;index:idx_activity" json:"tx_hash"`
	Time            time.Time `gorm:"type:datetime;not null;index:idx_activity" json:"time"`
	Event           string    `gorm:"type:varchar(255);not null;" json:"event"`
	Expiry          time.Time `gorm:"type:datetime;not null;" json:"expiry"`
	GeneSharing     string    `gorm:"type:varchar(42);not null;" json:"geneSharing"`
	Metadata        string    `gorm:"type:varchar(42);not null;" json:"metadata"`
	From            string    `gorm:"type:varchar(42);not null;" json:"from"`
	To              string    `gorm:"type:varchar(42);not null;" json:"to"`
}
