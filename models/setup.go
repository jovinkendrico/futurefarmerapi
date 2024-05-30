package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Connect to MySQL server without specifying a database
	dsn := "root:@tcp(localhost:3307)/?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to MySQL")
	}

	// Check if the database exists, and create it if it doesn't
	createDatabaseIfNotExists(db, "futurefarmerapi")

	// Connect to the `futurefarmerapi` database
	dsn = "root:@tcp(localhost:3307)/futurefarmerapi?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to futurefarmerapi database")
	}

	// Perform auto migrations
	err = db.AutoMigrate(&User{}, &SensorData{}, &RelayStatus{}, &RelayConfig{}, &RelayHistory{})
	if err != nil {
		panic("failed to migrate database")
	}
	db.Migrator().DropTable(&User{}, &SensorData{}, &RelayStatus{}, &RelayConfig{}, &RelayHistory{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&SensorData{})
	db.AutoMigrate(&RelayStatus{})
	db.AutoMigrate(&RelayConfig{})
	db.AutoMigrate(&RelayHistory{})

	// Assign the connected DB to the global variable
	DB = db

	// Create initial data if necessary
	createRelayStatus()
	createRelayConfig()
}

func createDatabaseIfNotExists(db *gorm.DB, dbName string) {
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := db.Exec(sql).Error; err != nil {
		panic(fmt.Sprintf("failed to create database %s: %v", dbName, err))
	}
}

func createRelayStatus() {
	relayStatus := RelayStatus{
		Ph_up:    1,
		Ph_down:  1,
		Nut_a:    1,
		Nut_b:    1,
		Fan:      1,
		Light:    1,
		CreatedAt: time.Now(),
	}

	// Insert the new record into the database
	result := DB.Create(&relayStatus)
	if result.Error != nil {
		panic("failed to insert relay status record")
	}
}

func createRelayConfig() {
	relayConfig := RelayConfig{
		Ph_up:   1,
		Ph_down: 20,
		Nut_A:   20,
		Nut_B:   20,
		Fan:     20,
		Light:   20,
		IsSync:   1,
		CreatedAt: time.Now(),
	}

	// Insert the new record into the database
	result := DB.Create(&relayConfig)
	if result.Error != nil {
		panic("failed to insert relay status record")
	}
}
