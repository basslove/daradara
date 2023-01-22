package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/jmoiron/sqlx"
)

func NewSightCategoryRepository(client *sqlx.DB) repository.SightCategoryRepository {
	return &SightCategoryRepository{Repository: Repository{client: client}}
}

type SightCategoryRepository struct {
	Repository
}

func (r *SightCategoryRepository) FindOneByID(ctx context.Context, sightCategoryID uint64, isValid bool) (*model.SightCategory, error) {
	var value model.SightCategory

	stmt, err := r.client.Preparex("SELECT * FROM sight_categories WHERE id = $1 AND is_valid = $2")
	if err != nil {
		return nil, fmt.Errorf("SightCategoryRepository FindOneByID: %w", err)
	}
	if err = stmt.GetContext(ctx, &value, sightCategoryID, isValid); err != nil {
		return nil, fmt.Errorf("SightCategoryRepository FindOneByID: %w", err)
	}

	return &value, nil
}

func (r *SightCategoryRepository) FindByName(ctx context.Context, name string, offset, limit uint64) ([]*model.SightCategory, error) {
	values := make([]*model.SightCategory, 0)

	m := model.SightCategory{}
	qb := sq.Select("*").From(m.TableName()).PlaceholderFormat(sq.Dollar)
	qb = qb.Where("is_valid = ?", true)

	if len(name) != 0 {
		qb = qb.Where("name LIKE ?", fmt.Sprint("%", name, "%"))
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	if limit > 0 {
		qb = qb.Limit(limit)
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("SightCategoryRepository FindByName: %w", err)
	}
	stmt, err := r.client.Preparex(sql)
	if err != nil {
		return nil, fmt.Errorf("SightCategoryRepository FindByName: %w", err)
	}
	if err = stmt.SelectContext(ctx, &values, args...); err != nil {
		return nil, fmt.Errorf("SightCategoryRepository FindByName: %w", err)
	}

	return values, nil
}
