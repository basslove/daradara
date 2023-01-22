package model

import (
	"time"
)

type Customer struct {
	ID                 uint64    `db:"id" json:"id"`
	Email              string    `db:"email" json:"email"`
	CryptedPassword    string    `db:"crypted_password" json:"crypted_password"`
	Name               string    `db:"name" json:"name"`
	Gender             string    `db:"gender" json:"gender"`
	Generation         string    `db:"generation" json:"generation"`
	DisplayName        string    `db:"display_name" json:"display_name"`
	Birthday           time.Time `db:"birthday" json:"birthday"`
	PhoneNumber        string    `db:"phone_number" json:"phone_number"`
	Introduction       string    `db:"introduction" json:"introduction"`
	ImageURL           string    `db:"image_url" json:"image_url"`
	AllowPlanDisplayed bool      `db:"allow_plan_displayed" json:"allow_plan_displayed"`
	IsValid            bool      `db:"is_valid" json:"is_valid"`
	LastAccessedAt     time.Time `db:"last_accessed_at" json:"last_accessed_at"`
	LastLoggedInAt     time.Time `db:"last_logged_in_at" json:"last_logged_in_at"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
}

func (m *Customer) TableName() string {
	return "customers"
}
