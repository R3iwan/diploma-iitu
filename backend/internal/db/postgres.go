package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/r3iwan/mse-business-go/internal/config"
)

func ConnectPostgres() *pgx.Conn {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error reading configs")
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
	}

	fmt.Println("Connected to Postgres successfully!")
	return conn
}
