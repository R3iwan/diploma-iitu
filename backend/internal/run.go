package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/r3iwan/mse-business-go/internal/config"
	"github.com/r3iwan/mse-business-go/internal/db"
)

func RunProgram() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Loaded Config: %+v\n", cfg)

	conn := db.ConnectPostgres()
	defer conn.Close(context.Background())
}
