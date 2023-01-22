package middleware

import (
	"github.com/basslove/daradara/internal/api/domain/service"
	"github.com/labstack/echo/v4"
)

func CustomerSetting(router Router, ca service.CustomerAuthenticator) {
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := GetToken(ctx)
			if token == "" {
				return next(ctx)
			}
			customer, err := ca.FindOneByToken(ctx.Request().Context(), token)
			if err != nil {
				return err
			}
			if customer != nil {
				SetCustomer(ctx, customer)
			}
			return next(ctx)
		}
	})
}
