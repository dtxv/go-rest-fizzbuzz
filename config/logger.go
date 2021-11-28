package config

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

// Config logger struct
type LoggerConfig struct {
	JSON  bool
	Level zerolog.Level
}

// SeverityHook duplicates the level attribute into a severity attribute.
type SeverityHook struct{}

// Run logger
func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("severity", level.String())
	}
}

// New creates a new logger instance and returns it.
func New(config *LoggerConfig) (logger zerolog.Logger) {
	if config.JSON {
		logger = zerolog.New(os.Stdout)
	} else {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return logger.Hook(SeverityHook{}).
		With().
		Caller().
		Timestamp().
		Logger().
		Level(config.Level)
}
