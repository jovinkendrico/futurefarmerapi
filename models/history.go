//untuk insert history
package models

import "time"

type RelayHistory struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"type:varchar(255)" json:"type"` 
	Status    string    `gorm:"type:varchar(255)" json:"status"` 
	CreatedAt   time.Time `json:"created_at"`
}