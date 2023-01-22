package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
	"time"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
	TestDB TestDBConfig
}

type ServerConfig struct {
	Port uint `default:"1323"`
}

type DBConfig struct {
	Debug               bool          `default:"false"`
	PsqlHostName        string        `default:"localhost" split_words:"true"`
	PsqlPort            int           `default:"9432" split_words:"true"`
	PsqlDatabaseName    string        `default:"dev01" split_words:"true"`
	PsqlDatabaseUser    string        `default:"dev01" split_words:"true"`
	PsqlDatabasePass    string        `default:"dev01" split_words:"true"`
	PsqlConnMaxIdleTime time.Duration `default:"0" split_words:"true"`
	PsqlMaxOpenConns    int           `default:"30" split_words:"true"`
	PsqlMaxIdleConns    int           `default:"30" split_words:"true"`
	PsqlConnMaxLifetime time.Duration `default:"5m" split_words:"true"`
}

// 本来は以下不要(環境単位にenv切替すればよし). but env環境構築めんどいので設ける.
type TestDBConfig struct {
	Debug               bool          `default:"true"`
	PsqlHostName        string        `default:"localhost" split_words:"true"`
	PsqlPort            int           `default:"9433" split_words:"true"`
	PsqlDatabaseName    string        `default:"test01" split_words:"true"`
	PsqlDatabaseUser    string        `default:"test01" split_words:"true"`
	PsqlDatabasePass    string        `default:"test01" split_words:"true"`
	PsqlConnMaxIdleTime time.Duration `default:"0" split_words:"true"`
	PsqlMaxOpenConns    int           `default:"30" split_words:"true"`
	PsqlMaxIdleConns    int           `default:"30" split_words:"true"`
	PsqlConnMaxLifetime time.Duration `default:"5m" split_words:"true"`
}

func Get() Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("can't read .env: %v", err)
	}

	once.Do(func() {
		if err := envconfig.Process("server", &conf.Server); err != nil {
			log.Fatal(err.Error())
		}
		if err := envconfig.Process("db", &conf.DB); err != nil {
			log.Fatal(err.Error())
		}
		if err := envconfig.Process("testdb", &conf.TestDB); err != nil {
			log.Fatal(err.Error())
		}

	})
	return conf
}
