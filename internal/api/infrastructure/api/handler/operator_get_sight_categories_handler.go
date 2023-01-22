package handler

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	usecase "github.com/basslove/daradara/internal/api/usecase/operator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OperatorGetSightCategoriesHandler struct {
	Auth      OperatorAuth
	InputPort usecase.GetSightCategoriesInputPort
}

func (h *OperatorGetSightCategoriesHandler) OperatorGetSightCategories(ctx echo.Context, params openapi_service.OperatorGetSightCategoriesParams) error {
	operator, err := h.Auth.fetchFromContext(ctx)
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

	output, err := h.InputPort.Execute(ctx.Request().Context(), operator, name, offset, limit)
	if err != nil {
		return err
	}

	sightCategories := openapi_service.BuildOperatorGetSightCategoriesResponse(output.SightCategories)
	return ctx.JSON(http.StatusOK, openapi_service.OperatorGetSightCategoriesResponse{SightCategories: sightCategories})
}
