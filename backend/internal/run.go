package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/config"
	"github.com/r3iwan/mse-business-go/internal/db"
	"github.com/r3iwan/mse-business-go/internal/repository"
	"github.com/r3iwan/mse-business-go/internal/routes"
	"github.com/r3iwan/mse-business-go/internal/services"
)

func RunProgram() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Loaded Config: %+v\n", cfg)

	conn := db.ConnectPostgres()
	defer conn.Close(context.Background())

	authRepo := repository.NewAuthRepository(conn)
	authService := services.NewAuthService(authRepo)

	r := gin.Default()
	routes.RegisterAuthRoutes(r, authService)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
