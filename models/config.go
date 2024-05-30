package models

import "time"

type RelayConfig struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph_up          float64   `gorm:"type:decimal(18,2)" json:"ph_up"`
	Ph_down         float64   `gorm:"type:decimal(18,2)" json:"ph_down"`
	Nut_A float64   `gorm:"type:decimal(18,2)" json:"nut_a"`
	Nut_B    float64   `gorm:"type:decimal(18,2)" json:"nut_B"`
	Fan    float64   `gorm:"type:decimal(18,2)" json:"fan"`
	Light    float64   `gorm:"type:decimal(18,2)" json:"light"`
	IsSync   int64 `gorm:"type:integer" json:"is_sync"`
	CreatedAt   time.Time `json:"created_at"`
}
