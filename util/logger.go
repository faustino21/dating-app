package util

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func NewLog(logLevel string) {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = logger
}
