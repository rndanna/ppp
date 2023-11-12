package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Postgresql struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Name     string `mapstructure:"name"`
		} `mapstructure:"postgresql"`
	} `mapstructure:"database"`
	RabbitMQ struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Login    string `mapstructure:"login"`
		Password string `mapstructure:"password"`
	} `mapstructure:"rabbitmq"`
}

func Init() (*Config, error) {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	return &cfg, nil
}
