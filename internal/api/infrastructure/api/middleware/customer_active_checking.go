package middleware

import (
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/labstack/echo/v4"
)

func inValidCustomerAllowedPaths() []string {
	return []string{"/v1/sign_in", "/v1/sign_up"}
}

func CustomerActiveChecking(router Router) {
	router.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(eCtx echo.Context) error {
				customer := GetCustomer(eCtx)
				if customer == nil {
					return next(eCtx)
				}
				if customer.IsValid {
					return next(eCtx)
				}

				for _, path := range inValidCustomerAllowedPaths() {
					if eCtx.Path() == path {
						return next(eCtx)
					}
				}
				return errors.ErrUnauthorized
			}
		})
}
