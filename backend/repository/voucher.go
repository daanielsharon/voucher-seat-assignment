package repository

import (
	"errors"
	"voucher-seat-assignment/models"

	"gorm.io/gorm"
)

type VoucherRepositoryInterface interface {
	CheckVoucher(flightNumber string, flightDate string) (*models.Voucher, error)
	GenerateVoucher(voucher *models.Voucher) error
}

type VoucherRepositoryImpl struct {
	DB *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepositoryInterface {
	return &VoucherRepositoryImpl{DB: db}
}

func (v *VoucherRepositoryImpl) CheckVoucher(flightNumber string, flightDate string) (*models.Voucher, error) {
	var voucher models.Voucher

	err := v.DB.Where("flight_number = ? AND flight_date = ?", flightNumber, flightDate).First(&voucher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &voucher, nil
}

func (v *VoucherRepositoryImpl) GenerateVoucher(voucher *models.Voucher) error {
	return v.DB.Create(voucher).Error
}
