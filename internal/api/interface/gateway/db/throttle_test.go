package db_test

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/interface/gateway/db"
	"github.com/basslove/daradara/internal/api/pkg/testutil"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestThrottleRepository_FindOneByKeyAndType(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupThrottleRepository())

	repository := db.NewThrottleRepository(testutil.PsqlDB())

	t.Run("exists", func(t *testing.T) {
		value, err := repository.FindOneByKeyAndType(ctx, "test1", "test-type1")
		require.NoError(t, err)
		require.Equal(t, "test1", value.Key)
		require.Equal(t, "test-type1", value.KeyType)
		require.NotEmpty(t, value.HashKey)
	})

	t.Run("not exists", func(t *testing.T) {
		value, err := repository.FindOneByKeyAndType(ctx, "aaa", "test-type1")
		require.NoError(t, err)
		require.Nil(t, value)
	})
}

func TestThrottleRepository_Create(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupThrottleRepository())

	repository := db.NewThrottleRepository(testutil.PsqlDB())

	t.Run("success", func(t *testing.T) {
		key := "hoge1"
		keyType := "hoge-type1"
		m := model.NewThrottle(key, keyType)

		err := repository.Create(ctx, m)
		require.NoError(t, err)

		value, err := repository.FindOneByKeyAndType(ctx, key, keyType)
		require.NoError(t, err)
		require.Equal(t, key, value.Key)
		require.Equal(t, keyType, value.KeyType)
		require.NotEmpty(t, value.HashKey)
	})

	t.Run("fail", func(t *testing.T) {
		key := "test1"
		keyType := "test-type1"
		m := model.NewThrottle(key, keyType)

		err := repository.Create(ctx, m)
		require.Error(t, err)
	})
}

func TestThrottleRepository_Update(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupThrottleRepository())

	repository := db.NewThrottleRepository(testutil.PsqlDB())

	t.Run("success", func(t *testing.T) {
		key := "test1"
		keyType := "test-type1"

		value, err := repository.FindOneByKeyAndType(ctx, key, keyType)
		require.NoError(t, err)
		require.Equal(t, 0, value.Count)

		value.Count = 5

		err = repository.Update(ctx, value)
		require.NoError(t, err)

		value, err = repository.FindOneByKeyAndType(ctx, key, keyType)
		require.NoError(t, err)
		require.Equal(t, 5, value.Count)
	})
}

func TestThrottleRepository_Delete(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupThrottleRepository())

	repository := db.NewThrottleRepository(testutil.PsqlDB())

	t.Run("success", func(t *testing.T) {
		key := "test1"
		keyType := "test-type1"

		value, err := repository.FindOneByKeyAndType(ctx, key, keyType)
		require.NoError(t, err)
		require.Equal(t, key, value.Key)
		require.Equal(t, keyType, value.KeyType)
		require.NotEmpty(t, value.HashKey)

		err = repository.Delete(ctx, value)
		require.NoError(t, err)

		value, err = repository.FindOneByKeyAndType(ctx, key, keyType)
		require.NoError(t, err)
		require.Nil(t, value)
	})
}

func setupThrottleRepository() error {
	testutil.CleanPsqlDB()

	t1 := model.NewThrottle("test1", "test-type1")
	t1.BlockExpiredAt = time.Now()
	t2 := model.NewThrottle("test2", "test-type2")
	t2.BlockExpiredAt = time.Now()

	throttles := []*model.Throttle{t1, t2}
	for _, m := range throttles {
		q := sq.Insert(m.TableName()).Columns("hash_key", "key", "key_type", "count_expired_at", "block_expired_at")
		q = q.Values(m.HashKey, m.Key, m.KeyType, m.CountExpiredAt, m.BlockExpiredAt).PlaceholderFormat(sq.Dollar).RunWith(testutil.PsqlDB())
		_, err := q.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
