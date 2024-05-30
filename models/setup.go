package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/futurefarmerapi?parseTime=true"))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&User{}, &SensorData{}, &RelayStatus{}, &RelayConfig{}, &RelayHistory{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&SensorData{})
	db.AutoMigrate(&RelayStatus{})
	db.AutoMigrate(&RelayConfig{})
	db.AutoMigrate(&RelayHistory{})
	DB = db

	createRelayStatus()
	createRelayConfig()
}

func createRelayStatus(){
	relayStatus := RelayStatus{
		Ph_up:   0,
		Ph_down: 0,
		Nut_A:   0,
		Nut_B:   0,
		Fan:     0,
		Light:   0,
		CreatedAt: time.Now(),
	}

	// Insert the new record into the database
	result := DB.Create(&relayStatus)
	if result.Error != nil {
		panic("failed to insert relay status record")
	}
}

func createRelayConfig(){
	relayConfig := RelayConfig{
		Ph_up:   20,
		Ph_down: 20,
		Nut_A:   20,
		Nut_B:   20,
		Fan:     20,
		Light:   20,
		CreatedAt: time.Now(),
	}

	// Insert the new record into the database
	result := DB.Create(&relayConfig)
	if result.Error != nil {
		panic("failed to insert relay status record")
	}
}