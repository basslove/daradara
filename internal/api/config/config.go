package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	Server      ServerConfig
	WebFrontend WebFrontendConfig
	Storage     StorageConfig
	DB          DBConfig
	TestDB      TestDBConfig
}

type ServerConfig struct {
	Port      uint   `default:"1323"`
	JwtSecret string `default:"secret"`
	Debug     bool   `default:"false"`
}

type WebFrontendConfig struct {
	CORSOrigins string `default:"*" split_words:"true"`
}

type StorageConfig struct {
	StaticBucket string `default:"kokufu-ichiba-static-dev" split_words:"true"`
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
	_, file, _, _ := runtime.Caller(0)
	envPath := filepath.Join(filepath.Join(filepath.Dir(file), "..", "..", ".."), ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Printf("can't read .env: %v", err)
	}

	once.Do(func() {
		if err := envconfig.Process("server", &conf.Server); err != nil {
			log.Fatal(err.Error())
		}
		if err := envconfig.Process("webfront", &conf.WebFrontend); err != nil {
			log.Fatal(err.Error())
		}
		if err := envconfig.Process("storage", &conf.Storage); err != nil {
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
