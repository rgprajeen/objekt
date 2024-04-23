package config

import (
	"sync"

	"github.com/alecthomas/kong"
)

type globalConfig struct {
	Hostname       string    `help:"Hostname of the Objekt server" default:"localhost" short:"H"`
	Port           int       `help:"Port of the Objekt server" default:"8080" short:"p"`
	StorageBackend string    `help:"Select the storage backend for data" enum:"memory,database" default:"memory"`
	Log            LogConfig `embed:"" prefix:"log."`
	DB             DBConfig  `embed:"" prefix:"db."`
}

const (
	StorageBackendMemory   = "memory"
	StorageBackendDatabase = "database"
)

var config globalConfig
var once sync.Once

func Get() *globalConfig {
	once.Do(func() {
		kong.Parse(&config,
			kong.Name("objekt"),
			kong.Description("A object storage facade for the modern cloud"),
			kong.DefaultEnvars("OBJEKT"),
			kong.Configuration(kong.JSON, "/etc/objekt/config.json", "~/.objekt/config.json", ".objekt.json"),
			kong.UsageOnError())
	})

	return &config
}
