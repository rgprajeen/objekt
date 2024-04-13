package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	obj_http "go.prajeen.com/objekt/internal/adapter/http"
	"go.prajeen.com/objekt/internal/adapter/storage/postgres"
	"go.prajeen.com/objekt/internal/adapter/storage/postgres/repository"
	"go.prajeen.com/objekt/internal/config"
	"go.prajeen.com/objekt/internal/core/service"
	"go.prajeen.com/objekt/internal/logger"
)

func main() {
	cliConfig := config.Parse()

	logConfig := &logger.Config{Level: cliConfig.LogLevel}
	log := logConfig.Get()

	dbConf := config.NewDB("localhost", 5432, "objekt_adm", "objekt123", "postgres", "objekt_db", map[string]string{
		"sslmode": "disable",
	})
	db, err := postgres.NewDB(context.Background(), dbConf)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	bucketRepo := repository.NewBucketRepository(db)
	fileRepo := repository.NewFileRepository(db)
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
