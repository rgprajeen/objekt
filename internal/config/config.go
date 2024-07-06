package config // import github.com/attoleap/objekt/internal/config

import (
	"sync"

	"github.com/alecthomas/kong"
)

type globalConfig struct {
	Http            HttpConfig  `embed:"" prefix:"http."`
	PersistenceMode string      `help:"Select the persistence mode for data" enum:"memory,database" default:"memory"`
	Log             LogConfig   `embed:"" prefix:"log."`
	DB              DBConfig    `embed:"" prefix:"db."`
	Cache           CacheConfig `embed:"" prefix:"cache."`
	Local           LocalConfig `embed:"" prefix:"local."`
	AWS             AWSConfig   `embed:"" prefix:"aws."`
	OCI             OCIConfig   `embed:"" prefix:"oci."`
}

const (
	PersistenceModeMemory   = "memory"
	PersistenceModeDatabase = "database"
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
