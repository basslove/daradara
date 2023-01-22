package middleware

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"strconv"
	"time"
)

func RequestLogging(router Router) {
	router.Use(echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogLatency:       true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogRoutePath:     true,
		LogRequestID:     true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogValuesFunc: func(eCtx echo.Context, rv echoMiddleware.RequestLoggerValues) error {
			info := struct {
				Time         string `json:"time"`
				ID           string `json:"id"`
				RemoteIP     string `json:"remote_ip"`
				Host         string `json:"host"`
				Method       string `json:"method"`
				URI          string `json:"uri"`
				RoutePath    string `json:"route_path"`
				UserAgent    string `json:"user_agent"`
				Status       int    `json:"status"`
				Error        string `json:"error"`
				Latency      int64  `json:"latency"`
				LatencyHuman string `json:"latency_human"`
				BytesIn      int    `json:"bytes_in"`
				BytesOut     int    `json:"bytes_out"`
			}{
				Time:         time.Now().Format(time.RFC3339Nano),
				ID:           rv.RequestID,
				RemoteIP:     rv.RemoteIP,
				Host:         rv.Host,
				Method:       rv.Method,
				URI:          rv.URI,
				RoutePath:    rv.RoutePath,
				UserAgent:    rv.UserAgent,
				Status:       rv.Status,
				LatencyHuman: rv.Latency.String(),
				Latency:      int64(rv.Latency),
				BytesOut:     int(rv.ResponseSize),
			}
			// json
			if rv.Error != nil {
				b, _ := json.Marshal(rv.Error.Error())
				info.Error = string(b[1 : len(b)-1])
			}

			// content length
			if hcl := eCtx.Request().Header.Get(echo.HeaderContentLength); hcl != "" {
				hcli, err := strconv.Atoi(hcl)
				if err == nil {
					info.BytesIn = hcli
				}
			}

			// finale
			b, err := json.Marshal(info)
			if err != nil {
				return err
			}
			b = append(b, "\n"...)
			_, err = eCtx.Logger().Output().Write(b)
			return err
		},
	}))
}
