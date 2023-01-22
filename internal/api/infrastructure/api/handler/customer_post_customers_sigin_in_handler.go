package handler

import (
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	usecase "github.com/basslove/daradara/internal/api/usecase/customer"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerPostCustomersSignInHandler struct {
	InputPort usecase.PostCustomersSignInInputPort
}

func (h *CustomerPostCustomersSignInHandler) CustomerPostCustomersSignIn(ctx echo.Context) error {
	var request openapi_service.CustomerPostCustomersSignInRequestBody
	if err := ctx.Bind(&request); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), request.Email, request.Password, ctx.RealIP())
	if err != nil {
		return err
	}

	var customer *openapi_service.Customer
	if output.Customer != nil {
		customer = openapi_service.BuildCustomerPostCustomersSignInResponse(output.Customer)
	}

	return ctx.JSON(http.StatusCreated, openapi_service.CustomerPostCustomersSignInResponse{Customer: customer, Token: &output.Token})
}
