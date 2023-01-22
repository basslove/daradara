package forms

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	"github.com/go-playground/validator/v10"
	"time"
)

type CustomerPostCustomersForm struct {
	Email                string `validate:"required"`
	Password             string `validate:"required"`
	PasswordConfirmation string `validate:"required"`
	Name                 string `validate:"required"`
	DisplayName          string `validate:"required"`
	PhoneNumber          string `validate:"required"`
	Introduction         string `validate:"required"`
	AllowPlansDisplayed  bool
	Birthday             time.Time `validate:"required"`
	Gender               string    `validate:"required"`
	Generation           string    `validate:"required"`
}

func NewCustomerPostCustomersForm(req openapi_service.CustomerPostCustomersRequestBody) *CustomerPostCustomersForm {
	return &CustomerPostCustomersForm{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		Name:                 req.Name,
		DisplayName:          req.DisplayName,
		PhoneNumber:          req.PhoneNumber,
		Introduction:         req.Introduction,
		AllowPlansDisplayed:  req.AllowPlansDisplayed,
		Birthday:             time.Date(req.Birthday.Year(), req.Birthday.Month(), req.Birthday.Day(), 0, 0, 0, 0, time.UTC),
		Gender:               string(req.Gender),
		Generation:           string(req.Generation),
	}
}

func (fm CustomerPostCustomersForm) Validate() error {
	v := validator.New()
	if err := v.Struct(fm); err != nil {
		return err
	}
	return nil
}
