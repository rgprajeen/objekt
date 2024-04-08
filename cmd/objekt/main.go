package main

import (
	"fmt"
	"net/http"

	ohttp "go.prajeen.com/objekt/internal/adapter/http"
	"go.prajeen.com/objekt/internal/adapter/storage/memory/repository"
	"go.prajeen.com/objekt/internal/config"
	"go.prajeen.com/objekt/internal/core/service"
	"go.prajeen.com/objekt/internal/logger"
)

func main() {
	cliConfig := config.Parse()

	logConfig := &logger.Config{Level: cliConfig.LogLevel}
	log := logConfig.Get()

	repo := repository.NewBucketRespository()
	svc := service.NewBucketService(repo)
	handler := ohttp.NewBucketHandler(svc)

	listener := fmt.Sprintf("%s:%d", cliConfig.Hostname, cliConfig.Port)
	log.Info().Str("listener", listener).Msgf("Starting Objekt Server at http://%s", listener)
	log.Fatal().Err(http.ListenAndServe(listener, handler.GetRouter())).Msg("Objekt server closed")
}
