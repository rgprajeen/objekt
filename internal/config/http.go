package config

import "fmt"

type HttpConfig struct {
	Hostname string `help:"Hostname of the Objekt server" default:"localhost" short:"H"`
	Port     int    `help:"Port of the Objekt server" default:"8080" short:"p"`
}

func (h *HttpConfig) ListenerURL() string {
	return fmt.Sprintf("http://%s:%d", h.Hostname, h.Port)
}
