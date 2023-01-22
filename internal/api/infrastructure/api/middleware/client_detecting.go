package middleware

import (
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

const clientDetectingKey = "daradara:clientDetectingKey"

func ClientDetecting(router Router) {
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			client := newClient(ctx.Request())
			SetClient(ctx, client)

			req := ctx.Request()
			ctxLogger := logger.FromContext(req.Context())

			clg := ctxLogger.With(
				zap.String("version", client.Version),
				zap.String("platform", string(client.Platform)),
				zap.String("device", client.Device),
			)
			ctx.SetRequest(req.WithContext(logger.NewContext(req.Context(), clg)))

			return next(ctx)
		}
	})
}

func SetClient(ctx echo.Context, client *model.Client) {
	ctx.Set(clientDetectingKey, client)
}

func GetClient(ctx echo.Context) *model.Client {
	client := ctx.Get(clientDetectingKey)
	if v, ok := client.(*model.Client); ok {
		return v
	}
	return &model.Client{}
}

func newClient(r *http.Request) *model.Client {
	return model.NewClient(
		r.Header.Get("X-Client-Version"),
		r.Header.Get("X-Client-System-Version"),
		r.Header.Get("X-Client-Platform"),
		r.Header.Get("X-Client-Device"),
	)
}
