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
)

func TestSightGenreRepository_FindOneByID(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupSightGenreRepository())

	repository := db.NewSightGenreRepository(testutil.PsqlDB())

	t.Run("exists", func(t *testing.T) {
		value, err := repository.FindOneByID(ctx, 1, true)
		require.NoError(t, err)
		require.Equal(t, uint64(1), value.ID)
	})

	t.Run("not exists", func(t *testing.T) {
		_, err := repository.FindOneByID(ctx, 999, true)
		require.Error(t, err)
	})

	t.Run("not exists by soft_deleted", func(t *testing.T) {
		_, err := repository.FindOneByID(ctx, 3, true)
		require.Error(t, err)
	})
}

func TestSightGenreRepository_FindByNameAndCategoryID(t *testing.T) {
	ctx := context.Background()
	require.NoError(t, setupSightGenreRepository())

	repository := db.NewSightGenreRepository(testutil.PsqlDB())

	t.Run("exists by name", func(t *testing.T) {
		values, err := repository.FindByNameAndCategoryID(ctx, "test", 0, 0, 30)
		require.NoError(t, err)
		require.Len(t, values, 2)
	})

	t.Run("exists by sight_category_id", func(t *testing.T) {
		values, err := repository.FindByNameAndCategoryID(ctx, "", 1, 0, 30)
		require.NoError(t, err)
		require.Len(t, values, 2)
	})

	t.Run("not exists", func(t *testing.T) {
		values, err := repository.FindByNameAndCategoryID(ctx, "ueeeei", 0, 0, 30)
		require.NoError(t, err)
		require.Len(t, values, 0)
	})

	t.Run("not exists by soft_deleted", func(t *testing.T) {
		values, err := repository.FindByNameAndCategoryID(ctx, "hoge", 2, 0, 30)
		require.NoError(t, err)
		require.Len(t, values, 0)
	})
}

func setupSightGenreRepository() error {
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
		query := sq.Insert("sight_categories").Columns("id", "name", "is_valid")
		query = query.Values(c.ID, c.Name, c.IsValid).RunWith(testutil.PsqlDB()).PlaceholderFormat(sq.Dollar)
		_, err := query.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, g := range sightGenres {
		query := sq.Insert("sight_genres").Columns("id", "name", "sight_category_id", "is_valid")
		query = query.Values(g.ID, g.Name, g.SightCategoryID, g.IsValid).RunWith(testutil.PsqlDB()).PlaceholderFormat(sq.Dollar)
		_, err := query.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
