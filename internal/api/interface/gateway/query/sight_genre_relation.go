package query

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/query"
	"github.com/basslove/daradara/internal/api/interface/gateway/db"
	"github.com/jmoiron/sqlx"
)

func NewSightGenreRelationQuery(client *sqlx.DB) query.SightGenreRelationQuery {
	return &SightGenreRelationQuery{Repository: *db.NewRepository(client)}
}

type SightGenreRelationQuery struct {
	db.Repository
}

func (r *SightGenreRelationQuery) FindByNameAndCategoryID(ctx context.Context, name string, sightCategoryID, offset, limit uint64) ([]*model.SightGenreRelation, error) {
	values := make([]*model.SightGenreRelation, 0)

	selectFields := []string{
		"sg.id AS id",
		"sg.name AS name",
		"sg.image_url AS image_url",
		"sg.is_valid AS is_valid",
		"sg.created_at AS created_at",
		"sg.updated_at AS updated_at",
		"sg.sight_category_id AS sight_category_id",
		"sc.name AS sight_category_name",
	}

	sgm := model.SightGenre{}
	scm := model.SightCategory{}
	qb := sq.Select(selectFields...).From(sgm.TableNameAlias()).PlaceholderFormat(sq.Dollar)
	qb = qb.Join(fmt.Sprintf("%s ON sg.sight_category_id = sc.id", scm.TableNameAlias()))
	qb = qb.Where("sc.is_valid = ?", true)
	qb = qb.Where("sg.is_valid = ?", true)

	if len(name) != 0 {
		qb = qb.Where("sg.name LIKE ?", fmt.Sprint("%", name, "%"))
	}

	if sightCategoryID > 0 {
		qb = qb.Where("sg.sight_category_id = ?", sightCategoryID)
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	if limit > 0 {
		qb = qb.Limit(limit)
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("SightGenreRelationQuery FindByNameAndCategoryID: %w", err)
	}
	stmt, err := r.Repository.Client().Preparex(sql)
	if err != nil {
		return nil, fmt.Errorf("SightGenreRelationQuery FindByNameAndCategoryID: %w", err)
	}
	if err = stmt.SelectContext(ctx, &values, args...); err != nil {
		return nil, fmt.Errorf("SightGenreRelationQuery FindByNameAndCategoryID: %w", err)
	}

	return values, nil
}
