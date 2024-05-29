package models

import "time"

type RelayConfig struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Ph_up     float64   `gorm:"type:int" json:"ph_up"`
	Ph_down   float64   `gorm:"type:int" json:"ph_down"`
	Nut_a     float64   `gorm:"type:int" json:"nut_a"`
	Nut_b     float64   `gorm:"type:int" json:"nut_B"`
	Fan       float64   `gorm:"type:int" json:"fan"`
	Light     float64   `gorm:"type:int" json:"light"`
	CreatedAt time.Time `json:"created_at"`
}
