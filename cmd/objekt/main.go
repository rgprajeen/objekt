package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	obj_http "go.prajeen.com/objekt/internal/adapter/http"
	m_repo "go.prajeen.com/objekt/internal/adapter/storage/memory/repository"
	"go.prajeen.com/objekt/internal/adapter/storage/postgres"
	p_repo "go.prajeen.com/objekt/internal/adapter/storage/postgres/repository"
	"go.prajeen.com/objekt/internal/config"
	"go.prajeen.com/objekt/internal/core/port"
	"go.prajeen.com/objekt/internal/core/service"
	"go.prajeen.com/objekt/internal/logger"
)

func main() {
	cliConfig := config.Get()

	log := logger.Get()

	var bucketRepo port.BucketRepository
	var fileRepo port.FileRepository

	if cliConfig.StorageBackend == config.StorageBackendDatabase {
		db, err := postgres.NewDB(context.Background())
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to database")
		}

		bucketRepo = p_repo.NewBucketRepository(db)
		fileRepo = p_repo.NewFileRepository(db)
	} else {
		bucketRepo = m_repo.NewBucketRepository()
		fileRepo = m_repo.NewFileRepository(bucketRepo)
	}

	bucketSvc := service.NewBucketService(log, bucketRepo, fileRepo)
	fileSvc := service.NewFileService(log, bucketRepo, fileRepo)

	router := httprouter.New()
	bucketHandler := obj_http.NewBucketHandler(log, router, bucketSvc)
	fileHandler := obj_http.NewFileHandler(log, router, fileSvc)
	bucketHandler.AddRoutes()
	fileHandler.AddRoutes()

	listener := fmt.Sprintf("%s:%d", cliConfig.Hostname, cliConfig.Port)
	log.Info().Str("listener", listener).Msgf("Starting Objekt Server at http://%s", listener)
	log.Fatal().Err(http.ListenAndServe(listener, router)).Msg("Objekt server closed")
}
