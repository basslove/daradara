package repository

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
)

type OperatorRepository interface {
	FindOneByEmail(context.Context, string) (*model.Operator, error)
	Create(ctx context.Context, customer *model.Operator) (int64, error)
	Update(ctx context.Context, customer *model.Operator) error
}
