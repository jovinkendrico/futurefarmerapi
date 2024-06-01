package models

import "time"

type RelayStatus struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph_up          int64   `gorm:"type:integer" json:"ph_up"`
	Is_manual_1    int64   `gorm:"type:integer; default:0"  json:"is_manual_1"`
	Ph_down         int64   `gorm:"type:integer" json:"ph_down"`
	Is_manual_2    int64   `gorm:"type:integer; default:0"  json:"is_manual_2"`
	Nut_a int64   `gorm:"type:integer" json:"nut_a"`
	Is_manual_3    int64   `gorm:"type:integer; default:0"  json:"is_manual_3"`
	Nut_b    int64   `gorm:"type:integer" json:"nut_B"`
	Is_manual_4    int64   `gorm:"type:integer; default:0"  json:"is_manual_4"`
	Fan    int64   `gorm:"type:integer" json:"fan"`
	Is_manual_5    int64   `gorm:"type:integer; default:0"  json:"is_manual_5"`
	Light    int64   `gorm:"type:integer" json:"light"`
	Is_manual_6    int64   `gorm:"type:integer; default:0"  json:"is_manual_6"`
	CreatedAt   time.Time `json:"created_at"`
}
