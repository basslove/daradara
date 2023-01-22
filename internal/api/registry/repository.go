package registry

import (
	"github.com/basslove/daradara/internal/api/domain/query"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/basslove/daradara/internal/api/interface/gateway/db"
	q "github.com/basslove/daradara/internal/api/interface/gateway/query"
)

type Repository struct {
	// repository
	CustomerRepository      repository.CustomerRepository
	SightCategoryRepository repository.SightCategoryRepository
	SightGenreRepository    repository.SightGenreRepository
	ThrottleRepository      repository.ThrottleRepository
	OperatorRepository      repository.OperatorRepository

	// query
	SightGenreRelationQuery query.SightGenreRelationQuery
}

func NewRepository(options ...RepositoryOption) *Repository {
	conf := &RepositoryConfig{}
	for _, o := range options {
		o.Apply(conf)
	}
	r := Repository{}

	if conf.psqldb != nil {
		// repository
		r.CustomerRepository = db.NewCustomerRepository(conf.psqldb)
		r.SightCategoryRepository = db.NewSightCategoryRepository(conf.psqldb)
		r.SightGenreRepository = db.NewSightGenreRepository(conf.psqldb)
		r.ThrottleRepository = db.NewThrottleRepository(conf.psqldb)
		r.OperatorRepository = db.NewOperatorRepository(conf.psqldb)

		// query
		r.SightGenreRelationQuery = q.NewSightGenreRelationQuery(conf.psqldb)
	}

	return &r
}
