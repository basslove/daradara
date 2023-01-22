package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

var (
	Logger     *zap.Logger
	BaseLogger *zap.SugaredLogger
)

func init() {
	fmt.Println("logger init")
}

type contextKey = struct{}

func Info(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Info(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	fromContext(ctx).Errorf(format, args...)
}

func NewContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func fromContext(ctx context.Context) *zap.SugaredLogger {
	if logger := ctx.Value(contextKey{}); logger != nil {
		return logger.(*zap.SugaredLogger)
	}
	return BaseLogger
}
