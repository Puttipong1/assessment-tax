package main

import (
	"net/http"

	"github.com/Puttipong1/assessment-tax/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.NewConfig()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.Logger.Fatal(e.Start(config.ServerConfig().Port()))
}
