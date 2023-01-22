package operator_test

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/domain/model"
	qm "github.com/basslove/daradara/internal/api/domain/query/mock"
	"github.com/basslove/daradara/internal/api/errors"
	usecase "github.com/basslove/daradara/internal/api/usecase/operator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type getSightGenresTestCase struct {
	name                        string
	operator                    *model.Operator
	sightGenreRelationQueryMock func(ctx context.Context, ctrl *gomock.Controller) *qm.MockSightGenreRelationQuery
	expectedErr                 error
	expectedOutput              *usecase.GetSightGenresOutput
}

func TestGetSightGenresInteractor_Execute(t *testing.T) {
	t.Parallel()

	sightGenreRelations := []*model.SightGenreRelation{
		{
			ID: 1,
		},
	}
	operator := &model.Operator{ID: 1}

	testCases := []getSightGenresTestCase{
		{
			name:     "operator is nil",
			operator: nil,
			sightGenreRelationQueryMock: func(ctx context.Context, ctrl *gomock.Controller) *qm.MockSightGenreRelationQuery {
				q := qm.NewMockSightGenreRelationQuery(ctrl)
				return q
			},
			expectedErr:    errors.ErrOperatorNilNotAllowed,
			expectedOutput: nil,
		},
		{
			name:     "count is zero",
			operator: operator,
			sightGenreRelationQueryMock: func(ctx context.Context, ctrl *gomock.Controller) *qm.MockSightGenreRelationQuery {
				q := qm.NewMockSightGenreRelationQuery(ctrl)
				q.EXPECT().FindByNameAndCategoryID(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*model.SightGenreRelation{}, nil)
				return q
			},
			expectedErr:    nil,
			expectedOutput: &usecase.GetSightGenresOutput{SightGenreRelations: []*model.SightGenreRelation{}},
		},
		{
			name:     "db unexpected error",
			operator: operator,
			sightGenreRelationQueryMock: func(ctx context.Context, ctrl *gomock.Controller) *qm.MockSightGenreRelationQuery {
				q := qm.NewMockSightGenreRelationQuery(ctrl)
				q.EXPECT().FindByNameAndCategoryID(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))
				return q
			},
			expectedErr:    fmt.Errorf("error"),
			expectedOutput: nil,
		},
		{
			name:     "success",
			operator: operator,
			sightGenreRelationQueryMock: func(ctx context.Context, ctrl *gomock.Controller) *qm.MockSightGenreRelationQuery {
				q := qm.NewMockSightGenreRelationQuery(ctrl)
				q.EXPECT().FindByNameAndCategoryID(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(sightGenreRelations, nil)
				return q
			},
			expectedErr:    nil,
			expectedOutput: &usecase.GetSightGenresOutput{SightGenreRelations: sightGenreRelations},
		},
	}

	for _, tc := range testCases {
		func(tc getSightGenresTestCase) {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				ctx := context.Background()

				output, err := usecase.NewGetSightGenresUsecase(tc.sightGenreRelationQueryMock(ctx, ctrl)).Execute(ctx, tc.operator, "", 0, 0, 0)
				assert.Equal(t, tc.expectedErr, err)
				assert.Equal(t, tc.expectedOutput, output)
			})
		}(tc)
	}
}
