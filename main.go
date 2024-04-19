package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Puttipong1/assessment-tax/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := config.NewConfig()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		err := e.Start(config.ServerConfig().Port())
		if err != nil && err != http.ErrServerClosed { // Start server
			fmt.Println("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Info(err)
	}
}
