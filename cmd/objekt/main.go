package main

import (
	"context"
	"net/http"
	_ "net/http/pprof"

	obj_http "github.com/attoleap/objekt/internal/adapter/http"
	m_repo "github.com/attoleap/objekt/internal/adapter/persistence/memory/repository"
	"github.com/attoleap/objekt/internal/adapter/persistence/postgres"
	p_repo "github.com/attoleap/objekt/internal/adapter/persistence/postgres/repository"
	"github.com/attoleap/objekt/internal/adapter/storage"
	"github.com/attoleap/objekt/internal/config"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/attoleap/objekt/internal/core/service"
	"github.com/attoleap/objekt/internal/logger"
	"github.com/julienschmidt/httprouter"
)

func main() {
	cliConfig := config.Get()

	log := logger.Get()

	var bucketRepo port.BucketRepository
	var fileRepo port.FileRepository

	if cliConfig.PersistenceMode == config.PersistenceModeDatabase {
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

	storageRepoProvider := &storage.StorageRepositoryProvider{}
	bucketSvc := service.NewBucketService(log, bucketRepo, fileRepo, storageRepoProvider)
	fileSvc := service.NewFileService(log, bucketRepo, fileRepo)

	router := httprouter.New()
	if cliConfig.Http.PprofEnabled {
		router.Handler(http.MethodGet, "/debug/pprof/*item", http.DefaultServeMux)
		log.Debug().Str("endpoint_format", "/debug/pprof/*item").Msg("Added routes to pprof endpoints")
	}
	bucketHandler := obj_http.NewBucketHandler(log, router, bucketSvc)
	fileHandler := obj_http.NewFileHandler(log, router, fileSvc)
	bucketHandler.AddRoutes()
	fileHandler.AddRoutes()

	listener := cliConfig.Http.ListenerURL()
	log.Info().Str("listener", listener).Msgf("Starting Objekt Server at http://%s", listener)
	log.Fatal().Err(http.ListenAndServe(listener, router)).Msg("Objekt server closed")
}
