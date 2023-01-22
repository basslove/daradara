package logger

import (
	"github.com/basslove/daradara/internal/api/config"
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

type LogOption func(zc *zap.Config)

func newLoggerConfig(opts ...LogOption) zap.Config {
	var zc zap.Config
	if config.Get().Server.Debug {
		zc = zapdriver.NewDevelopmentConfig()
	} else {
		zc = zapdriver.NewProductionConfig()
	}
	for _, o := range opts {
		o(&zc)
	}
	return zc
}

func OutputStdout() LogOption {
	return LogOption(func(zc *zap.Config) { zc.OutputPaths = []string{"stdout"} })
}
