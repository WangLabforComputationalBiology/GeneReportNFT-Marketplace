package models

import "time"

type GNFT struct {
	Id            string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	CollectionId  string    `gorm:"uniqueIndex:idx_gnft;type:varchar(32)" json:"collection_id"`
	Identifier    string    `gorm:"uniqueIndex:idx_gnft;type:varchar(32)" json:"identifier"`
	Category      string    `gorm:"type:varchar(32)" json:"category"`
	Address       string    `gorm:"type:varchar(42)" json:"address"`
	TokenStandard string    `gorm:"type:varchar(32)" json:"token_standard"`
	Name          string    `gorm:"index:idx_name;type:varchar(32)" json:"name"`
	Description   string    `gorm:"type:text" json:"description"`
	Supply        int       `gorm:"type:int" json:"supply"`
	IsMinted      bool      `gorm:"type:boolean;default:false" json:"is_minted"`
	CreatedAt     time.Time `gorm:"type:datetime" json:"created_at"`
}
