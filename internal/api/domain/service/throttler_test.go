package service_test

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	repoMock "github.com/basslove/daradara/internal/api/domain/repository/mock"
	"github.com/basslove/daradara/internal/api/domain/service"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestThrottler_IsBlocked(t *testing.T) {
	ctx := context.Background()
	key := "test"
	keyType := model.ThrottleKeyTypeIP

	t.Run("repo return is error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(nil, errors.ErrInternalServerError)
		throttler := service.NewThrottler(MockThrottleRepository)

		result, err := throttler.IsBlocked(ctx, key, keyType)
		require.Error(t, err)
		require.False(t, result)
	})

	t.Run("repo return is nil(not found)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(nil, nil)
		throttler := service.NewThrottler(MockThrottleRepository)

		result, err := throttler.IsBlocked(ctx, key, keyType)
		require.NoError(t, err)
		require.False(t, result)
	})

	t.Run("repo return is unblocked data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := model.NewThrottle(key, keyType)

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(m, nil)
		throttler := service.NewThrottler(MockThrottleRepository)

		result, err := throttler.IsBlocked(ctx, key, keyType)
		require.NoError(t, err)
		require.False(t, result)
	})

	t.Run("repo return is blocked data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := model.NewThrottle(key, keyType)
		m.Block()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(m, nil)
		throttler := service.NewThrottler(MockThrottleRepository)

		result, err := throttler.IsBlocked(ctx, key, keyType)
		require.NoError(t, err)
		require.True(t, result)
	})
}

func TestThrottler_Increase(t *testing.T) {
	ctx := context.Background()
	key := "test"
	keyType := model.ThrottleKeyTypeIP

	t.Run("repo return is error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(nil, errors.ErrInternalServerError)
		throttler := service.NewThrottler(MockThrottleRepository)

		err := throttler.Increase(ctx, key, keyType)
		require.Error(t, err)
	})

	t.Run("first time", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(nil, nil)
		MockThrottleRepository.EXPECT().Create(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, m *model.Throttle) error {
			require.Equal(t, key, m.Key)
			require.Equal(t, 1, m.Count)
			require.Equal(t, time.Time{}, m.BlockExpiredAt)
			require.True(t, m.CountExpiredAt.After(time.Now()))
			return nil
		})

		throttler := service.NewThrottler(MockThrottleRepository)
		err := throttler.Increase(ctx, key, keyType)
		require.NoError(t, err)
	})

	t.Run("count expired", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := model.NewThrottle(key, keyType)
		m.CountExpiredAt = time.Now().Add(-time.Hour * 10)
		m.Count = 5

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(m, nil)
		MockThrottleRepository.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, m *model.Throttle) error {
			require.Equal(t, key, m.Key)
			require.Equal(t, 1, m.Count)
			require.Equal(t, time.Time{}, m.BlockExpiredAt)
			require.True(t, m.CountExpiredAt.After(time.Now()))
			return nil
		})

		throttler := service.NewThrottler(MockThrottleRepository)
		err := throttler.Increase(ctx, key, keyType)
		require.NoError(t, err)
	})

	t.Run("blocked expired", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := model.NewThrottle(key, keyType)
		m.CountExpiredAt = time.Now().Add(-time.Hour * 10)
		m.BlockExpiredAt = time.Now().Add(-time.Hour * 10)
		m.Count = 10

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(m, nil)
		MockThrottleRepository.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, m *model.Throttle) error {
			require.Equal(t, key, m.Key)
			require.Equal(t, 1, m.Count)
			require.Equal(t, time.Time{}, m.BlockExpiredAt)
			require.True(t, m.CountExpiredAt.After(time.Now()))
			return nil
		})

		throttler := service.NewThrottler(MockThrottleRepository)
		err := throttler.Increase(ctx, key, keyType)
		require.NoError(t, err)
	})

	t.Run("block", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := model.NewThrottle(key, keyType)
		m.CountExpiredAt = time.Now().Add(time.Hour * 10)
		m.BlockExpiredAt = time.Now()
		m.Count = 9

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().FindOneByKeyAndType(ctx, key, keyType).Return(m, nil)
		MockThrottleRepository.EXPECT().Update(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, m *model.Throttle) error {
			require.Equal(t, key, m.Key)
			require.Equal(t, 10, m.Count)
			require.True(t, m.BlockExpiredAt.After(time.Now()))
			require.True(t, m.CountExpiredAt.After(time.Now()))
			return nil
		})

		throttler := service.NewThrottler(MockThrottleRepository)
		err := throttler.Increase(ctx, key, keyType)
		require.NoError(t, err)
	})
}

func TestThrottler_Clear(t *testing.T) {
	ctx := context.Background()
	key := "test"
	keyType := model.ThrottleKeyTypeIP

	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().Delete(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, m *model.Throttle) error {
			require.Equal(t, key, m.Key)
			return nil
		})

		throttler := service.NewThrottler(MockThrottleRepository)
		err := throttler.Clear(ctx, key, keyType)
		require.NoError(t, err)
	})
	t.Run("ng", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		MockThrottleRepository := repoMock.NewMockThrottleRepository(ctrl)
		MockThrottleRepository.EXPECT().Delete(ctx, gomock.Any()).Return(errors.ErrInternalServerError)

		throttler := service.NewThrottler(MockThrottleRepository)
		err := throttler.Clear(ctx, key, keyType)
		require.Error(t, err)
	})
}
