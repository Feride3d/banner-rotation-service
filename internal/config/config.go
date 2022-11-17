package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Logger   LoggerConfig   // `yaml:"logger"`
	DB       DBConfig       // `yaml:"db"`
	Server   ServerConfig   //`yaml:"server"`
	Rabbitmq RabbitmqConfig //`yaml:"rabbitmq"`
}

type LoggerConfig struct {
	Level string // `yaml:"level"`
	File  string // `yaml:"file"`
}

type DBConfig struct {
	User     string // `yaml:"user"`
	Password string // `yaml:"password"`
	Host     string // `yaml:"host"`
	Port     string // `yaml:"port"`
	DBName   string // `yaml:"dbname"`
	SSLMode  string // `yaml:"sslmode"`
}

type RabbitmqConfig struct {
	URI       string // `yaml:"uri"`
	Queuename string // `yaml:"queuename"`
}

type ServerConfig struct {
	Host     string // `yaml:"host"`
	Port     string // `yaml:"port"`
	GrpcPort string // `yaml:"grpcport"`
}

func NewConfig(file string) (conf Config) {
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Errorf("reading config error: %w", err))
	}
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(fmt.Errorf("ummarshal to config struct is failed: %w", err))
	}
	return conf
}
