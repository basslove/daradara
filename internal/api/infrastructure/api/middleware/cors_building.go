package middleware

import (
	"github.com/basslove/daradara/internal/api/config"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"strings"
)

func CORSBuilding() (echoMiddleware.CORSConfig, error) {
	corsConfig := echoMiddleware.DefaultCORSConfig
	origins := strings.Split(config.Get().WebFrontend.CORSOrigins, "~")
	corsConfig.AllowOrigins = make([]string, 0, len(origins))
	//for _, origin := range origins {
	//	u, err := url.Parse(origin)
	//	if err != nil {
	//		return echoMiddleware.CORSConfig{}, err
	//	}
	//	corsConfig.AllowOrigins = append(origins, fmt.Sprintf("%s://%s", u.Scheme, u.Host))
	//}
	return corsConfig, nil
}
