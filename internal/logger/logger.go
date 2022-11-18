package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New(ctx context.Context, level string, file string) *Logger {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Println("failed to open logFile")
		return nil
	}
	log := logrus.New()
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		log.Println("failed to parse log level")
	}
	log.SetLevel(lv)
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	go func() {
		<-ctx.Done()
		if err := f.Close(); err != nil {
			log.Println("failed to close logFile")
			return
		}
		fmt.Println("LogFile was closed")
	}()

	return &Logger{logger: log}
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Logger) Trace(msg string) {
	l.logger.Trace(msg)
}
