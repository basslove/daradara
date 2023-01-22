package registry

import (
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/basslove/daradara/internal/api/interface/gateway/db"
)

type Repository struct {
	SightCategoryRepository repository.SightCategoryRepository
}

func NewRepository(options ...RepositoryOption) *Repository {
	conf := &RepositoryConfig{}
	for _, o := range options {
		o.Apply(conf)
	}
	r := Repository{}

	if conf.psqldb != nil {
		r.SightCategoryRepository = db.NewSightCategoryRepository(conf.psqldb)
	}

	return &r
}
