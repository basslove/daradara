package db

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	client *sqlx.DB
}

func NewRepository(client *sqlx.DB) *Repository {
	return &Repository{client: client}
}

func (r *Repository) Client() *sqlx.DB {
	return r.client
}
