package handler

import (
	"github.com/basslove/daradara/internal/api/registry"
	"github.com/basslove/daradara/internal/api/usecase"
)

type Handler struct {
	GetSightCategoriesHandler
}

func NewHandler(repository *registry.Repository) (*Handler, error) {
	return &Handler{
		GetSightCategoriesHandler: GetSightCategoriesHandler{
			InputPort: usecase.NewGetSightCategoriesUsecase(repository.SightCategoryRepository),
		},
	}, nil
}
