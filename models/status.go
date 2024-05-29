package models

import "time"

type RelayStatus struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph_up          float64   `gorm:"type:decimal(18,2)" json:"ph_up"`
	Ph_down         float64   `gorm:"type:decimal(18,2)" json:"ph_down"`
	Nut_A float64   `gorm:"type:decimal(18,2)" json:"nut_a"`
	Nut_B    float64   `gorm:"type:decimal(18,2)" json:"nut_B"`
	Fan    float64   `gorm:"type:decimal(18,2)" json:"fan"`
	Light    float64   `gorm:"type:decimal(18,2)" json:"light"`
	CreatedAt   time.Time `json:"created_at"`
}


