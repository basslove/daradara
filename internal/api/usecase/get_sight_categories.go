package usecase

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
)

type GetSightCategoriesInputPort interface {
	Execute(ctx context.Context, name string, offset, limit uint64) (*GetSightCategoriesOutput, error)
}

type GetSightCategoriesOutput struct {
	SightCategories []*model.SightCategory
}

type GetSightCategoriesInteractor struct {
	sightCategoryRepository repository.SightCategoryRepository
}

func NewGetSightCategoriesUsecase(sc repository.SightCategoryRepository) GetSightCategoriesInputPort {
	return &GetSightCategoriesInteractor{sightCategoryRepository: sc}
}

func (i *GetSightCategoriesInteractor) Execute(ctx context.Context, name string, offset, limit uint64) (*GetSightCategoriesOutput, error) {
	sightCategories, err := i.sightCategoryRepository.FindByName(ctx, name, offset, limit)
	if err != nil {
		return nil, err
	}

	return &GetSightCategoriesOutput{SightCategories: sightCategories}, nil
}
