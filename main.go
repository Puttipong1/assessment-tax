package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		err := server.Echo.Start(server.Config.ServerConfig().Port())
		if err != nil {
			log.Fatal().Err(err).Msg(common.ShutDownServerMessage)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Warn().Msg("Shutting down...")
	if err := server.DB.DB.Close(); err != nil {
		log.Fatal().Err(err)
	}
	if err := server.Echo.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
