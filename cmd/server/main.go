package main

import (
	"net/http"

	"github.com/YutaIke/go-api-experiment/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func main() {
	config, err := config.NewConfig(".env.local") // TODO: set file name for each environment
	if err != nil {
		log.Fatal(err)
	}

	logger, _ := zap.NewProduction()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}()
	zap.ReplaceGlobals(logger)

	e := echo.New()
	e.Logger.SetLevel(log.ERROR) // e.Logger is error logger, not access logger.

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRequestID:     true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       true,
		LogUserAgent:     true,
		LogReferer:       true,
		LogContentLength: true,
		LogRoutePath:     true,
		LogStatus:        true,
		LogHeaders:       []string{"user-agent", "x-forwarded-for", "x-real-ip", "x-request-id"},
		LogQueryParams:   []string{"user"},
		LogLatency:       true,
		LogResponseSize:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			zap.L().Info("access log",
				zap.String("request_id", values.RequestID),
				zap.String("protocol", values.Protocol),
				zap.String("remote_ip", values.RemoteIP),
				zap.String("host", values.Host),
				zap.String("method", values.Method),
				zap.String("URI", values.URI),
				zap.String("URIPath", values.URIPath),
				zap.String("user_agent", values.UserAgent),
				zap.String("referer", values.Referer),
				zap.String("content_length", values.ContentLength),
				zap.String("route_path", values.RoutePath),
				zap.Int("status", values.Status),
				zap.Any("headers", values.Headers),
				zap.Any("query_params", values.QueryParams),
				zap.Duration("latency", values.Latency),
				zap.Int64("response_size", values.ResponseSize),
			)
			return nil
		},
	}))

	// Common settings for all routes
	e.Use(middleware.Recover()) // recover from panic
	e.Use(middleware.Secure())  // provide protection against injection attacks
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
	}))
	e.Use(middleware.RequestID())
	// TODO: Error Handling to send error to Sentory or so.
	// e.HTTPErrorHandler = func(err error, c echo.Context) { ... }
	e.GET("/", hello)

	e.Logger.Fatal(e.Start(config.Address))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
