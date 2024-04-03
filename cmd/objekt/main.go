package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.prajeen.com/objekt/internal/config"
	"go.prajeen.com/objekt/internal/logger"
)

func main() {
	cliConfig := config.Parse()

	router := httprouter.New()
	router.GET("/objekt", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		logger.Get().Debug().Msg("Received request to /objekt")
		fmt.Fprint(w, "Objekt Server")
	})

	listener := fmt.Sprintf("%s:%d", cliConfig.Hostname, cliConfig.Port)
	logger.Get().Info().Str("listener", listener).Msgf("Starting Objekt Server at http://%s", listener)
	logger.Get().Fatal().Err(http.ListenAndServe(listener, router)).Msg("Objekt server closed")
}
