package repository

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
)

type ThrottleRepository interface {
	FindOneByKeyAndType(context.Context, string, string) (*model.Throttle, error)
	Create(context.Context, *model.Throttle) error
	Update(context.Context, *model.Throttle) error
	Delete(context.Context, *model.Throttle) error
}
