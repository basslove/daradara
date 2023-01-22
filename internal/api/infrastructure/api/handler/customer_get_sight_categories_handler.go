package handler

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	usecase "github.com/basslove/daradara/internal/api/usecase/customer"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerGetSightCategoriesHandler struct {
	Auth      CustomerAuth
	InputPort usecase.GetSightCategoriesInputPort
}

func (h *CustomerGetSightCategoriesHandler) CustomerGetSightCategories(ctx echo.Context, params openapi_service.CustomerGetSightCategoriesParams) error {
	customer, err := h.Auth.fetchFromContext(ctx)
	if err != nil {
		return err
	}

	var limit uint64
	if params.Limit != nil && *params.Limit >= 0 {
		limit = uint64(*params.Limit)
	}
	var offset uint64
	if params.Offset != nil && *params.Offset >= 0 {
		offset = uint64(*params.Offset)
	}
	var name string
	if params.Name != nil {
		name = *params.Name
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), customer, name, offset, limit)
	if err != nil {
		return err
	}

	sightCategories := openapi_service.BuildCustomerGetSightCategoriesResponse(output.SightCategories)
	return ctx.JSON(http.StatusOK, openapi_service.CustomerGetSightCategoriesResponse{SightCategories: sightCategories})
}
