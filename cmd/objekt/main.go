package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	obj_http "go.prajeen.com/objekt/internal/adapter/http"
	"go.prajeen.com/objekt/internal/adapter/storage/memory/repository"
	"go.prajeen.com/objekt/internal/config"
	"go.prajeen.com/objekt/internal/core/service"
	"go.prajeen.com/objekt/internal/logger"
)

func main() {
	cliConfig := config.Parse()

	logConfig := &logger.Config{Level: cliConfig.LogLevel}
	log := logConfig.Get()

	bucketRepo := repository.NewBucketRespository()
	fileRepo := repository.NewFileRepository(bucketRepo)
	bucketSvc := service.NewBucketService(log, bucketRepo)
	fileSvc := service.NewFileService(bucketRepo, fileRepo)

	router := httprouter.New()
	bucketHandler := obj_http.NewBucketHandler(log, router, bucketSvc)
	fileHandler := obj_http.NewFileHandler(router, fileSvc)
	bucketHandler.AddRoutes()
	fileHandler.AddRoutes()

	listener := fmt.Sprintf("%s:%d", cliConfig.Hostname, cliConfig.Port)
	log.Info().Str("listener", listener).Msgf("Starting Objekt Server at http://%s", listener)
	log.Fatal().Err(http.ListenAndServe(listener, router)).Msg("Objekt server closed")
}
