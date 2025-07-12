package main

import (
	"net/http"
	"os"
	"voucher-seat-assignment/config"
	"voucher-seat-assignment/models"
	"voucher-seat-assignment/routes"
	"voucher-seat-assignment/validation"
)

func main() {
	validation.Init()
	db, err := config.ConnectDatabase()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if os.Getenv("RESET_DB") == "true" {
		db.Exec("DELETE FROM vouchers")
		os.Exit(0)
	}

	db.AutoMigrate(&models.Voucher{})
	http.ListenAndServe(":8080", routes.SetupRouter(db))
}
