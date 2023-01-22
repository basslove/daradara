package repository

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
)

type SightGenreRepository interface {
	FindOneByID(context.Context, uint64, bool) (*model.SightGenre, error)
	FindByNameAndCategoryID(context.Context, string, uint64, uint64, uint64) ([]*model.SightGenre, error)
}
