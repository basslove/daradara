package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewThrottleRepository(client *sqlx.DB) repository.ThrottleRepository {
	return &ThrottleRepository{Repository: Repository{client: client}}
}

type ThrottleRepository struct {
	Repository
}

func (r *ThrottleRepository) FindOneByKeyAndType(ctx context.Context, key, keyType string) (*model.Throttle, error) {
	var values []*model.Throttle

	m := model.NewThrottle(key, keyType)

	stmt, err := r.client.Preparex("SELECT * FROM throttles WHERE hash_key = $1 AND key_type = $2")
	if err != nil {
		return nil, fmt.Errorf("ThrottleRepository FindOneByKeyAndType: %w", err)
	}

	if err = stmt.SelectContext(ctx, &values, m.HashKey, m.KeyType); err != nil {
		return nil, fmt.Errorf("ThrottleRepository FindOneByKeyAndType: %w", err)
	}

	if len(values) > 0 {
		return values[0], nil
	}

	return nil, nil
}

func (r *ThrottleRepository) Create(ctx context.Context, m *model.Throttle) error {
	sql := fmt.Sprintf("INSERT INTO %s (hash_key, key_type, key, count, count_expired_at, block_expired_at) VALUES (:hash_key, :key_type, :key, :count, :count_expired_at, :block_expired_at);", m.TableName())
	_, err := r.client.NamedQueryContext(ctx, sql, m)
	if err != nil {
		return fmt.Errorf("ThrottleRepository Create: %w", err)
	}

	return nil
}

func (r *ThrottleRepository) Update(ctx context.Context, m *model.Throttle) error {
	qb := sq.Update(m.TableName()).PlaceholderFormat(sq.Dollar).RunWith(r.client)
	qb = qb.Set("count", m.Count)
	qb = qb.Set("count_expired_at", m.CountExpiredAt)
	qb = qb.Set("block_expired_at", m.BlockExpiredAt)
	qb = qb.Set("updated_at", time.Now())
	qb = qb.Where("hash_key = ?", m.HashKey)
	qb = qb.Where("key_type = ?", m.KeyType)

	_, err := qb.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ThrottleRepository Update: %w", err)
	}

	return nil
}

func (r *ThrottleRepository) Delete(ctx context.Context, m *model.Throttle) error {
	qb := sq.Delete(m.TableName()).PlaceholderFormat(sq.Dollar).RunWith(r.client)
	qb = qb.Where("hash_key = ?", m.HashKey)
	qb = qb.Where("key_type = ?", m.KeyType)

	_, err := qb.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("ThrottleRepository Delete: %w", err)
	}

	return nil
}
