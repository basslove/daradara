package model

import (
	"time"
)

type Operator struct {
	ID              uint64    `db:"id" json:"id"`
	Email           string    `db:"email" json:"email"`
	CryptedPassword string    `db:"crypted_password" json:"crypted_password"`
	Name            string    `db:"name" json:"name"`
	DisplayName     string    `db:"display_name" json:"display_name"`
	ImageURL        string    `db:"image_url" json:"image_url"`
	Level           string    `db:"level" json:"level"`
	Belong          string    `db:"belong" json:"belong"`
	IsGod           bool      `db:"is_god" json:"is_god"`
	IsValid         bool      `db:"is_valid" json:"is_valid"`
	LastAccessedAt  time.Time `db:"last_accessed_at" json:"last_accessed_at"`
	LastLoggedInAt  time.Time `db:"last_logged_in_at" json:"last_logged_in_at"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

func (m *Operator) TableName() string {
	return "operators"
}
