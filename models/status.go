package models

import "time"

type RelayStatus struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph_up          int64   `gorm:"type:integer" json:"ph_up"`
	Ph_down         int64   `gorm:"type:integer" json:"ph_down"`
	Nut_a int64   `gorm:"type:integer" json:"nut_a"`
	Nut_b    int64   `gorm:"type:integer" json:"nut_B"`
	Fan    int64   `gorm:"type:integer" json:"fan"`
	Light    int64   `gorm:"type:integer" json:"light"`
	CreatedAt   time.Time `json:"created_at"`
}
