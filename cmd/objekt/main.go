package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/julienschmidt/httprouter"
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
		fmt.Fprint(w, "Objekt Server")
	})

	listener := fmt.Sprintf("%s:%d", cli.Hostname, cli.Port)
	fmt.Printf("Starting Objekt Server on http://%s\n", listener)
	http.ListenAndServe(listener, router)
}
