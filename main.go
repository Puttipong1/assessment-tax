package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	log := logger.Get()
	config := config.NewConfig()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		time.Sleep(90 * time.Second)
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		err := e.Start(config.ServerConfig().Port())
		if err != nil && err != http.ErrServerClosed { // Start server
			log.Fatal().Err(err).Msg("Shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
