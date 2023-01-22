package model

import (
	"fmt"
	"time"
)

type SightGenre struct {
	ID              uint64    `db:"id" json:"id"`
	SightCategoryID uint64    `db:"sight_category_id" json:"sight_category_id"`
	Name            string    `db:"name" json:"name"`
	ImageURL        string    `db:"image_url" json:"image_url"`
	IsValid         bool      `db:"is_valid" json:"is_valid"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

func (m *SightGenre) TableName() string {
	return "sight_genres"
}

func (m *SightGenre) TableNameAlias() string {
	return fmt.Sprintf("%s AS sg", m.TableName())
}
