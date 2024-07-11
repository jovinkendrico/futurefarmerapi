package models

import "time"

type SensorData struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph          float64   `gorm:"type:decimal(18,2)" json:"ph"`
	Tds         int64     `gorm:"type:integer" json:"tds"`
	Temperature float64   `gorm:"type:decimal(18,2)" json:"temperature"`
	Humidity    float64   `gorm:"type:decimal(18,2)" json:"humidity"`
	CreatedAt   time.Time `json:"created_at"`
}
