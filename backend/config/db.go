package config

import (
	"log"
	"voucher-seat-assignment/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("vouchers.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = database.AutoMigrate(&models.Voucher{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return database, nil
}
