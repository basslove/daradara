package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/jmoiron/sqlx"
)

func NewSightGenreRepository(client *sqlx.DB) repository.SightGenreRepository {
	return &SightGenreRepository{Repository: Repository{client: client}}
}

type SightGenreRepository struct {
	Repository
}

func (r *SightGenreRepository) FindOneByID(ctx context.Context, sightGenreID uint64, isValid bool) (*model.SightGenre, error) {
	var value model.SightGenre

	stmt, err := r.client.Preparex("SELECT * FROM sight_genres WHERE id = $1 AND is_valid = $2")
	if err != nil {
		return nil, fmt.Errorf("SightGenreRepository FindOneByID: %w", err)
	}
	if err = stmt.GetContext(ctx, &value, sightGenreID, isValid); err != nil {
		return nil, fmt.Errorf("SightGenreRepository FindOneByID: %w", err)
	}

	return &value, nil
}

func (r *SightGenreRepository) FindByNameAndCategoryID(ctx context.Context, name string, sightCategoryID, offset, limit uint64) ([]*model.SightGenre, error) {
	values := make([]*model.SightGenre, 0)

	m := model.SightGenre{}
	qb := sq.Select("*").From(m.TableName()).PlaceholderFormat(sq.Dollar)
	qb = qb.Where("is_valid = ?", true)

	if len(name) != 0 {
		qb = qb.Where("name LIKE ?", fmt.Sprint("%", name, "%"))
	}
	if sightCategoryID > 0 {
		qb = qb.Where("sight_category_id = ?", sightCategoryID)
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	if limit > 0 {
		qb = qb.Limit(limit)
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("SightGenreRepository FindByNameAndCategoryID: %w", err)
	}
	stmt, err := r.client.Preparex(sql)
	if err != nil {
		return nil, fmt.Errorf("SightGenreRepository FindByNameAndCategoryID: %w", err)
	}
	if err = stmt.SelectContext(ctx, &values, args...); err != nil {
		return nil, fmt.Errorf("SightGenreRepository FindByNameAndCategoryID: %w", err)
	}

	return values, nil
}
