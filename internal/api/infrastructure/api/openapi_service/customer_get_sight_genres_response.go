package openapi_service

import (
	"github.com/basslove/daradara/internal/api/domain/model"
)

func BuildCustomerGetSightGenresResponse(ms []*model.SightGenreRelation) []SightGenre {
	sightGenres := make([]SightGenre, 0, len(ms))

	for _, m := range ms {
		sc := SightCategory{
			Id:   m.SightCategoryID,
			Name: m.SightCategoryName,
		}
		sg := SightGenre{
			Id:            m.ID,
			Name:          m.Name,
			ImageUrl:      m.ImageURL,
			SightCategory: sc,
		}
		sightGenres = append(sightGenres, sg)
	}

	return sightGenres
}
