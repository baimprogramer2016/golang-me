package entity

import "time"

type Encounter struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Poli      string `json:"poli"`
	Diagnosa  string `json:"diagnosa"`
	Umur      int    `json:"umur"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Encounter) TableName() string {
	return "encounter"
}
