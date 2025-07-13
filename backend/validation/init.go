package validation

import (
	"time"

	"github.com/go-playground/validator"
)

var Validate *validator.Validate

func Init() {
	Validate = validator.New()
	Validate.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())
		return err == nil
	})
	Validate.RegisterValidation("aircraft", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "ATR" || fl.Field().String() == "Airbus 320" || fl.Field().String() == "Boeing 737 Max"
	})
}
