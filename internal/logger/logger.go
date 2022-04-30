package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger  struct {
	logger zerolog.Logger
}

zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

func New(level string) *Logger {
	return &Logger{}
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	// TODO
}

// TODO
