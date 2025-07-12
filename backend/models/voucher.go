package models

import "time"

type Voucher struct {
	ID           uint `gorm:"primaryKey"`
	CrewName     string
	CrewID       string
	FlightNumber string `gorm:"index:idx_flight_date"`
	FlightDate   string `gorm:"index:idx_flight_date"`
	AircraftType string
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    time.Time
}

// data transfer object (DTO)
type DTOVoucherRequest struct {
	CrewName     string `json:"name" validate:"required"`
	CrewID       string `json:"id" validate:"required"`
	FlightNumber string `json:"flightNumber" validate:"required"`
	FlightDate   string `json:"date" validate:"required,datetime=2006-01-02"`
	AircraftType string `json:"aircraft" validate:"required,oneof=ATR 'Airbus 320' 'Boeing 737 Max'"`
}

// built to return to client
type DTOVoucherGenerationResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}

type DTOVoucherCheckResponse struct {
	Exists bool `json:"exists"`
}

type DTOVoucherCheckRequest struct {
	FlightNumber string `json:"flightNumber" validate:"required"`
	FlightDate   string `json:"date" validate:"required,datetime=2006-01-02"`
}
