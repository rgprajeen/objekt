package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/julienschmidt/httprouter"
	"go.prajeen.com/objekt/internal/logger"
)

type CLI struct {
	Hostname string `help:"Hostname of the Objekt server" default:"localhost" short:"H"`
	Port     int    `help:"Port of the Objekt server" default:"8080" short:"p"`
}

func main() {
	cli := CLI{}
	kong.Parse(&cli)

	router := httprouter.New()
	router.GET("/objekt", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		logger.Get().Debug().Msg("Received request to /objekt")
		fmt.Fprint(w, "Objekt Server")
	})

	listener := fmt.Sprintf("%s:%d", cli.Hostname, cli.Port)
	logger.Get().Info().Str("listener", listener).Msgf("Starting Objekt Server at http://%s", listener)
	logger.Get().Fatal().Err(http.ListenAndServe(listener, router)).Msg("Objekt server closed")
}
