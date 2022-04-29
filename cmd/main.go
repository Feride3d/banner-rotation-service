package main

import (
	"log"
	"strings"

	"github.com/Eun/go-hit/doctest/server"
	"github.com/Feride3d/banner-rotation-service/internal/handler"
	"github.com/Feride3d/banner-rotation-service/internal/service"
	"github.com/Feride3d/banner-rotation-service/internal/storage"
	"github.com/spf13/viper"
)

// инициализация вайпера
func initViper(configPath string) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("")

	viper.SetDefault("log", "debug")
	viper.SetDefault("listen", "localhost:8080")
	viper.SetDefault("db.url", "_")

}

// инициализация конфигурационных файлов
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig() // считывает значения конфигов и записывает их во внутр объект вайпера
}

func main() {
	if err := initConfig(); err != nil { // ициализируем конфиги
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	bd := storage.Storage() // объявляем зависимости
	services := service.NewService(bd)
	handlers := handler.NewHandler(services)

	s := new(server.Server)                                                       // ициализируем экземпляр сервера
	if err := s.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil { // запуск сервера на порту по ключу из вайпера
		log.Fatalf("error occured while running http server: %s", err.Error()) // при ошибке вывод ошибки и выход из программы
	}
}
