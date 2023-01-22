package handler

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	usecase "github.com/basslove/daradara/internal/api/usecase/customer"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerGetSightGenresHandler struct {
	Auth      CustomerAuth
	InputPort usecase.GetSightGenresInputPort
}

func (h *CustomerGetSightGenresHandler) CustomerGetSightGenres(ctx echo.Context, params openapi_service.CustomerGetSightGenresParams) error {
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
	var sightCategoryID uint64
	if params.SightCategoryId != nil && *params.SightCategoryId >= 0 {
		sightCategoryID = uint64(*params.SightCategoryId)
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), customer, name, sightCategoryID, offset, limit)
	if err != nil {
		return err
	}

	sightGenres := openapi_service.BuildCustomerGetSightGenresResponse(output.SightGenreRelations)
	return ctx.JSON(http.StatusOK, openapi_service.CustomerGetSightGenresResponse{SightGenres: sightGenres})
}
