package postgresql

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewClient(ctx context.Context, conf config.DBConfig) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.PsqlDatabaseUser, conf.PsqlDatabasePass, conf.PsqlHostName, conf.PsqlPort, conf.PsqlDatabaseName)
	db, err := sqlx.ConnectContext(ctx, "postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(conf.PsqlConnMaxIdleTime) // 最大アイドル時間。0は無制限
	db.SetMaxOpenConns(conf.PsqlMaxOpenConns)       // オープン接続の最大数。0は無制限
	db.SetMaxIdleConns(conf.PsqlMaxIdleConns)       // アイドル接続の最大数。0は保持しない
	db.SetConnMaxLifetime(conf.PsqlConnMaxLifetime) // 最大接続再利用時間。0は無制限
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewTestClient(ctx context.Context, conf config.TestDBConfig) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", conf.PsqlDatabaseUser, conf.PsqlDatabasePass, conf.PsqlHostName, conf.PsqlPort, conf.PsqlDatabaseName)
	db, err := sqlx.ConnectContext(ctx, "postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(conf.PsqlConnMaxIdleTime) // 最大アイドル時間。0は無制限
	db.SetMaxOpenConns(conf.PsqlMaxOpenConns)       // オープン接続の最大数。0は無制限
	db.SetMaxIdleConns(conf.PsqlMaxIdleConns)       // アイドル接続の最大数。0は保持しない
	db.SetConnMaxLifetime(conf.PsqlConnMaxLifetime) // 最大接続再利用時間。0は無制限
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
