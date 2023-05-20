package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Database string
}

func GetConfig() (Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "DEV"
	}

	switch env {
	case "DEV":
		viper.SetConfigFile("configs/config.dev.json")
	case "PROD":
		viper.SetConfigFile("configs/config.prod.json")
	default:
		return Config{}, fmt.Errorf("unexpected env %v", env)
	}

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
