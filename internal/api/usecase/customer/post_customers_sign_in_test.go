package customer_test

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	repository "github.com/basslove/daradara/internal/api/domain/repository/mock"
	service "github.com/basslove/daradara/internal/api/domain/service/mock"
	"github.com/basslove/daradara/internal/api/errors"
	usecase "github.com/basslove/daradara/internal/api/usecase/customer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type postCustomersSignInArgs struct {
	email     string
	pass      string
	IPAddress string
}

type postCustomersSignInTestCase struct {
	name           string
	args           *postCustomersSignInArgs
	setMocks       func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler)
	expectedErr    error
	expectedOutput *usecase.PostCustomersSignInOutput
}

func TestPostCustomersSignInInteractor_Execute(t *testing.T) {
	t.Parallel()

	args := &postCustomersSignInArgs{
		email:     "email",
		pass:      "pass",
		IPAddress: "0.0.0.1",
	}
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(args.pass), bcrypt.DefaultCost)

	testCases := []postCustomersSignInTestCase{
		{
			name: "throttler IsBlocked is error",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, errors.ErrInternalServerError)
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "customer is blocked",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(true, nil)
			},
			expectedErr:    errors.ErrUnauthorized,
			expectedOutput: nil,
		},
		{
			name: "customer FindOneByEmail is error",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(nil, errors.ErrInternalServerError)
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "customer FindOneByEmail is nil & throttler Increase is error",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(nil, nil)
				throttlerService.EXPECT().Increase(ctx, args.IPAddress, gomock.Any()).Return(errors.ErrInternalServerError)
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "customer FindOneByEmail is nil & throttler Increase is ok",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(nil, nil)
				throttlerService.EXPECT().Increase(ctx, args.IPAddress, gomock.Any()).Return(nil)
			},
			expectedErr:    errors.ErrNotFound,
			expectedOutput: nil,
		},
		{
			name: "customer password is ng & throttler Increase is error",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(&model.Customer{CryptedPassword: "hoge", Email: args.email}, nil)
				throttlerService.EXPECT().Increase(ctx, args.IPAddress, gomock.Any()).Return(errors.ErrInternalServerError)
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "customer password is ng & throttler Increase is ok",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(&model.Customer{CryptedPassword: "hoge", Email: args.email}, nil)
				throttlerService.EXPECT().Increase(ctx, args.IPAddress, gomock.Any()).Return(nil)
			},
			expectedErr:    errors.ErrBadRequest,
			expectedOutput: nil,
		},
		{
			name: "customer Update is error",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(&model.Customer{CryptedPassword: string(hashPass), Email: args.email}, nil)
				customerRepo.EXPECT().Update(ctx, gomock.Any()).Return(errors.ErrInternalServerError)
			},
			expectedErr:    errors.ErrInternalServerError,
			expectedOutput: nil,
		},
		{
			name: "success",
			args: args,
			setMocks: func(ctx context.Context, customerRepo *repository.MockCustomerRepository, throttlerService *service.MockThrottler) {
				throttlerService.EXPECT().IsBlocked(ctx, args.IPAddress, gomock.Any()).Return(false, nil)
				customerRepo.EXPECT().FindOneByEmail(ctx, args.email).Return(&model.Customer{CryptedPassword: string(hashPass), Email: args.email}, nil)
				customerRepo.EXPECT().Update(ctx, gomock.Any()).Return(nil)
			},
			expectedErr:    nil,
			expectedOutput: &usecase.PostCustomersSignInOutput{Customer: &model.Customer{Email: args.email}},
		},
	}

	for _, tc := range testCases {
		func(tc postCustomersSignInTestCase) {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				ctx := context.Background()

				customerRepoMock := repository.NewMockCustomerRepository(ctrl)
				throttlerSerMock := service.NewMockThrottler(ctrl)
				tc.setMocks(ctx, customerRepoMock, throttlerSerMock)

				output, err := usecase.NewPostCustomersSignInUsecase(customerRepoMock, throttlerSerMock).Execute(ctx, tc.args.email, tc.args.pass, tc.args.IPAddress)
				assert.Equal(t, tc.expectedErr, err)
				if output != nil {
					assert.Equal(t, tc.expectedOutput.Customer.Email, output.Customer.Email)
				}
			})
		}(tc)
	}
}
