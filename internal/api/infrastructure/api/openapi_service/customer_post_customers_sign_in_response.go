package openapi_service

import (
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/pkg/util"
)

func BuildCustomerPostCustomersSignInResponse(m *model.Customer) *Customer {
	if m == nil {
		return nil
	}

	birthday := util.TimeToStrYMD(m.Birthday)
	gender := CustomerGender(m.Gender)
	generation := CustomerGeneration(m.Generation)

	customer := Customer{
		Email:               &m.Email,
		Name:                &m.Name,
		Birthday:            &birthday,
		Introduction:        &m.Introduction,
		PhoneNumber:         &m.PhoneNumber,
		DisplayName:         &m.DisplayName,
		LastAccessedAt:      &m.LastAccessedAt,
		LastLoggedInAt:      &m.LastLoggedInAt,
		AllowPlansDisplayed: &m.AllowPlanDisplayed,
		Gender:              &gender,
		Generation:          &generation,
	}

	return &customer
}
