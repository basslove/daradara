package handler

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	usecase "github.com/basslove/daradara/internal/api/usecase/operator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OperatorGetSightGenresHandler struct {
	Auth      OperatorAuth
	InputPort usecase.GetSightGenresInputPort
}

func (h *OperatorGetSightGenresHandler) OperatorGetSightGenres(ctx echo.Context, params openapi_service.OperatorGetSightGenresParams) error {
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
	var sightCategoryID uint64
	if params.SightCategoryId != nil && *params.SightCategoryId >= 0 {
		sightCategoryID = uint64(*params.SightCategoryId)
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), operator, name, sightCategoryID, offset, limit)
	if err != nil {
		return err
	}

	sightGenres := openapi_service.BuildOperatorGetSightGenresResponse(output.SightGenreRelations)
	return ctx.JSON(http.StatusOK, openapi_service.OperatorGetSightGenresResponse{SightGenres: sightGenres})
}
