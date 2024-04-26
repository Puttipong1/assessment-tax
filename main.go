package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/route"
)

func main() {
	log := config.Logger()
	server := server.NewServer()
	route.ConfigureRoutes(server)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		err := server.Echo.Start(server.Config.ServerConfig().Port())
		if err != nil && err != http.ErrServerClosed { // Start server
			log.Fatal().Err(err).Msg(common.ShutDownServerMessage)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.DB.DB.Close(); err != nil {
		log.Fatal().Err(err).Msg(common.ShutDownServerMessage)
	}
	if err := server.Echo.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg(common.ShutDownServerMessage)
	}
}
