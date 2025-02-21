package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/config"
)

func ConnectPostgres() *pgx.Conn {
	cfg, err := config.NewPostgresConfig()
	if err != nil {
		fmt.Printf("Error converting POSTGRES_PORT (%s): %v, using default 5432\n", cfg.Port, err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
	}

	fmt.Println("Connected to Postgres successfully!")
	return conn
}
