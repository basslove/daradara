package query

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
)

type SightGenreRelationQuery interface {
	FindByNameAndCategoryID(context.Context, string, uint64, uint64, uint64) ([]*model.SightGenreRelation, error)
}
