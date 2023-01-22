package query_test

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/interface/gateway/query"
	"github.com/basslove/daradara/internal/api/pkg/testutil"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSightGenreRelationQuery_FindByNameAndCategoryID(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupSightGenreRelationQuery())

	q := query.NewSightGenreRelationQuery(testutil.PsqlDB())

	t.Run("exists", func(t *testing.T) {
		values, err := q.FindByNameAndCategoryID(ctx, "test", 1, 0, 0)
		require.NoError(t, err)
		require.Len(t, values, 2)
		require.Equal(t, uint64(1), values[0].SightCategoryID)
	})

	t.Run("not exists", func(t *testing.T) {
		values, err := q.FindByNameAndCategoryID(ctx, "test", 999, 0, 0)
		require.NoError(t, err)
		require.Len(t, values, 0)
	})

	t.Run("not exists by soft_deleted", func(t *testing.T) {
		values, err := q.FindByNameAndCategoryID(ctx, "hoge", 0, 0, 0)
		require.NoError(t, err)
		require.Len(t, values, 0)
	})
}

func setupSightGenreRelationQuery() error {
	testutil.CleanPsqlDB()
	sightCategories := []*model.SightCategory{
		{
			ID:      uint64(1),
			Name:    "test1",
			IsValid: true,
		},
		{
			ID:      uint64(2),
			Name:    "test2",
			IsValid: true,
		},
	}
	sightGenres := []*model.SightGenre{
		{
			ID:              uint64(1),
			Name:            "test1",
			SightCategoryID: uint64(1),
			IsValid:         true,
		},
		{
			ID:              uint64(2),
			Name:            "test2",
			SightCategoryID: uint64(1),
			IsValid:         true,
		},
		{
			ID:              uint64(3),
			Name:            "hoge",
			SightCategoryID: uint64(2),
		},
	}
	for _, c := range sightCategories {
		q := sq.Insert("sight_categories").Columns("id", "name", "is_valid")
		q = q.Values(c.ID, c.Name, c.IsValid).RunWith(testutil.PsqlDB()).PlaceholderFormat(sq.Dollar)
		_, err := q.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, g := range sightGenres {
		q := sq.Insert("sight_genres").Columns("id", "name", "sight_category_id", "is_valid")
		q = q.Values(g.ID, g.Name, g.SightCategoryID, g.IsValid).RunWith(testutil.PsqlDB()).PlaceholderFormat(sq.Dollar)
		_, err := q.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
