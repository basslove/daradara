package handler

import (
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	"github.com/basslove/daradara/internal/api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetSightCategoriesHandler struct {
	InputPort usecase.GetSightCategoriesInputPort
}

func (h *GetSightCategoriesHandler) GetSightCategories(ctx echo.Context, params openapi_service.GetSightCategoriesParams) error {
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

	output, err := h.InputPort.Execute(ctx.Request().Context(), name, offset, limit)
	if err != nil {
		return err
	}

	sightCategories := make([]openapi_service.SightCategory, 0, len(output.SightCategories))
	for _, o := range output.SightCategories {
		s := openapi_service.SightCategory{
			Id:   o.ID,
			Name: o.Name,
		}
		sightCategories = append(sightCategories, s)
	}

	return ctx.JSON(http.StatusOK, openapi_service.GetSightCategoriesResponse{SightCategories: sightCategories})
}
