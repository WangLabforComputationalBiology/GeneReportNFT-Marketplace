package models

type GNFT struct {
	ProfileId     string `gorm:"index:idx_gnft;type:varchar(32)" json:"profileId"`
	Identifier    string `gorm:"index:idx_gnft;type:varchar(32)" json:"identifier"`
	Category      string `gorm:"type:varchar(32)" json:"category"`
	Address       string `gorm:"type:varchar(42)" json:"address"`
	TokenStandard string `gorm:"type:varchar(32)" json:"token_standard"`
	Name          string `gorm:"index:idx_name;type:varchar(32)" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	Supply        int    `gorm:"type:int" json:"supply"`
	IsMinted      bool   `gorm:"type:boolean;default:false" json:"is_minted"`
}
