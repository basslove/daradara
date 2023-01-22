package repository

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
)

type CustomerRepository interface {
	FindOneByName(context.Context, string) (*model.Customer, error)
	FindOneByEmail(context.Context, string) (*model.Customer, error)
	Create(ctx context.Context, customer *model.Customer) (int64, error)
	Update(ctx context.Context, customer *model.Customer) error
}
