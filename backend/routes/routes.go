package routes

import (
	"net/http"
	"voucher-seat-assignment/controllers"
	"voucher-seat-assignment/repository"
	"voucher-seat-assignment/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := services.NewVoucherService(voucherRepository)
	voucherController := controllers.NewVoucherController(voucherService)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Post("/check", voucherController.CheckVoucher)
		r.Post("/generate", voucherController.GenerateVoucher)
	})

	return r
}
