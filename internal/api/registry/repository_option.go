package registry

import "github.com/jmoiron/sqlx"

type RepositoryConfig struct {
	psqldb *sqlx.DB
}

type RepositoryOption interface {
	Apply(*RepositoryConfig)
}

type RepositoryOptionPsqlDB struct {
	client *sqlx.DB
}

func (o RepositoryOptionPsqlDB) Apply(conf *RepositoryConfig) {
	conf.psqldb = o.client
}

func WithPsql(client *sqlx.DB) RepositoryOption {
	return RepositoryOptionPsqlDB{
		client: client,
	}
}
