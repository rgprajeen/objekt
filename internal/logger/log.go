package logger

import (
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/attoleap/objekt/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var log zerolog.Logger

func Get() *zerolog.Logger {
	once.Do(func() {
		logConfig := config.Get().Log
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			return filepath.Base(file) + ":" + strconv.Itoa(line)
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:           os.Stdout,
			TimeFormat:    time.RFC3339,
			FieldsExclude: []string{"git_revision", "go_version"},
		}

		if logConfig.Mode == config.LogModeProduction {
			logFile := &lumberjack.Logger{
				Filename:   logConfig.File,
				MaxSize:    5,
				MaxAge:     14,
				MaxBackups: 10,
				Compress:   true,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, logFile)
		}

		var gitRevision string
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		log = zerolog.New(output).
			Level(logConfig.Level).
			With().
			Timestamp().
			Caller().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return &log
}
