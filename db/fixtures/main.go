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
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE operators restart identity CASCADE")
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE throttles restart identity CASCADE")
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE customers restart identity CASCADE")
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE sight_genres restart identity CASCADE")
	sqlxConn.ExecContext(ctx, "TRUNCATE TABLE sight_categories restart identity CASCADE")

	// INSERT
	tx := sqlxConn.MustBegin()
	// sight_categories
	tx.NamedExecContext(ctx, "INSERT INTO sight_categories (id, name) VALUES (:id, :name)", map[string]interface{}{"id": 1, "name": "観光施設"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_categories (id, name) VALUES (:id, :name)", map[string]interface{}{"id": 2, "name": "宿泊"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_categories (id, name) VALUES (:id, :name)", map[string]interface{}{"id": 3, "name": "飲食"})
	// sight_genres
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name, image_url) VALUES (:id, :sight_category_id, :name, :image_url)", map[string]interface{}{"id": 1, "sight_category_id": 1, "name": "公園", "image_url": "https://s3-ap-northeast-1.amazonaws.com/aaaa.jp/images/genre/%E3%82%A2%E3%82%A48A0.jpg"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name) VALUES (:id, :sight_category_id, :name)", map[string]interface{}{"id": 2, "sight_category_id": 1, "name": "温泉"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name) VALUES (:id, :sight_category_id, :name)", map[string]interface{}{"id": 3, "sight_category_id": 1, "name": "美術館"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name, image_url) VALUES (:id, :sight_category_id, :name, :image_url)", map[string]interface{}{"id": 4, "sight_category_id": 2, "name": "旅館", "image_url": "https://s3-ap-northeast-1.amazonaws.com/aaaa.jp/images/genre/%E3%82%A2%E3%82%A48A0.jpg"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name) VALUES (:id, :sight_category_id, :name)", map[string]interface{}{"id": 5, "sight_category_id": 2, "name": "ホテル"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name, image_url) VALUES (:id, :sight_category_id, :name, :image_url)", map[string]interface{}{"id": 6, "sight_category_id": 3, "name": "アイスクリーム", "image_url": "https://s3-ap-northeast-1.amazonaws.com/aaaa.jp/images/genre/%E3%82%A2%E3%82%A48A0.jpg"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name) VALUES (:id, :sight_category_id, :name)", map[string]interface{}{"id": 7, "sight_category_id": 3, "name": "イタリアン"})
	tx.NamedExecContext(ctx, "INSERT INTO sight_genres (id, sight_category_id, name) VALUES (:id, :sight_category_id, :name)", map[string]interface{}{"id": 8, "sight_category_id": 3, "name": "スパニッシュ"})
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

	sightGenres := make([]*model.SightGenre, 0)
	err = sqlxConn.SelectContext(ctx, &sightGenres, "SELECT * FROM sight_genres ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range sightGenres {
		fmt.Println(v)
	}

	fmt.Println("############## set dev seed : END ##############")
}
