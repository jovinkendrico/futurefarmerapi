package models

import "time"

type SensorData struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph          float64   `gorm:"type:decimal(18,2)" json:"ph"`
	Tds         float64   `gorm:"type:decimal(18,2)" json:"tds"`
	Temperature float64   `gorm:"type:decimal(18,2)" json:"temperature"`
	Humidity    float64   `gorm:"type:decimal(18,2)" json:"humidity"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}