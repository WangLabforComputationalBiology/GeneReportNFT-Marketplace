package models

import (
	"time"
)

type Collection struct {
	Id             string    `gorm:"primaryKey;type:varchar(32)" json:"Id"`
	Name           string    `gorm:"type:varchar(32)" json:"name"`
	Address        string    `gorm:"type:varchar(42)" json:"address"`
	Description    string    `gorm:"type:text" json:"description"`
	Creator        string    `gorm:"type:varchar(42)" json:"creator"`
	CreatorEarning string    `gorm:"type:varchar(32)" json:"creator_earning"`
	RequiredZone   string    `gorm:"type:varchar(32)" json:"required_zone"`
	CreatedAt      time.Time `gorm:"type:datetime" json:"created_at"`
}
