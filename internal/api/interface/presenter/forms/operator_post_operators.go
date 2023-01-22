package forms

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	"github.com/go-playground/validator/v10"
)

type OperatorPostOperatorsForm struct {
	Email                string `validate:"required"`
	Password             string `validate:"required"`
	PasswordConfirmation string `validate:"required"`
	Name                 string `validate:"required"`
	DisplayName          string `validate:"required"`
	IsGod                bool
	Level                string `validate:"required"`
	Belong               string `validate:"required"`
}

func NewOperatorPostOperatorsForm(req openapi_service.OperatorPostOperatorsRequestBody) *OperatorPostOperatorsForm {
	return &OperatorPostOperatorsForm{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		Name:                 req.Name,
		DisplayName:          req.DisplayName,
		IsGod:                req.IsGod,
		Level:                string(req.Level),
		Belong:               string(req.Belong),
	}
}

func (fm OperatorPostOperatorsForm) Validate() error {
	v := validator.New()
	if err := v.Struct(fm); err != nil {
		return err
	}
	return nil
}
