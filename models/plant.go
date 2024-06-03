package models

import "time"

type Plant struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"varchar(200)" json:"nama"`
	Tanggal   time.Time `gorm:"date" json:"tanggal"`
	Umur      float64   `gorm:"decimal(18,2)" json:"umur"`
	CreatedAt time.Time `gorm:"datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"datetime" json:"updated_at"`
}
