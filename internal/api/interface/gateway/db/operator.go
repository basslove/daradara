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

func NewOperatorRepository(client *sqlx.DB) repository.OperatorRepository {
	return &OperatorRepository{Repository: Repository{client: client}}
}

type OperatorRepository struct {
	Repository
}

func (r *OperatorRepository) FindOneByEmail(ctx context.Context, email string) (*model.Operator, error) {
	values := make([]*model.Operator, 0)

	m := model.Operator{}
	qb := sq.Select("*").From(m.TableName()).PlaceholderFormat(sq.Dollar)
	qb = qb.Where("email = ?", email)

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("OperatorRepository FindOneByEmail: %w", err)
	}
	stmt, err := r.client.Preparex(sql)
	if err != nil {
		return nil, fmt.Errorf("OperatorRepository FindOneByEmail: %w", err)
	}
	if err = stmt.SelectContext(ctx, &values, args...); err != nil {
		return nil, fmt.Errorf("OperatorRepository FindOneByEmail: %w", err)
	}

	var value *model.Operator
	if len(values) > 0 {
		value = values[0]
	}

	return value, nil
}

func (r *OperatorRepository) Create(ctx context.Context, c *model.Operator) (int64, error) {
	var newID int64

	sql := fmt.Sprintf("INSERT INTO %s (email, crypted_password, name, level, belong, display_name, image_url, is_god, is_valid) VALUES (:email, :crypted_password, :name, :level, :belong, :display_name, :image_url, :is_god, :is_valid) RETURNING id;", c.TableName())
	rows, err := r.client.NamedQueryContext(ctx, sql, c)
	if err != nil {
		return 0, fmt.Errorf("OperatorRepository Create: %w", err)
	}
	if rows.Next() {
		rows.Scan(&newID)
	}

	return newID, nil
}

func (r *OperatorRepository) Update(ctx context.Context, m *model.Operator) error {
	m.UpdatedAt = time.Now()

	sql := fmt.Sprintf("UPDATE %s SET email=:email, crypted_password=:crypted_password, name=:name, level=:level, belong=:belong, display_name=:display_name, image_url=:image_url, is_god=:is_god, is_valid=:is_valid, last_accessed_at=:last_accessed_at, last_logged_in_at=:last_logged_in_at, updated_at=:updated_at WHERE id = :id", m.TableName())
	_, err := r.client.NamedQueryContext(ctx, sql, m)
	if err != nil {
		return fmt.Errorf("OperatorRepository Update: %w", err)
	}

	return nil
}
