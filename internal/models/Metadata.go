package models

import (
	"time"
)

type Metadata struct {
	DataHash                   string    `gorm:"primaryKey;type:varchar(32)" json:"data_hash"`
	ProfileID                  string    `gorm:"type:varchar(36)" json:"profile_id"`
	Format                     string    `gorm:"type:varchar(32)" json:"format"`
	Sex                        bool      `gorm:"type:varchar(32)" json:"sex"`
	Category                   string    `gorm:"type:varchar(32)" json:"category"`
	Owner                      string    `gorm:"type:varchar(32)" json:"owner"`
	Name                       string    `gorm:"index:idx_name;type:varchar(32)" json:"name"`
	Description                string    `gorm:"type:text" json:"description"`
	ContractAddress            string    `gorm:"type:varchar(42)" json:"contract_address"`
	GeneSharingContractAddress string    `gorm:"type:varchar(42)" json:"gene_sharing_contract_address"`
	IsSharable                 bool      `gorm:"type:tinyint(1)" json:"is_sharable"`
	IsHidden                   bool      `gorm:"type:tinyint(1)" json:"is_hidden"`
	CreatedAt                  time.Time `gorm:"type:datetime" json:"created_at"`
}
