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

	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(config.Address))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
