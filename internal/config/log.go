package config

import "github.com/rs/zerolog"

type LogConfig struct {
	Level zerolog.Level `help:"Log level" enum:"debug,error,fatal,info,panic,warn" default:"info"`
	Mode  string        `help:"Log Mode" enum:"development,production" default:"development"`
	File  string        `help:"Log file to use in production mode" type:"path" default:"objekt.log"`
}

const (
	LogModeProduction  = "production"
	LogModeDevelopment = "development"
)
