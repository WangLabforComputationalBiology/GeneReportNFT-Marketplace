package models

type Listing struct {
	Id                string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	CollectionId      string `gorm:"index:idx_listing;type:varchar(32)" json:"collection_id"`
	Identifier        string `gorm:"index:idx_listing;type:varchar(32)" json:"identifier"`
	SaleAmount        uint   `gorm:"index:idx_listing;type:int" json:"sale_amount"`
	Price             uint   `gorm:"type:int" json:"price"`
	Offerer           string `gorm:"type:varchar(42)" json:"offerer"`
	CreatorFee        uint   `gorm:"type:bigint" json:"creator_fee"`
	Creator           string `gorm:"type:varchar(42)" json:"creator"`
	OffererFee        uint   `gorm:"type:bigint" json:"offerer_fee"`
	Salt              string `gorm:"type:varchar(32)" json:"salt"`
	Signature         string `gorm:"type:varchar(132)" json:"signature"`
	RemainingQuantity int    `gorm:"type:int" json:"remaining_quantity"`
	Finalized         bool   `gorm:"type:bool;default:false" json:"finalized"`
}
