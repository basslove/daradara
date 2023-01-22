package openapi_service

import (
	"github.com/basslove/daradara/internal/api/domain/model"
)

func BuildOperatorPostOperatorsSignInResponse(m *model.Operator) *Operator {
	if m == nil {
		return nil
	}

	level := OperatorLevel(m.Level)
	belong := OperatorBelong(m.Belong)

	operator := Operator{
		Email:          &m.Email,
		Name:           &m.Name,
		DisplayName:    &m.DisplayName,
		LastAccessedAt: &m.LastAccessedAt,
		LastLoggedInAt: &m.LastLoggedInAt,
		IsGod:          &m.IsGod,
		Level:          &level,
		Belong:         &belong,
	}

	return &operator
}
