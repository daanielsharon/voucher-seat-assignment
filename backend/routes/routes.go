package routes

import (
	"net/http"
	"voucher-seat-assignment/controllers"
	"voucher-seat-assignment/repository"
	"voucher-seat-assignment/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := services.NewVoucherService(voucherRepository)
	voucherController := controllers.NewVoucherController(voucherService)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/check", voucherController.CheckVoucher)
		r.Post("/generate", voucherController.GenerateVoucher)
	})

	return r
}
