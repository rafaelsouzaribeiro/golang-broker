package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Conf struct {
	PasswordRedis    string `mapstructure:"PASSWORD_REDIS"`
	RabbitmqUser     string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	RabbitmqPassWord string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	RabbitmqVhost    string `mapstructure:"RABBITMQ_DEFAULT_VHOST"`
}

func LoadConfig(path string) (*Conf, error) {
	_ = godotenv.Overload(filepath.Join(path, ".env"))

	cfg := &Conf{
		PasswordRedis:    os.Getenv("PASSWORD_REDIS"),
		RabbitmqUser:     os.Getenv("RABBITMQ_DEFAULT_USER"),
		RabbitmqPassWord: os.Getenv("RABBITMQ_DEFAULT_PASS"),
		RabbitmqVhost:    os.Getenv("RABBITMQ_DEFAULT_VHOST"),
	}

	return cfg, nil
}
