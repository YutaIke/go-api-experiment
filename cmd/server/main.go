package main

import (
	"net/http"

	"github.com/YutaIke/go-api-experiment/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	config, err := config.NewConfig(".env.local") // TODO: set file name for each environment
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Logger.SetLevel(log.ERROR) // e.Logger is error logger, not access logger.

	// Common settings for all routes
	e.Use(middleware.Logger())  // access logger
	e.Use(middleware.Recover()) // recover from panic
	e.Use(middleware.Secure())  // provide protection against injection attacks
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
	}))
	// TODO: Error Handling to send error to Sentory or so.
	// e.HTTPErrorHandler = func(err error, c echo.Context) { ... }

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(config.Address))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
