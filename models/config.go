package models

type RelayConfig struct {
	Id      int64   `gorm:"primaryKey" json:"id"`
	Ph_up   float64 `gorm:"type:decimal(18,2)" json:"ph_up"`
	Ph_down float64 `gorm:"type:decimal(18,2)" json:"ph_down"`
	Nut_A   float64 `gorm:"type:decimal(18,2)" json:"nut_a"`
	Nut_B   float64 `gorm:"type:decimal(18,2)" json:"nut_b"`
	Fan     float64 `gorm:"type:decimal(18,2)" json:"fan"`
	Light   float64 `gorm:"type:decimal(18,2)" json:"light"`
	IsSync  int64   `gorm:"type:integer" json:"is_sync"`
}

type LevelConfig struct {
	Id          int64   `gorm:"primaryKey" json:"id"`
	Ph_low      float64 `gorm:"type:decimal(18,2)" json:"ph_low"`
	Ph_high     float64 `gorm:"type:decimal(18,2)" json:"ph_high"`
	Tds         float64 `gorm:"type:decimal(18,2)" json:"tds"`
	Temperature float64 `gorm:"type:decimal(18,2)" json:"temperature"`
	Humidity    float64 `gorm:"type:decimal(18,2)" json:"humidity"`
}
