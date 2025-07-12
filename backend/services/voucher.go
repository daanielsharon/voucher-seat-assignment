package services

import (
	"voucher-seat-assignment/models"
	"voucher-seat-assignment/repository"
	"voucher-seat-assignment/utils"
)

type VoucherServiceInterface interface {
	CheckVoucher(flightNumber string, flightDate string) (*models.DTOVoucherCheckResponse, error)
	GenerateVoucher(voucher *models.DTOVoucherRequest) (*models.DTOVoucherGenerationResponse, error)
}

type VoucherServiceImpl struct {
	repository repository.VoucherRepositoryInterface
}

func NewVoucherService(repository repository.VoucherRepositoryInterface) VoucherServiceInterface {
	return &VoucherServiceImpl{
		repository: repository,
	}
}

func (v *VoucherServiceImpl) CheckVoucher(flightNumber string, flightDate string) (*models.DTOVoucherCheckResponse, error) {
	data, err := v.repository.CheckVoucher(flightNumber, flightDate)
	if err != nil {
		return nil, err
	}

	return &models.DTOVoucherCheckResponse{
		Exists: data != nil,
	}, nil
}

func (v *VoucherServiceImpl) GenerateVoucher(voucher *models.DTOVoucherRequest) (*models.DTOVoucherGenerationResponse, error) {
	seats, err := utils.GenerateRandomSeats(voucher.AircraftType)
	if err != nil {
		return nil, err
	}

	vouchers := []string{}
	vouchers = append(vouchers, seats...)

	if err := v.repository.GenerateVoucher(&models.Voucher{
		CrewName:     voucher.CrewName,
		CrewID:       voucher.CrewID,
		FlightNumber: voucher.FlightNumber,
		FlightDate:   voucher.FlightDate,
		AircraftType: voucher.AircraftType,
		Seat1:        vouchers[0],
		Seat2:        vouchers[1],
		Seat3:        vouchers[2],
	}); err != nil {
		return nil, err
	}

	return &models.DTOVoucherGenerationResponse{
		Success: true,
		Seats:   vouchers,
	}, nil
}
