package model

import (
	"fmt"
	"time"
)

type SightCategory struct {
	ID        uint64    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	IsValid   bool      `db:"is_valid" json:"is_valid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (m *SightCategory) TableName() string {
	return "sight_categories"
}

func (m *SightCategory) TableNameAlias() string {
	return fmt.Sprintf("%s AS sc", m.TableName())
}
