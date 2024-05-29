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
	dsn := "root:@tcp(localhost:3306)/?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to MySQL")
	}

	// Check if the database exists, and create it if it doesn't
	createDatabaseIfNotExists(db, "futurefarmerapi")

	// Connect to the `futurefarmerapi` database
	dsn = "root:@tcp(localhost:3306)/futurefarmerapi?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to futurefarmerapi database")
	}

	// Perform auto migrations
	err = db.AutoMigrate(&User{}, &SensorData{}, &RelayStatus{}, &RelayConfig{}, &RelayHistory{})
	if err != nil {
		panic("failed to migrate database")
	}

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
		Ph_up:     0,
		Ph_down:   0,
		Nut_a:     0,
		Nut_b:     0,
		Fan:       0,
		Light:     0,
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
		Ph_up:     20,
		Ph_down:   20,
		Nut_a:     20,
		Nut_b:     20,
		Fan:       20,
		Light:     20,
		CreatedAt: time.Now(),
	}

	// Insert the new record into the database
	result := DB.Create(&relayConfig)
	if result.Error != nil {
		panic("failed to insert relay status record")
	}
}
