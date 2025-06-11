package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPC        GRPC
	StoragePath string `env:"STORAGE_PATH" env-required:"true"`
}
type GRPC struct {
	Port    string        `env:"GRPC_PORT_API_DB" env-default:"8081"`
	TimeOut time.Duration `env:"GRPC_TIME_OUT_API_DB"`
}

func InitConfig(log *slog.Logger) *Config {
	cfgPath := ".env"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Debug("config not found from", "CfgPath", cfgPath, "Error:", err.Error())
		log.Info("not found config")
		os.Exit(1)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Debug("cant read configfile", "err:", err.Error())
		log.Info("read config file to failed")
		os.Exit(1)
	}
	return &cfg
}
