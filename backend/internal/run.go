package internal

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/config"
	"github.com/r3iwan/mse-business-go/internal/db"
	"github.com/r3iwan/mse-business-go/internal/delivery"
	"github.com/r3iwan/mse-business-go/internal/repository"
	"github.com/r3iwan/mse-business-go/internal/routes"
	"github.com/r3iwan/mse-business-go/internal/services"
)

func RunProgram() {
	_, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	conn := db.ConnectPostgres()
	defer conn.Close(context.Background())

	authRepo := repository.NewAuthRepository(conn)
	authService := services.NewAuthService(authRepo)
	authHandler := delivery.NewAuthHandler(authService)
	compRepo := repository.NewCompanyRepository(conn)
	compService := services.NewCompanyService(compRepo)
	compHandler := delivery.NewCompHandler(compService)
	superAdminRepo := repository.NewSuperAdminRepository(conn)
	superAdminService := services.NewSuperAdminServices(superAdminRepo, authService, compService)

	r := gin.Default()
	routes.RegisterAuthRoutes(r, *authHandler)
	routes.RegisterCompanyRoutes(r, *compHandler)
	routes.RegisterSuperAdminRoutes(r, superAdminService, *authHandler, *compHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
