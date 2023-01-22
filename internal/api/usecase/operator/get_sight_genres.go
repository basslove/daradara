package operator

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/query"
	"github.com/basslove/daradara/internal/api/errors"
)

type GetSightGenresInputPort interface {
	Execute(ctx context.Context, operator *model.Operator, name string, sightCategoryID, offset, limit uint64) (*GetSightGenresOutput, error)
}

type GetSightGenresOutput struct {
	SightGenreRelations []*model.SightGenreRelation
}

type GetSightGenresInteractor struct {
	sightGenreRelationQuery query.SightGenreRelationQuery
}

func NewGetSightGenresUsecase(q query.SightGenreRelationQuery) GetSightGenresInputPort {
	return &GetSightGenresInteractor{sightGenreRelationQuery: q}
}

func (i *GetSightGenresInteractor) Execute(ctx context.Context, operator *model.Operator, name string, sightCategoryID, offset, limit uint64) (*GetSightGenresOutput, error) {
	if operator == nil {
		return nil, errors.ErrOperatorNilNotAllowed
	}

	SightGenreRelations, err := i.sightGenreRelationQuery.FindByNameAndCategoryID(ctx, name, sightCategoryID, offset, limit)
	if err != nil {
		return nil, err
	}

	return &GetSightGenresOutput{SightGenreRelations: SightGenreRelations}, nil
}
