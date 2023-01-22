package middleware

import "github.com/labstack/echo/v4"

type Router interface {
	Use(...echo.MiddlewareFunc)
}
