package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r3iwan/mse-business-go/internal/delivery"
	"github.com/r3iwan/mse-business-go/internal/services"
)

// RegisterAuthRoutes registers authentication-related routes.
func RegisterAuthRoutes(r *gin.Engine, authService services.AuthService) {
	// Create a new handler with the injected AuthService
	handler := delivery.NewAuthHandler(authService)

	// Define the routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.RegisterCustomerHandler)
		auth.POST("/login", handler.LoginCustomerHanlder)
	}
}
