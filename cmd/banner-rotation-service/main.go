package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Feride3d/banner-rotation-service/internal/config"
	"github.com/Feride3d/banner-rotation-service/internal/logger"
	"github.com/Feride3d/banner-rotation-service/internal/queues"
	internalgrpc "github.com/Feride3d/banner-rotation-service/internal/server/grpc"
	"github.com/Feride3d/banner-rotation-service/internal/service"
	"github.com/Feride3d/banner-rotation-service/internal/service/mab"
	"github.com/Feride3d/banner-rotation-service/internal/storage"
	"github.com/streadway/amqp"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.yaml", "Path to configuration file")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	// Init configuration
	config := config.NewConfig(configFile)

	// Init logger
	logg := logger.New(ctx, config.Logger.Level, config.Logger.File)

	// Init storage
	storage, err := startStorage(ctx, config.DB)
	if err != nil {
		logg.Error("failed to init new storage: " + err.Error())
	}

	// Start RabbitMQ
	rb, err := amqp.Dial(config.Rabbitmq.URI)
	if err != nil {
		logg.Error(fmt.Errorf("failed to start rabbit queue, %w", err).Error())
	}

	// Init queue
	publisher := queues.NewPublisher(config.Rabbitmq.Queuename, rb)

	// Init mab algorithm
	mab := mab.NewMab()

	// Init application
	rotator := service.New(logg, storage, mab, publisher)

	// Init server
	server, err := internalgrpc.NewHttpServer(rotator, config.Server.Host, config.Server.Port, config.Server.GrpcPort)
	if err != nil {
		logg.Error("failed to start new server: " + err.Error())
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("rotation service is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

// Start storage
func startStorage(ctx context.Context, dbcfg config.DBConfig) (storage *storage.Storage, err error) {
	conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		dbcfg.User, dbcfg.Password, dbcfg.Host, dbcfg.Port,
		dbcfg.DBName, dbcfg.SSLMode)
	if storage, err = storage.Connect(ctx, conn); err != nil {
		return nil, fmt.Errorf("failed to start postgres database: %w", err)
	}
	return storage, nil
}
