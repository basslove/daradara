package repository

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
)

type SightCategoryRepository interface {
	FindOneByID(context.Context, uint64, bool) (*model.SightCategory, error)
	FindByName(context.Context, string, uint64, uint64) ([]*model.SightCategory, error)
}
