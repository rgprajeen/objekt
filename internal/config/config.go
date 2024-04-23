package config

import (
	"sync"

	"github.com/alecthomas/kong"
)

type GlobalConfig struct {
	Hostname string    `help:"Hostname of the Objekt server" default:"localhost" short:"H"`
	Port     int       `help:"Port of the Objekt server" default:"8080" short:"p"`
	Log      LogConfig `embed:"" prefix:"log."`
	DB       DBConfig  `embed:"" prefix:"db."`
}

var config GlobalConfig
var once sync.Once

func Load() {
	once.Do(func() {
		kong.Parse(&config,
			kong.Name("objekt"),
			kong.Description("A object storage facade for the modern cloud"),
			kong.DefaultEnvars("OBJEKT"),
			kong.Configuration(kong.JSON, "/etc/objekt/config.json", "~/.objekt/config.json", ".objekt.json"),
			kong.UsageOnError())
	})
}

func Global() *GlobalConfig {
	return &config
}

func Log() *LogConfig {
	return &config.Log
}
