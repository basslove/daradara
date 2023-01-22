package service

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
)

type Throttler interface {
	IsBlocked(ctx context.Context, key, keyType string) (bool, error)
	Increase(ctx context.Context, key, keyType string) error
	Clear(ctx context.Context, key, keyType string) error
}

type throtterImpl struct {
	throttleRepository repository.ThrottleRepository
}

func NewThrottler(tr repository.ThrottleRepository) Throttler {
	return &throtterImpl{throttleRepository: tr}
}

func (r *throtterImpl) IsBlocked(ctx context.Context, key, keyType string) (bool, error) {
	throttle, err := r.throttleRepository.FindOneByKeyAndType(ctx, key, keyType)
	if err != nil {
		return false, err
	}
	if throttle == nil {
		return false, nil
	}

	return throttle.IsBlocked(), nil
}

func (r *throtterImpl) Increase(ctx context.Context, key, keyType string) error {
	throttle, err := r.throttleRepository.FindOneByKeyAndType(ctx, key, keyType)
	if err != nil {
		return err
	}
	if throttle == nil {
		m := model.NewThrottle(key, keyType)
		if err = r.throttleRepository.Create(ctx, m); err != nil {
			return err
		}
		return nil
	}

	throttle.Increase()

	return r.throttleRepository.Update(ctx, throttle)
}

func (r *throtterImpl) Clear(ctx context.Context, key, keyType string) error {
	throttle := model.NewThrottle(key, keyType)

	return r.throttleRepository.Delete(ctx, throttle)
}
