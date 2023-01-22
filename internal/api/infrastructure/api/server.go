package api

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/config"
	"github.com/basslove/daradara/internal/api/infrastructure/api/handler"
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	"github.com/basslove/daradara/internal/api/infrastructure/db/postgresql"
	"github.com/basslove/daradara/internal/api/logger"
	"github.com/basslove/daradara/internal/api/registry"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) error {
	defer logger.Info(ctx, "shutdown completed")

	conf := config.Get()

	// signal interrupt
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// env := os.Getenv("GOLANG_ENV")
	psqldb, err := postgresql.NewClient(ctx, conf.DB)
	if err != nil {
		return fmt.Errorf("error postgresql.NewClient: %w", err)
	}
	defer psqldb.Close()

	options := []registry.RepositoryOption{
		registry.WithPsql(psqldb),
	}
	repository := registry.NewRepository(options...)
	fmt.Println(repository)

	// echo framework
	e := echo.New()

	// v1 api
	apiRouter := e.Group("/v1")

	// middleware
	e.Use(echoMiddleware.Recover())
	corsConfig := echoMiddleware.DefaultCORSConfig
	e.Use(echoMiddleware.CORSWithConfig(corsConfig))

	// openAPI service
	handlersBuilder, err := handler.NewHandler(repository)
	if err != nil {
		return err
	}
	openapi_service.RegisterHandlers(apiRouter, handlersBuilder)

	// shutdown
	go func() {
		<-ctx.Done()
		fmt.Println("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, "Shutdown error: %w", err)
		}
	}()

	// Start Server
	e.Logger.Error(e.Start(fmt.Sprintf(":1323")))

	return nil
}
