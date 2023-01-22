package operator_test

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/domain/model"
	rm "github.com/basslove/daradara/internal/api/domain/repository/mock"
	"github.com/basslove/daradara/internal/api/errors"
	usecase "github.com/basslove/daradara/internal/api/usecase/operator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type getSightCategoriesTestCase struct {
	name                        string
	operator                    *model.Operator
	sightCategoryRepositoryMock func(ctx context.Context, ctrl *gomock.Controller) *rm.MockSightCategoryRepository
	expectedErr                 error
	expectedOutput              *usecase.GetSightCategoriesOutput
}

func TestGetSightCategoriesInteractor_Execute(t *testing.T) {
	t.Parallel()

	offset := uint64(0)
	limit := uint64(0)
	name := "test"

	sightCategories := []*model.SightCategory{
		{
			ID: 1,
		},
	}
	operator := &model.Operator{ID: 1}

	testCases := []getSightCategoriesTestCase{
		{
			name:     "operator is nil",
			operator: nil,
			sightCategoryRepositoryMock: func(ctx context.Context, ctrl *gomock.Controller) *rm.MockSightCategoryRepository {
				repo := rm.NewMockSightCategoryRepository(ctrl)
				return repo
			},
			expectedErr:    errors.ErrOperatorNilNotAllowed,
			expectedOutput: nil,
		},
		{
			name:     "count is zero",
			operator: operator,
			sightCategoryRepositoryMock: func(ctx context.Context, ctrl *gomock.Controller) *rm.MockSightCategoryRepository {
				repo := rm.NewMockSightCategoryRepository(ctrl)
				repo.EXPECT().FindByName(ctx, name, offset, limit).Return([]*model.SightCategory{}, nil)
				return repo
			},
			expectedErr:    nil,
			expectedOutput: &usecase.GetSightCategoriesOutput{SightCategories: []*model.SightCategory{}},
		},
		{
			name:     "db unexpected error",
			operator: operator,
			sightCategoryRepositoryMock: func(ctx context.Context, ctrl *gomock.Controller) *rm.MockSightCategoryRepository {
				repo := rm.NewMockSightCategoryRepository(ctrl)
				repo.EXPECT().FindByName(ctx, name, offset, limit).Return(nil, fmt.Errorf("error"))
				return repo
			},
			expectedErr:    fmt.Errorf("error"),
			expectedOutput: nil,
		},
		{
			name:     "success",
			operator: operator,
			sightCategoryRepositoryMock: func(ctx context.Context, ctrl *gomock.Controller) *rm.MockSightCategoryRepository {
				repo := rm.NewMockSightCategoryRepository(ctrl)
				repo.EXPECT().FindByName(ctx, name, offset, limit).Return(sightCategories, nil)
				return repo
			},
			expectedErr:    nil,
			expectedOutput: &usecase.GetSightCategoriesOutput{SightCategories: sightCategories},
		},
	}

	for _, tc := range testCases {
		func(tc getSightCategoriesTestCase) {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				ctx := context.Background()

				output, err := usecase.NewGetSightCategoriesUsecase(tc.sightCategoryRepositoryMock(ctx, ctrl)).Execute(ctx, tc.operator, name, offset, limit)
				assert.Equal(t, tc.expectedErr, err)
				assert.Equal(t, tc.expectedOutput, output)
			})
		}(tc)
	}
}
