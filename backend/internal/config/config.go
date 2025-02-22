package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
}

func NewConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	postgresConfig, err := NewPostgresConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Postgres: *postgresConfig,
	}, nil
}

func NewPostgresConfig() (*PostgresConfig, error) {
	return &PostgresConfig{
		DBName:   viper.GetString("POSTGRES_DB"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		User:     viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}, nil
}
