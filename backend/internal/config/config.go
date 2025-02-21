package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
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
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}

	// Navigate to the project root by getting the parent directory once
	projectRoot := filepath.Dir(currentDir)
	envPath := filepath.Join(projectRoot, ".env")

	// Use os.Open() to manually read the .env file
	envFile, err := os.Open(envPath)
	if err != nil {
		log.Println("Error opening .env file:", err)
	} else {
		fmt.Println(".env file opened successfully.")
	}

	// Manually parse the .env file
	envMap, err := godotenv.Parse(envFile)
	if err != nil {
		log.Println("Error parsing .env file:", err)
	} else {
		for key, value := range envMap {
			os.Setenv(key, value)
		}
	}

	// Use Viper to get environment variables
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
