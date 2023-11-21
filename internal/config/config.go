package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

type Config struct {
	MySQL   `yaml:"mysql"`
	MongoDB `yaml:"mongoDB"`
}

type MySQL struct {
	UserName string `yaml:"MYSQL_USER" env-default:"root"`
	Password string `yaml:"MYSQL_PASSWORD" env-required:"true"`
	DbName   string `yaml:"MYSQL_DATABASE" env-required:"true"`
	Port     int    `yaml:"MYSQL_PORT" env-default:"8080"`
	Host     string `yaml:"MYSQL_HOST" env-default:"localhost:8080"`
}

type MongoDB struct {
	Host       string `yaml:"MONGO_HOST" env-default:"localhost"`
	Port       string `yaml:"MONGO_PORT" env-default:"27017"`
	DataBase   string `yaml:"DBNAME" env-default:"admin"`
	Collection string `yaml:"COLLECTION" env-required:"true"`
}

func MustLoad() *Config {
	configPath := "../internal/config/config.yaml"

	var cfg Config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error("error with file: %v", err)
	}
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		slog.Error("error with parse: %v", err)
	}
	return &cfg
}
