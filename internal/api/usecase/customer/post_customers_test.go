package customer_test

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	repository "github.com/basslove/daradara/internal/api/domain/repository/mock"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/interface/presenter/forms"
	usecase "github.com/basslove/daradara/internal/api/usecase/customer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type postCustomersTestCase struct {
	name           string
	form           *forms.CustomerPostCustomersForm
	setMocks       func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository)
	expectedErr    error
	expectedOutput *usecase.PostCustomersOutput
}

func TestPostCustomersInteractor_Execute(t *testing.T) {
	t.Parallel()

	newID := uint64(1)

	testCases := []postCustomersTestCase{
		{
			name: "form is nil",
			form: nil,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "operator FindOneByEmail is error",
			form: &forms.CustomerPostCustomersForm{Email: "email"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, errors.ErrNotFound),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    errors.ErrNotFound,
			expectedOutput: nil,
		},
		{
			name: "operator FindOneByEmail is exist",
			form: &forms.CustomerPostCustomersForm{Email: "email"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(&model.Operator{Email: "email"}, nil),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    errors.ErrAlreadyExists,
			expectedOutput: nil,
		},
		{
			name: "customer FindOneByEmail is error",
			form: &forms.CustomerPostCustomersForm{Email: "email"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, errors.ErrNotFound),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    errors.ErrNotFound,
			expectedOutput: nil,
		},
		{
			name: "customer FindOneByEmail exist",
			form: &forms.CustomerPostCustomersForm{Email: "email"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().FindOneByEmail(ctx, "email").Return(&model.Customer{Email: "email"}, nil),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    errors.ErrAlreadyExists,
			expectedOutput: nil,
		},
		{
			name: "password != password confirmation",
			form: &forms.CustomerPostCustomersForm{Email: "email", Password: "pass", PasswordConfirmation: "aaaa"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    errors.ErrBadRequest,
			expectedOutput: nil,
		},
		{
			name: "customer Create is error",
			form: &forms.CustomerPostCustomersForm{Email: "email", Password: "pass", PasswordConfirmation: "pass"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().Create(ctx, gomock.Any()).Return(int64(0), errors.ErrInternalServerError),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "success",
			form: &forms.CustomerPostCustomersForm{Email: "email", Password: "pass", PasswordConfirmation: "pass"},
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, operatorRepo *repository.MockOperatorRepository) {
				mocks := []*gomock.Call{
					operatorRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().FindOneByEmail(ctx, "email").Return(nil, nil),
					customerRepo.EXPECT().Create(ctx, gomock.Any()).Return(int64(1), nil),
				}
				gomock.InOrder(mocks...)
			},
			expectedErr:    nil,
			expectedOutput: &usecase.PostCustomersOutput{Customer: &model.Customer{ID: newID}},
		},
	}

	for _, tc := range testCases {
		func(tc postCustomersTestCase) {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				ctx := context.Background()

				customerRepoMock := repository.NewMockCustomerRepository(ctrl)
				operatorRepoMock := repository.NewMockOperatorRepository(ctrl)
				tc.setMocks(ctx, customerRepoMock, operatorRepoMock)

				output, err := usecase.NewPostCustomersUsecase(customerRepoMock, operatorRepoMock).Execute(ctx, tc.form)
				assert.Equal(t, tc.expectedErr, err)
				if output != nil {
					assert.Equal(t, tc.expectedOutput.Customer.ID, output.Customer.ID)
				}
			})
		}(tc)
	}
}
