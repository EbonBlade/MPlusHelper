package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger returns zerolog instance.
func Logger() *zerolog.Logger {
	return &log.Logger
}
