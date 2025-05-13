package models

type Ownership struct {
	Owner        string `gorm:"index:idx_ownership;type:varchar(42)" json:"owner"`
	CollectionId string `gorm:"index:idx_ownership;type:varchar(32)" json:"collection_id"`
	Identifier   string `gorm:"index:idx_ownership;type:varchar(32)" json:"identifier"`
	Quantity     int    `gorm:"type:int" json:"quantity"`
}
