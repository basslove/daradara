package handler

import (
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	"github.com/basslove/daradara/internal/api/interface/presenter/forms"
	usecase "github.com/basslove/daradara/internal/api/usecase/customer"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerPostCustomersHandler struct {
	InputPort usecase.PostCustomersInputPort
}

func (h *CustomerPostCustomersHandler) CustomerPostCustomers(ctx echo.Context) error {
	var request openapi_service.CustomerPostCustomersRequestBody
	if err := ctx.Bind(&request); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	requestForm := forms.NewCustomerPostCustomersForm(request)
	if err := requestForm.Validate(); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), requestForm)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, openapi_service.CustomerPostCustomersResponse{Id: output.Customer.ID})
}
