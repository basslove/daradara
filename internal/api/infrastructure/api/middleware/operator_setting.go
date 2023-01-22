package middleware

import (
	"github.com/basslove/daradara/internal/api/domain/service"
	"github.com/labstack/echo/v4"
)

func OperatorSetting(router Router, oa service.OperatorAuthenticator) {
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := GetToken(ctx)
			if token == "" {
				return next(ctx)
			}
			operator, err := oa.FindOneByToken(ctx.Request().Context(), token)
			if err != nil {
				return err
			}
			if operator != nil {
				SetOperator(ctx, operator)
			}
			return next(ctx)
		}
	})
}
