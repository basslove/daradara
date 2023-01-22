package main

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/config"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/infrastructure/db/postgresql"
	"log"
)

func main() {
	ctx := context.Background()
	conf := config.Get()

	fmt.Println("############## set dev seed : START ##############")

	// conn
	sqlxConn, err := postgresql.NewClient(ctx, conf.DB)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer sqlxConn.Close()

	// TRUNCATE
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE sight_genres restart identity CASCADE")
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE sight_categories restart identity CASCADE")

	// INSERT
	tx := sqlxConn.MustBegin()
	// sight_categories
	tx.NamedExecContext(ctx, "INSERT INTO sight_categories (name) VALUES (:name)", map[string]interface{}{"name": "観光施設"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_categories (name) VALUES (:name)", map[string]interface{}{"name": "宿泊"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_categories (name) VALUES (:name)", map[string]interface{}{"name": "飲食"})
	tx.Commit()

	// SELECT
	sightCategories := make([]*model.SightCategory, 0)
	err = sqlxConn.SelectContext(ctx, &sightCategories, "SELECT * FROM sight_categories ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range sightCategories {
		fmt.Println(v)
	}

	fmt.Println("############## set dev seed : END ##############")
}
