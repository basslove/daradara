package testutil

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/config"
	testPSQL "github.com/basslove/daradara/internal/api/infrastructure/db/postgresql"
	dbRepo "github.com/basslove/daradara/internal/api/interface/gateway/db"
	"github.com/jmoiron/sqlx"
	"gopkg.in/khaiql/dbcleaner.v2"
	"log"
	"sync"
)

var (
	Cleaner    = dbcleaner.New()
	psqldb     *sqlx.DB
	once       sync.Once
	repo       *dbRepo.Repository
	conf       config.TestDBConfig
	testTables []string
)

func init() {
	var err error
	if psqldb, err = NewPsqlConnection(); err != nil {
		log.Fatal(err)
	}

	repo = dbRepo.NewRepository(psqldb)
	testTables, err = Tables()
	if err != nil {
		log.Fatal(err)
	}
}

func PsqlDB() *sqlx.DB {
	return psqldb
}

func Repository() *dbRepo.Repository {
	return repo
}

func NewPsqlConnection() (*sqlx.DB, error) {
	return testPSQL.NewTestClient(context.Background(), dbConfig())
}

func dbConfig() config.TestDBConfig {
	once.Do(func() {
		conf = config.Get().TestDB
	})

	return conf
}

func CleanPsqlDB() {
	for _, t := range testTables {
		psqldb.Exec(fmt.Sprintf("TRUNCATE TABLE %s restart identity CASCADE", t))
	}
}

type pgTable struct {
	SchemaName string `db:"schemaname" json:"schemaname"`
	TableName  string `db:"tablename" json:"tablename"`
}

func Tables() ([]string, error) {
	const queryStr = `SELECT schemaname, tablename FROM pg_tables WHERE schemaname NOT LIKE 'pg_%' AND schemaname != 'information_schema' AND tablename != 'schema_migrations'`
	rows, err := psqldb.QueryxContext(context.Background(), queryStr)
	if err != nil {
		return nil, fmt.Errorf("TestTables(): %w", err)
	}

	tableNames := make([]string, 0)
	for rows.Next() {
		var pt pgTable
		if err := rows.StructScan(&pt); err != nil {
			return nil, fmt.Errorf("TestTables(): %w", err)
		}
		if pt.SchemaName == "public" {
			tableNames = append(tableNames, pt.TableName)
		}
	}

	return tableNames, nil
}
