package models

type GenoType struct {
	ProfileId   string `json:"profile_id" gorm:"column:profile_id"`
	ReportId    int    `json:"report_id" gorm:"column:report_id"`
	Type        string `json:"type" gorm:"column:type"`
	Genotype    string `json:"genotype" gorm:"column:genotype"`
	Summary     string `json:"summary" gorm:"column:summary"`
	TSummary    string `json:"tsummary" gorm:"column:tsummary"`
	Score       string `json:"score" gorm:"column:score"`
	Rsid        string `json:"rsid" gorm:"column:rsid"`
	Gene        string `json:"gene" gorm:"column:gene"`
	OrValue     string `json:"or_value,omitempty" gorm:"column:or_value"`
	Orientation string `json:"orientation,omitempty" gorm:"column:orientation"`
}
