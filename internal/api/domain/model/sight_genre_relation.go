package model

import "time"

/*
 - sight_genres(*)
 - sight_categories
*/

type SightGenreRelation struct {
	ID                uint64    `db:"id"`
	Name              string    `db:"name"`
	ImageURL          string    `db:"image_url"`
	IsValid           bool      `db:"is_valid"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	SightCategoryID   uint64    `db:"sight_category_id"`
	SightCategoryName string    `db:"sight_category_name"`
}
