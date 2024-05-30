package models

import "time"

type RelayStatus struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph_up          float64   `gorm:"type:integer" json:"ph_up"`
	Ph_down         float64   `gorm:"type:integer" json:"ph_down"`
	Nut_A float64   `gorm:"type:integer" json:"nut_a"`
	Nut_B    float64   `gorm:"type:integer" json:"nut_B"`
	Fan    float64   `gorm:"type:integer" json:"fan"`
	Light    float64   `gorm:"type:integer" json:"light"`
	CreatedAt   time.Time `json:"created_at"`
}


