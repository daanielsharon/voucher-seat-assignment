package controllers

import (
	"fmt"
	"voucher-seat-assignment/models"
	"voucher-seat-assignment/services"
	"voucher-seat-assignment/utils"
	"voucher-seat-assignment/validation"

	"net/http"
)

type VoucherControllerInterface interface {
	CheckVoucher(w http.ResponseWriter, r *http.Request)
	GenerateVoucher(w http.ResponseWriter, r *http.Request)
}

type VoucherControllerImpl struct {
	service services.VoucherServiceInterface
}

func NewVoucherController(service services.VoucherServiceInterface) VoucherControllerInterface {
	return &VoucherControllerImpl{
		service: service,
	}
}

func (v *VoucherControllerImpl) CheckVoucher(w http.ResponseWriter, r *http.Request) {
	var req models.DTOVoucherCheckRequest

	if err := utils.JSONRequest(w, r, &req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err := validation.Validate.Struct(req)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request")
		return
	}

	voucher, err := v.service.CheckVoucher(req.FlightNumber, req.FlightDate)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "failed to check voucher")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "voucher checked successfully", voucher)
}

func (v *VoucherControllerImpl) GenerateVoucher(w http.ResponseWriter, r *http.Request) {
	var req models.DTOVoucherRequest

	if err := utils.JSONRequest(w, r, &req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request")
		return
	}

	err := validation.Validate.Struct(req)
	if err != nil {
		fmt.Println("err", err)
		utils.JSONError(w, http.StatusBadRequest, "invalid request")
		return
	}

	voucherData, err := v.service.CheckVoucher(req.FlightNumber, req.FlightDate)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "failed to check voucher")
		return
	}

	if voucherData.Exists {
		utils.JSONError(w, http.StatusBadRequest, "voucher already created")
		return
	}

	voucher, err := v.service.GenerateVoucher(&req)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "failed to generate voucher")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "voucher generated successfully", voucher)
}
