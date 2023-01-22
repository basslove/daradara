package logger

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	Logger     *zap.Logger        // 高速, 低アロケーション, 構造化Logスタイル
	BaseLogger *zap.SugaredLogger // 低速, 高アロケーション, printLogスタイル
)

func init() {
	fmt.Println("init logger")

	logger, err := NewLogger()
	if err != nil {
		log.Fatal(errors.Errorf("fail. init logger: %v", err))
	}

	Logger = logger
	BaseLogger = logger.Sugar()
}

func NewLogger(opts ...LogOption) (*zap.Logger, error) {
	logger, err := newLoggerConfig(opts...).Build()
	if err != nil {
		return nil, errors.Errorf("fail. init logger: %v", err)
	}

	return logger, nil
}

type contextKey = struct{}

func NewContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func fromContext(ctx context.Context) *zap.SugaredLogger {
	if logger := ctx.Value(contextKey{}); logger != nil {
		return logger.(*zap.SugaredLogger)
	}
	return BaseLogger
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	return fromContext(ctx)
}

func With(ctx context.Context, args ...interface{}) *zap.SugaredLogger {
	return fromContext(ctx).With(args...)
}

func Info(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Info(args...)
}

func Infof(ctx context.Context, templateFmt string, args ...interface{}) {
	fromContext(ctx).Infof(templateFmt, args...)
}

func Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	fromContext(ctx).Infow(msg, keysAndValues...)
}

func Warn(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Warn(args...)
}

func Warnf(ctx context.Context, templateFmt string, args ...interface{}) {
	fromContext(ctx).Warnf(templateFmt, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Error(args...)
}

func Errorf(ctx context.Context, templateFmt string, args ...interface{}) {
	fromContext(ctx).Errorf(templateFmt, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Debug(args...)
}

func Debugf(ctx context.Context, templateFmt string, args ...interface{}) {
	fromContext(ctx).Debugf(templateFmt, args...)
}

func Fatal(args ...interface{}) {
	Error(context.TODO(), args...)
	os.Exit(1)
}

func Fatalf(templateFmt string, args ...interface{}) {
	Errorf(context.TODO(), templateFmt, args...)
	os.Exit(1)
}
