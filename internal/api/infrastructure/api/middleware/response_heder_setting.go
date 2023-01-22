package middleware

import "github.com/labstack/echo/v4"

var resHeaders = map[string]string{
	"Cache-control": "no-store",
	"Pragma":        "no-cache",
}

func ResponseHeaderSetting(router Router) {
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			resp := eCtx.Response()
			for k, v := range resHeaders {
				resp.Header().Set(k, v)
			}
			return next(eCtx)
		}
	})
}
