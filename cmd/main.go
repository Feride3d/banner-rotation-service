package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Feride3d/banner-rotation-service/internal/logger"
	"github.com/Feride3d/banner-rotation-service/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/configs/config.yml", "Path to configuration file")
}

func main() {
	config := NewConfig()
	Logg := logger.New(config.Logger.Level)

	storage := storage.New()
	service := service.New(Logg, storage)
	server := http.NewServer(service, config.Server.Host, config.Server.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			Logg.Error("failed to stop http server")
		}
	}()

	Logg.Info("start http server...")

	if err := server.Start(ctx); err != nil {
		Logg.Error("failed to start http server")
		cancel()
		os.Exit(1)
	}
}

func NewConfig() Config {
	return Config{}
}
