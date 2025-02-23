package entity

import (
	"time"
)

type GNFT struct {
	ID          int       `json:"id"`
	Name        string    `json:"file_name"`
	FileSig     string    `json:"file_sig"`
	FileShares  int       `json:"file_shares"`
	Data        string    `json:"data"`
	Race        int       `json:"race"`
	Age         int       `json:"age"`
	BloodType   int       `json:"blood_type"`
	Gender      bool      `json:"gender"`
	Height      float32   `json:"height"`
	Weight      float32   `json:"weight"`
	SmkStat     int       `json:"smk_stat"`
	AlcStat     int       `json:"alc_stat"`
	Other       string    `json:"other"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}
