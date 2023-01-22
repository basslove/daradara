package openapi_service

import (
	"github.com/basslove/daradara/internal/api/domain/model"
)

func BuildOperatorGetSightCategoriesResponse(ms []*model.SightCategory) []SightCategory {
	sightCategories := make([]SightCategory, 0, len(ms))

	for _, m := range ms {
		s := SightCategory{
			Id:   m.ID,
			Name: m.Name,
		}
		sightCategories = append(sightCategories, s)
	}

	return sightCategories
}
