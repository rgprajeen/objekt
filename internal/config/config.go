package config

import (
	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
)

type GlobalConfig struct {
	Hostname string        `help:"Hostname of the Objekt server" default:"localhost" short:"H"`
	Port     int           `help:"Port of the Objekt server" default:"8080" short:"p"`
	LogLevel zerolog.Level `help:"Log level" default:"info" short:"l"`
	DB       DBConfig      `embed:"" prefix:"db."`
}

func Parse() *GlobalConfig {
	gc := &GlobalConfig{}
	kong.Parse(gc,
		kong.Name("objekt"),
		kong.Description("A object storage facade for the modern cloud"),
		kong.DefaultEnvars("OBJEKT"),
		kong.Configuration(kong.JSON, "/etc/objekt/config.json", "~/.objekt/config.json", ".objekt.json"),
		kong.UsageOnError())
	return gc
}
